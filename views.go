package main

import (
	"time"

	"github.com/go-sharp/cerberus/v2"
	"github.com/pkg/errors"
)

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
	RecoveryActions []ServiceRecoveryActionItem `json:"recovery_actions,omitempty"`
	StopSignal      int                         `json:"stop_signal,omitempty"`

	// SCM Properties (Admin rights require to load this properties)
	Dependencies []string `json:"dependencies,omitempty"`
	ServiceUser  string   `json:"service_user,omitempty"`
	Password     *string  `json:"password,omitempty"`
	StartType    int      `json:"start_type,omitempty"`
}

// ServiceRecoveryActionItem is the frontend view for the recovery actions.
type ServiceRecoveryActionItem struct {
	ExitCode    int      `json:"exit_code"`
	Action      int      `json:"action"`
	Delay       int      `json:"delay"`
	MaxRestarts int      `json:"max_restarts"`
	ResetAfter  int      `json:"reset_after"`
	Program     string   `json:"program"`
	Arguments   []string `json:"arguments"`
}

// DependOnSvcItem item for the DependenciesList component.
type DependOnSvcItem struct {
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

// mapSvcCfgToSvcItem maps a cerberus SvcConfig to a ServiceItem.
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

// mapSvcCfgRecoveryActionToViewItem maps a cerberus.RecoveryAction to a view recovery action item.
func mapSvcCfgRecoveryActionToViewItem(actions map[int]cerberus.SvcRecoveryAction) []ServiceRecoveryActionItem {
	uiActions := []ServiceRecoveryActionItem{}
	for _, v := range actions {
		uiActions = append(uiActions, ServiceRecoveryActionItem{
			ExitCode:    v.ExitCode,
			Action:      int(v.Action),
			Delay:       v.Delay,
			MaxRestarts: v.MaxRestarts,
			ResetAfter:  int(v.ResetAfter.Seconds()),
			Program:     v.Program,
			Arguments:   v.Arguments,
		})
	}

	return uiActions
}

// mapServiceItemToSvcConfig maps a ServiceItem back to a cerberus.SvcConfig.
func mapServiceItemToSvcConfig(service map[string]interface{}) cerberus.SvcConfig {
	// Map Basic Configuration
	svc := cerberus.SvcConfig{
		Name:            service["name"].(string),
		ExePath:         service["exe_path"].(string),
		StartType:       cerberus.ManualStartType,
		StopSignal:      cerberus.NoSignal,
		RecoveryActions: map[int]cerberus.SvcRecoveryAction{},
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

	// Map Extendend configuration
	if v, ok := service["start_type"].(float64); ok {
		svc.StartType = cerberus.StartType(v)
	}

	if v, ok := service["stop_signal"].(float64); ok {
		svc.StopSignal = cerberus.StopSignal(v)
	}

	if v, ok := service["recovery_actions"].([]interface{}); ok {
		for _, item := range v {
			if i, ok := item.(map[string]interface{}); ok {
				action, err := mapRecoverViewItemToSvcCfgRecAction(i)
				if err != nil {
					continue
				}
				svc.RecoveryActions[action.ExitCode] = *action
			}
		}
	}

	if v, ok := service["dependencies"]; ok {
		svc.Dependencies = mapInfToStrA(v)
	}

	if v, ok := service["service_user"].(string); ok {
		svc.ServiceUser = v
	}

	if v, ok := service["password"].(string); ok && v != "" {
		svc.Password = &v
	}

	return svc
}

func mapRecoverViewItemToSvcCfgRecAction(action map[string]interface{}) (*cerberus.SvcRecoveryAction, error) {
	svcAction := cerberus.SvcRecoveryAction{}

	if v, ok := action["exit_code"].(float64); ok {
		svcAction.ExitCode = int(v)
	}

	if v, ok := action["action"].(float64); ok {
		svcAction.Action = cerberus.RecoveryAction(v)
	}

	switch svcAction.Action {
	case cerberus.NoAction:
	case cerberus.RestartAction, cerberus.RunProgramAction, cerberus.RunAndRestartAction:
	default:
		return nil, errors.New("invalid action typ")
	}

	if v, ok := action["delay"].(float64); ok {
		svcAction.Delay = int(v)
	}

	if v, ok := action["max_restarts"].(float64); ok {
		svcAction.MaxRestarts = int(v)
	}

	if v, ok := action["reset_after"].(float64); ok {
		svcAction.ResetAfter = time.Duration(time.Second * time.Duration(v))
	}

	if v, ok := action["program"].(string); ok {
		svcAction.Program = v
	}

	if v, ok := action["arguments"]; ok {
		svcAction.Arguments = mapInfToStrA(v)
	}

	return &svcAction, nil
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
