package models

type ProjectsConfig struct {
	ActiveProject string   `json:"active_project"`
	Projects      []string `json:"projects"`
}
