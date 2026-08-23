package storage

import (
	// system packages
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/johnathantam/TodoTerminal/internal/storage/models"
)

// CreateProjectsTerminalRoot creates the root directory for storing TodoTerminal projects in the current working directory.
func CreateProjectsTerminalRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	rootPath := filepath.Join(cwd, "TodoTerminal")
	if err := os.MkdirAll(rootPath, 0o755); err != nil {
		return "", err
	}

	return rootPath, nil
}

// CreateProjectsTerminalRootConfig creates a configuration file for the TodoTerminal projects in the root directory.
func CreateProjectsTerminalRootConfig(rootPath string) (string, error) {
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

func InitializeRoot() (string, error) {
	rootPath, err := CreateProjectsTerminalRoot()
	if err != nil {
		return "", err
	}

	if _, err := CreateProjectsTerminalRootConfig(rootPath); err != nil {
		return "", err
	}

	return rootPath, nil
}
