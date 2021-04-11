package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-sharp/cerberus/v2"
	"github.com/wailsapp/wails"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

// Services contains the application logic for service handling in the frontend.
type Services struct {
	r   *wails.Runtime
	log *wails.CustomLogger
	mgr *mgr.Mgr

	// Properties that will be bound
	IsAdmin bool
}

// WailsInit initialize the struct with the wails runtime.
func (s *Services) WailsInit(runtime *wails.Runtime) error {
	s.r = runtime
	s.log = runtime.Log.New("Services")

	s.log.Debug("connecting to the local service manager")
	m, err := mgr.Connect()
	if err != nil {
		s.log.Errorf("failed to connect to service manager: %v", err)
		s.IsAdmin = false

		return nil
	}
	s.mgr, s.IsAdmin = m, true
	s.log.Debug("successfully connected to the local service manager")
	return nil
}

// LoadOverviewServices loads the service data for the overview page.
func (s *Services) LoadOverviewServices() (svcs []OverviewServiceItem, err error) {

	services, err := s.getServices()
	if err != nil {
		s.log.Error(err.Error())
		return svcs, err
	}

	for _, s := range services {
		svcs = append(svcs, OverviewServiceItem{
			Name:        s.cbCfg.Name,
			Description: s.cbCfg.Desc,
			DisplayName: s.cbCfg.DisplayName,
			StartType:   s.mgrCfg.StartType,
			State:       uint32(s.status.State),
		})
	}

	return svcs, nil
}

func (s *Services) GetDependOnSvc() (svcs []DependOnSvcItem, err error) {
	s.log.Debug("loading depend on services")

	services, err := s.mgr.ListServices()
	if err != nil {
		return svcs, err
	}

	for _, sname := range services {
		sv, err := s.mgr.OpenService(sname)
		if err != nil {
			s.log.Errorf("skipping service '%v': %v", sname, err)
			continue
		}

		cfg, err := sv.Config()
		if err != nil {
			s.log.Errorf("skipping service '%v': %v", sname, err)
			continue
		}

		svcs = append(svcs, DependOnSvcItem{
			Name:        sname,
			DisplayName: cfg.DisplayName,
		})
	}

	return svcs, nil
}

// LoadService gets the current configuration for a service.
func (s *Services) LoadService(name string) (svc ServiceItem, err error) {
	cfg, err := cerberus.LoadServiceCfg(name)
	if err != nil {
		return ServiceItem{}, err
	}

	return mapSvcCfgToSvcItem(cfg), nil
}

// InstallService installs a new service.
func (s *Services) InstallService(data map[string]interface{}) (err error) {
	defer s.handlePanic(&err)

	svc := mapServiceItemToSvcConfig(data)

	if err := cerberus.InstallService(svc); err != nil {
		return err
	}

	if err := cerberus.UpdateService(svc); err != nil {
		_ = cerberus.RemoveService(svc.Name)
		return err
	}

	return nil
}

func (s *Services) UpdateService(data map[string]interface{}) (err error) {
	defer s.handlePanic(&err)

	svc := mapServiceItemToSvcConfig(data)

	if err := cerberus.UpdateService(svc); err != nil {
		return err
	}

	return nil
}

// ShowFileDialog displays a file dialog to request a file path.
func (s *Services) ShowFileDialog(data map[string]interface{}) string {
	dir := false
	if d, ok := data["dir"]; ok {
		dir = d.(bool)
	}

	if dir {
		return s.r.Dialog.SelectDirectory()
	}
	return s.r.Dialog.SelectFile()
}

// DeleteService removes the service with the given name.
func (s *Services) DeleteService(name string) (err error) {
	return cerberus.RemoveService(name)
}

// RestartService stops and restart a service and waits until completion.
func (s *Services) RestartService(name string) (err error) {
	defer s.handlePanic(&err)

	if err = s.StopService(name); err != nil {
		return err
	}

	return s.StartService(name)
}

// StopService stops a running service and waits until it is stopped.
func (s *Services) StopService(name string) (err error) {
	defer s.handlePanic(&err)

	service, err := s.mgr.OpenService(name)
	if err != nil {
		return err
	}

	status, err := service.Query()
	if err != nil {
		return err
	}

	if status.State == svc.Stopped {
		return nil
	}

	status, err = service.Control(svc.Stop)
	if err != nil {
		return err
	}

	cancel := time.After(30 * time.Second)
	for {
		status, err = service.Query()
		if err != nil {
			return err
		}

		if status.State == svc.Stopped {
			return nil
		}

		select {
		case <-cancel:
			return errors.New("timeout: failed to stop service")
		default:
		}
	}
}

// StartService starts the passed service and waits until it is running.
func (s *Services) StartService(name string) (err error) {
	defer s.handlePanic(&err)

	service, err := s.mgr.OpenService(name)
	if err != nil {
		return err
	}

	status, err := service.Query()
	if err != nil {
		return err
	}

	if status.State == svc.Running {
		return nil
	}

	if err := service.Start(); err != nil {
		return err
	}

	cancel := time.After(30 * time.Second)
	for {
		status, err = service.Query()
		if err != nil {
			return err
		}

		if status.State == svc.Running {
			return nil
		}

		select {
		case <-cancel:
			return errors.New("timeout: failed to start service")
		default:
		}
	}
}

func (s *Services) getServices() (svcs []serviceInfo, err error) {
	if !s.IsAdmin {
		return svcs, fmt.Errorf("access denied: you need to be an administrator")
	}

	s.log.Debug("loading serivce configuration")
	services, err := cerberus.LoadServicesCfg()
	if err != nil {
		return svcs, err
	}

	for _, service := range services {
		s.log.Debugf("open service %v\n", service.Name)
		sv, err := s.mgr.OpenService(service.Name)
		if err != nil {
			return []serviceInfo{}, err
		}

		s.log.Debugf("loading service mgr config for %v\n", service.Name)
		cfg, err := sv.Config()
		if err != nil {
			return svcs, err
		}

		s.log.Debugf("query service state for %v\n", service.Name)
		status, err := sv.Query()
		if err != nil {
			return svcs, err
		}

		svcs = append(svcs, serviceInfo{
			cbCfg:  service,
			mgrCfg: cfg,
			status: status,
		})
	}

	return svcs, nil
}

func (s *Services) handlePanic(err *error) {
	if p := recover(); p != nil {
		s.log.Errorf("panic in function call: %v", p)
		if err != nil {
			*err = fmt.Errorf("unhandled error occured: %v", p)
		}
	}
}

type serviceInfo struct {
	cbCfg  *cerberus.SvcConfig
	mgrCfg mgr.Config
	status svc.Status
}
