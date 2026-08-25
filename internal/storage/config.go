package storage

import (
	// system packages
	"encoding/json"
	"os"

	// local packages
	"github.com/johnathantam/TodoTerminal/internal/storage/models"
)

// AddProjectToConfig registers projectName in the shared config.json,
func AddProjectToConfig(appConfigPath, projectName string) error {
	data, err := os.ReadFile(appConfigPath)
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
	err = os.WriteFile(appConfigPath, updatedData, 0o644)
	if err != nil {
		return err
	}

	return nil
}

// RemoveProjectFromConfig removes projectName from the shared config.json,
func RemoveProjectFromConfig(appConfigPath, projectName string) error {
	data, err := os.ReadFile(appConfigPath)
	if err != nil {
		return err
	}

	// Unmarshal the JSON config
	var config models.ProjectsConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}

	// Remove the project from the list if it exists
	updatedProjects := make([]string, 0, len(config.Projects))
	found := false
	for _, existingProject := range config.Projects {
		if existingProject == projectName {
			found = true
			continue
		}
		updatedProjects = append(updatedProjects, existingProject)
	}

	if !found {
		// project doesn't exist, nothing to do
		return nil
	}

	config.Projects = updatedProjects

	// Marshal the updated config back to JSON
	updatedData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	// Write the updated JSON
	err = os.WriteFile(appConfigPath, updatedData, 0o644)
	if err != nil {
		return err
	}

	return nil
}

func ChangeActiveProjectInConfig(appConfigPath string, projectName string) error {
	data, err := os.ReadFile(appConfigPath)
	if err != nil {
		return err
	}

	// Unmarshal the JSON config
	var config models.ProjectsConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}

	// Make sure projectName is within the config
	found := false
	for _, existingProject := range config.Projects {
		if existingProject == projectName {
			found = true
			break
		}
	}
	if !found {
		// project doesn't exist, nothing to do
		return nil
	}

	// Switch the active project in the config with the new project
	config.ActiveProject = projectName

	// Marshal the updated config back to JSON
	updatedData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	// Write the updated JSON
	err = os.WriteFile(appConfigPath, updatedData, 0o644)
	if err != nil {
		return err
	}

	return nil
}
