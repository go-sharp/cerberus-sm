package main

import "github.com/go-sharp/cerberus/v2"

// OverviewServiceItem represents a service for the overview page.
type OverviewServiceItem struct {
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Description string `json:"description,omitempty"`
	State       uint32 `json:"state,omitempty"`
	StartType   uint32 `json:"start_type,omitempty"`
}

// ServiceItem represents a service.
type ServiceItem struct {
	// Base configuration
	Name        string   `json:"name,omitempty"`
	DisplayName string   `json:"display_name,omitempty"`
	Description string   `json:"description,omitempty"`
	ExePath     string   `json:"exe_path,omitempty"`
	WorkDir     string   `json:"work_dir,omitempty"`
	Args        []string `json:"args,omitempty"`
	Env         []string `json:"env,omitempty"`

	// Extended configurations
	RecoveryActions map[int]ServiceRecoveryActionItem `json:"recovery_actions,omitempty"`
	StopSignal      int                               `json:"stop_signal,omitempty"`

	// SCM Properties (Admin rights require to load this properties)
	Dependencies []string `json:"dependencies,omitempty"`
	ServiceUser  string   `json:"service_user,omitempty"`
	Password     *string  `json:"password,omitempty"`
	StartType    int      `json:"start_type,omitempty"`
}

// ServiceRecoveryActionItem is the frontend view for the recovery actions.
type ServiceRecoveryActionItem struct {
	ExitCode    int      `json:"exit_code,omitempty"`
	Action      int      `json:"action,omitempty"`
	Delay       int      `json:"delay,omitempty"`
	MaxRestarts int      `json:"max_restarts,omitempty"`
	ResetAfter  int      `json:"reset_after,omitempty"`
	Program     string   `json:"program,omitempty"`
	Arguments   []string `json:"arguments,omitempty"`
}

func mapSvcCfgToSvcItem(svcCfg *cerberus.SvcConfig) (svc ServiceItem) {
	return ServiceItem{
		Name:        svcCfg.Name,
		DisplayName: svcCfg.DisplayName,
		Description: svcCfg.Desc,
		ExePath:     svcCfg.ExePath,
		WorkDir:     svcCfg.WorkDir,
		Args:        svcCfg.Args,
		Env:         svcCfg.Env,

		RecoveryActions: mapSvcCfgRecoveryActionToViewItem(svcCfg.RecoveryActions),
		StopSignal:      int(svcCfg.StopSignal),
		Dependencies:    svcCfg.Dependencies,
		ServiceUser:     svcCfg.ServiceUser,
		StartType:       int(svcCfg.StartType),
	}
}

func mapSvcCfgRecoveryActionToViewItem(actions map[int]cerberus.SvcRecoveryAction) map[int]ServiceRecoveryActionItem {
	uiActions := map[int]ServiceRecoveryActionItem{}
	for k, v := range actions {
		uiActions[k] = ServiceRecoveryActionItem{
			ExitCode:    v.ExitCode,
			Action:      int(v.Action),
			Delay:       v.Delay,
			MaxRestarts: v.MaxRestarts,
			ResetAfter:  int(v.ResetAfter.Seconds()),
			Program:     v.Program,
			Arguments:   v.Arguments,
		}
	}
	return uiActions
}

func mapServiceItemToSvcConfig(service map[string]interface{}) cerberus.SvcConfig {
	svc := cerberus.SvcConfig{
		Name:    service["name"].(string),
		ExePath: service["exe_path"].(string),
	}

	if v, ok := service["display_name"].(string); ok {
		svc.DisplayName = v
	}

	if v, ok := service["description"].(string); ok {
		svc.Desc = v
	}

	if v, ok := service["work_dir"].(string); ok {
		svc.WorkDir = v
	}

	if v, ok := service["args"]; ok {
		svc.Args = mapInfToStrA(v)
	}

	if v, ok := service["env"]; ok {
		svc.Env = mapInfToStrA(v)
	}

	return svc
}

func mapInfToStrA(data interface{}) []string {
	var re []string
	if arr, ok := data.([]interface{}); ok {
		for _, v := range arr {
			if s, ok := v.(string); ok {
				re = append(re, s)
			}
		}
	}

	return re
}
