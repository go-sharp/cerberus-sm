package main

import (
	"fmt"

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
	// services, err := s.getServices()
	// if err != nil {
	// 	s.log.Error(err.Error())
	// 	return svcs, err
	// }

	// for _, s := range services {
	// 	svcs = append(svcs, OverviewServiceItem{
	// 		Name:        s.cbCfg.Name,
	// 		Description: s.cbCfg.Desc,
	// 		DisplayName: s.cbCfg.DisplayName,
	// 		StartType:   s.mgrCfg.StartType,
	// 		State:       uint32(s.status.State),
	// 	})
	// }

	svcs = []OverviewServiceItem{
		{DisplayName: "Test Service 1", Description: "A test service 1 description", Name: "TestSvc1", StartType: 2, State: 1},
		{DisplayName: "Test Service 2", Description: "A test service 2 description", Name: "TestSvc2", StartType: 4, State: 3},
		{DisplayName: "Test Service 1", Description: "A test service 1 description", Name: "TestSvc1", StartType: 2, State: 1},
		{DisplayName: "Test Service 2", Description: "A test service 2 description", Name: "TestSvc2", StartType: 4, State: 3},
		{DisplayName: "Test Service 1", Description: "A test service 1 description", Name: "TestSvc1", StartType: 2, State: 1},
		{DisplayName: "Test Service 2", Description: "A test service 2 description", Name: "TestSvc2", StartType: 4, State: 3},
		{DisplayName: "Test Service 1", Description: "A test service 1 description", Name: "TestSvc1", StartType: 2, State: 1},
		{DisplayName: "Test Service 2", Description: "A test service 2 description", Name: "TestSvc2", StartType: 4, State: 3},
		{DisplayName: "Test Service 1", Description: "A test service 1 description", Name: "TestSvc1", StartType: 2, State: 1},
		{DisplayName: "Test Service 2", Description: "A test service 2 description", Name: "TestSvc2", StartType: 4, State: 3},
		{DisplayName: "Test Service 1", Description: "A test service 1 description", Name: "TestSvc1", StartType: 2, State: 1},
		{DisplayName: "Test Service 2", Description: "A test service 2 description", Name: "TestSvc2", StartType: 4, State: 3},
		{DisplayName: "Test Service 1", Description: "A test service 1 description", Name: "TestSvc1", StartType: 2, State: 1},
		{DisplayName: "Test Service 2", Description: "A test service 2 description", Name: "TestSvc2", StartType: 4, State: 3},
	}

	return svcs, nil
}

// InstallService installs a new service.
func (s *Services) InstallService(data map[string]interface{}) error {
	defer s.handlePanic()

	svc := mapServiceItemToSvcConfig(data)
	return cerberus.InstallService(svc)
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

func (s *Services) handlePanic() {
	if p := recover(); p != nil {
		s.log.Errorf("panic in function call: %v", p)
	}
}

type serviceInfo struct {
	cbCfg  *cerberus.SvcConfig
	mgrCfg mgr.Config
	status svc.Status
}
