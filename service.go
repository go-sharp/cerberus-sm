package main

import (
	"github.com/go-sharp/cerberus/v2"
	"github.com/wailsapp/wails"
)

type Services struct {
	r        *wails.Runtime
	log      *wails.CustomLogger
	svcStore *wails.Store
}

func (s *Services) WailsInit(runtime *wails.Runtime) error {
	s.r = runtime
	s.log = runtime.Log.New("Services")
	s.svcStore = runtime.Store.New("Services", []*cerberus.SvcConfig{})

	return nil
}

func (s *Services) ReloadServices() error {
	svcs, err := cerberus.LoadServicesCfg()
	if err != nil {
		s.log.Error(err.Error())
		return err
	}

	s.svcStore.Set(svcs)

	return nil
}
