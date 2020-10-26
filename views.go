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
	Name        string   `json:"name,omitempty"`
	DisplayName string   `json:"display_name,omitempty"`
	Description string   `json:"description,omitempty"`
	ExePath     string   `json:"exe_path,omitempty"`
	WorkDir     string   `json:"work_dir,omitempty"`
	Args        []string `json:"args,omitempty"`
	Env         []string `json:"env,omitempty"`
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
