package app

import (
	// System packages.
	"fmt"
	"path/filepath"

	// Local packages.
	"github.com/johnathantam/TodoTerminal/internal/storage"
)

// AppContext contains application-wide information and state
// needed by commands and other parts of the application.
type AppContext struct {
	AppPaths storage.AppPaths
}

// LoadAppContext finds and validates the TodoTerminal application
// structure and creates an application context from it.
func LoadAppContext() (AppContext, error) {
	// Find the TodoTerminal application paths.
	appPaths, err := storage.FindLocationOfAppPaths()
	if err != nil {
		return AppContext{}, err
	}

	// Create the application context now that the structure is valid.
	return AppContext{
		AppPaths: appPaths,
	}, nil
}

// ProjectExists checks whether a project with the given name
// exists within the application's projects directory.
func (appContext AppContext) ProjectExists(projectName string) (bool, error) {
	return storage.ProjectExists(appContext.AppPaths.AppProjectsDirectoryPath, projectName)
}

// RequireProjectExists ensures that a project with the given name
// exists. An error is returned if the project does not exist
// or cannot be checked.
func (appContext AppContext) RequireProjectExists(projectName string) error {
	projectExists, err := appContext.ProjectExists(projectName)
	if err != nil {
		return err
	}

	if !projectExists {
		return fmt.Errorf("project %q does not exist", projectName)
	}

	return nil
}

func (appContext AppContext) FindProjectDirectoryPath(projectName string) (string, error) {
	if err := appContext.RequireProjectExists(projectName); err != nil {
		return "", err
	}

	return filepath.Join(
		appContext.AppPaths.AppProjectsDirectoryPath,
		projectName,
	), nil
}
