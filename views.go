package main

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
