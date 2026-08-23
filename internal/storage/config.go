package storage

import (
	// system packages
	"encoding/json"
	"os"

	// local packages
	"github.com/johnathantam/TodoTerminal/internal/storage/models"
)

// AddProjectToConfig registers projectName in the shared config.json,
func AddProjectToConfig(rootPath, projectName string) error {
	configPath := filepath.Join(rootPath, "config.json")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	// Unmarshal the JSON config
	var config models.ProjectsConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}

	// Add the new project to the list if it doesn't already exist
	for _, existingProject := range config.Projects {
		if existingProject == projectName {
			// project already exists
			return nil
		}
	}

	config.Projects = append(config.Projects, projectName)

	// Marshal the updated config back to JSON
	updatedData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	// Write the updated JSON
	err = os.WriteFile(configPath, updatedData, 0o644)
	if err != nil {
		return err
	}

	return nil
}
