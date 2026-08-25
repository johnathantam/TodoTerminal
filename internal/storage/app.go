package storage

import (
	// system packages
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/johnathantam/TodoTerminal/internal/storage/models"
)

// EnsureAppDirectory creates the root directory for storing TodoTerminal projects in the current working directory.
func EnsureAppDirectory() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	rootPath := filepath.Join(cwd, ".todoterminal")
	if err := os.MkdirAll(rootPath, 0o755); err != nil {
		return "", err
	}

	return rootPath, nil
}

func EnsureAppProjectsDirectory(rootPath string) (string, error) {
	projectsPath := filepath.Join(rootPath, "projects")
	if err := os.MkdirAll(projectsPath, 0o755); err != nil {
		return "", err
	}

	return projectsPath, nil
}

// EnsureAppConfig creates a configuration file for the TodoTerminal projects in the root directory.
func EnsureAppConfig(rootPath string) (string, error) {
	configPath := filepath.Join(rootPath, "config.json")

	if _, err := os.Stat(configPath); err == nil {
		// File already exists
		return configPath, nil
	} else if !os.IsNotExist(err) {
		// Some other filesystem error occurred
		return "", err
	}

	// Create project metadata
	metadata := models.ProjectsConfig{
		ActiveProject: "",
		Projects:      []string{},
	}

	// Marshal the metadata to JSON
	data, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return "", err
	}

	// Write the JSON data to a file
	err = os.WriteFile(configPath, data, 0o644)
	if err != nil {
		return "", err
	}

	return configPath, nil
}

// EnsureAppStructure ensures that the necessary directory structure and configuration file for the TodoTerminal application exist.
// It creates the root directory, projects directory, and configuration file if they do not already exist. It returns the path to the root directory and any error encountered during the process.
func CreateAppStructure() (AppPaths, error) {
	// Ensure there is a root folder to store our data
	appDirectoryPath, err := EnsureAppDirectory()
	if err != nil {
		return AppPaths{}, err
	}

	// Ensure there is a projects folder to store our data
	appProjectsDirectoryPath, err := EnsureAppProjectsDirectory(appDirectoryPath)
	if err != nil {
		return AppPaths{}, err
	}

	// Ensure there is a config.json file to store our data
	appConfigPath, err := EnsureAppConfig(appDirectoryPath)
	if err != nil {
		return AppPaths{}, err
	}

	return AppPaths{
		AppDirectoryPath:         appDirectoryPath,
		AppProjectsDirectoryPath: appProjectsDirectoryPath,
		AppConfigPath:            appConfigPath,
	}, nil
}
