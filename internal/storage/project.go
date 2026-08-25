package storage

import (
	// system packages
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	// local packages
	"github.com/johnathantam/TodoTerminal/internal/storage/models"
)

// EnsureProjectDir ensures a project's directory exists inside the given
// projects folder, creating it if needed.
func EnsureProjectDir(projectsDirPath string, projectName string) (string, error) {
	projectPath := filepath.Join(projectsDirPath, projectName)
	if err := os.MkdirAll(projectPath, 0o755); err != nil {
		return "", err
	}

	return projectPath, nil
}

// EnsureProjectMetadataFile ensures a project's metadata file exists,
// creating it with default values if it doesn't.
func EnsureProjectMetadataFile(projectDirPath string, projectName string) (string, error) {
	metadataPath := filepath.Join(projectDirPath, projectName+".json")
	if _, err := os.Stat(metadataPath); err == nil {
		return metadataPath, nil // already exists
	} else if !os.IsNotExist(err) {
		return "", err // some other filesystem error occurred
	}

	metadata := models.ProjectMetadata{
		Name:      projectName,
		CreatedAt: time.Now(),
	}

	data, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return "", err
	}

	if err := os.WriteFile(metadataPath, data, 0o644); err != nil {
		return "", err
	}

	return metadataPath, nil
}

// EnsureProjectTodoListFile ensures a project's todo list file exists,
// creating it empty if it doesn't.
func EnsureProjectTodoListFile(projectDirPath string, projectName string) (string, error) {
	todoListPath := filepath.Join(projectDirPath, projectName+"-todo-list.json")

	if _, err := os.Stat(todoListPath); err == nil {
		return todoListPath, nil // already exists
	} else if !os.IsNotExist(err) {
		return "", err // some other filesystem error occurred
	}

	todoList := models.TodoList{
		Todos: []models.TodoItem{},
	}

	data, err := json.MarshalIndent(todoList, "", "  ")
	if err != nil {
		return "", err
	}

	if err := os.WriteFile(todoListPath, data, 0o644); err != nil {
		return "", err
	}

	return todoListPath, nil
}

// CreateProjectStructure creates a new project — directory, metadata file,
// todo list file, and a config.json entry — and registers it as a new
// project. Returns an error if a project with this name already exists.
func CreateProjectStructure(projectsDirPath string, projectName string) (string, error) {
	projectDirectoryPath := filepath.Join(projectsDirPath, projectName)

	projectExists, err := ProjectExists(projectsDirPath, projectName)
	if err != nil {
		return "", err
	}
	if projectExists {
		return "", fmt.Errorf("project %q already exists", projectName)
	}

	if _, err := EnsureProjectDir(projectsDirPath, projectName); err != nil {
		return "", err
	}

	if _, err := EnsureProjectMetadataFile(projectDirectoryPath, projectName); err != nil {
		return "", err
	}

	if _, err := EnsureProjectTodoListFile(projectDirectoryPath, projectName); err != nil {
		return "", err
	}

	return projectDirectoryPath, nil
}

func RemoveProjectStructure(projectsDirPath string, projectName string) error {
	projectDirectoryPath := filepath.Join(projectsDirPath, projectName)

	// Make sure project exists
	projectExists, err := ProjectExists(projectsDirPath, projectName)
	if err != nil {
		return err
	}
	if !projectExists {
		return fmt.Errorf("project %q does not exist", projectName)
	}

	// Remove the project from the project directory path
	err = os.RemoveAll(projectDirectoryPath)
	if err != nil {
		return fmt.Errorf("failed to remove project %q: %w", projectName, err)
	}

	return nil
}
