package storage

import (
	// system packages
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	// local packages
	"github.com/johnathantam/TodoTerminal/internal/storage/models"
)

// CreateStorageDirectory creates a project directory for storing application data in the user's configuration directory.
func CreateProjectStorageDirectory(rootPath, projectName string) (string, error) {
	projectPath := filepath.Join(rootPath, projectName)
	if err := os.MkdirAll(projectPath, 0o755); err != nil {
		return "", err
	}

	return projectPath, nil
}

func CreateProjectStorageMetadataFile(projectPath, projectName string) (string, error) {
	projectStorageMetadataPath := filepath.Join(projectPath, projectName+".json")

	if _, err := os.Stat(projectStorageMetadataPath); err == nil {
		// File already exists
		return "", nil
	} else if !os.IsNotExist(err) {
		// Some other filesystem error occurred
		return "", err
	}

	// Create project metadata
	metadata := models.ProjectMetadata{
		Name:      projectName,
		CreatedAt: time.Now(),
	}

	// Marshal the metadata to JSON
	data, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return "", err
	}

	// Write the JSON data to a file
	err = os.WriteFile(projectStorageMetadataPath, data, 0o644)
	if err != nil {
		return "", err
	}

	return projectStorageMetadataPath, nil
}

func CreateProjectTodoListMetadataFile(projectPath, projectName string) (string, error) {
	projectTodoListMetadataPath := filepath.Join(projectPath, projectName+"-todo-list.json")

	if _, err := os.Stat(projectTodoListMetadataPath); err == nil {
		// File already exists
		return "", nil
	} else if !os.IsNotExist(err) {
		// Some other filesystem error occurred
		return "", err
	}

	// Create an empty todo list
	todoList := models.TodoList{
		Todos: []models.TodoItem{},
	}

	// Marshal the todo list to JSON
	data, err := json.MarshalIndent(todoList, "", "  ")
	if err != nil {
		return "", err
	}

	// Write the JSON data to a file
	err = os.WriteFile(projectTodoListMetadataPath, data, 0o644)
	if err != nil {
		return "", err
	}

	return projectTodoListMetadataPath, nil
}

func CreateProject(rootPath, projectName string) (string, error) {
	projectPath, err := CreateProjectStorageDirectory(rootPath, projectName)
	if err != nil {
		return "", err
	}

	if _, err := CreateProjectStorageMetadataFile(projectPath, projectName); err != nil {
		return "", err
	}

	if _, err := CreateProjectTodoListMetadataFile(projectPath, projectName); err != nil {
		return "", err
	}

	if err := AddProjectToConfig(rootPath, projectName); err != nil {
		return "", err
	}

	return projectPath, nil
}
