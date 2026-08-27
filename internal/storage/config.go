package storage

import (
	// system packages
	"encoding/json"
	"fmt"
	"os"

	// local packages
	"github.com/johnathantam/TodoTerminal/internal/storage/models"
)

func readProjectsConfig(appConfigPath string) (models.ProjectsConfig, error) {
	data, err := os.ReadFile(appConfigPath)
	if err != nil {
		return models.ProjectsConfig{}, err
	}

	var config models.ProjectsConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return models.ProjectsConfig{}, err
	}

	return config, nil
}

func writeProjectsConfig(appConfigPath string, config models.ProjectsConfig) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(appConfigPath, data, 0o644); err != nil {
		return err
	}

	return nil
}

func GetActiveProjectInConfig(appConfigPath string) (string, error) {
	config, err := readProjectsConfig(appConfigPath)
	if err != nil {
		return "", err
	}

	// Check the active project and return it
	return config.ActiveProject, nil
}

func GetProjectsInConfig(appConfigPath string) ([]string, error) {
	config, err := readProjectsConfig(appConfigPath)
	if err != nil {
		return nil, err
	}

	return config.Projects, nil
}

// AddProjectToConfig registers projectName in the shared config.json,
func AddProjectToConfig(appConfigPath, projectName string) error {
	config, err := readProjectsConfig(appConfigPath)
	if err != nil {
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

	err = writeProjectsConfig(appConfigPath, config)
	if err != nil {
		return err
	}

	return nil
}

// RemoveProjectFromConfig removes projectName from the shared config.json,
func RemoveProjectFromConfig(appConfigPath, projectName string) error {
	config, err := readProjectsConfig(appConfigPath)
	if err != nil {
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
		return fmt.Errorf("Project doesn't exist in the config")
	}

	config.Projects = updatedProjects

	err = writeProjectsConfig(appConfigPath, config)
	if err != nil {
		return err
	}

	return nil
}

func ChangeActiveProjectInConfig(appConfigPath string, projectName string) error {
	config, err := readProjectsConfig(appConfigPath)
	if err != nil {
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
		return fmt.Errorf("Can't switch project. Project %s doesn't exist", projectName)
	}

	// Switch the active project in the config with the new project
	config.ActiveProject = projectName

	err = writeProjectsConfig(appConfigPath, config)
	if err != nil {
		return err
	}

	return nil
}
