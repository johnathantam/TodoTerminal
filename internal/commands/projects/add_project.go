package projects

import (
	// system packages
	"fmt"

	// local packages
	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/app"
	"github.com/johnathantam/TodoTerminal/internal/storage"
)

func AddProject(appContext app.AppContext, commandArguments []string) error {
	if len(commandArguments) != 1 {
		return fmt.Errorf("usage: todo project add <project-name>")
	}

	// Grab project name
	projectName := commandArguments[0]
	// Check if the name is valid
	if _, err := storage.IsProjectNameValid(projectName); err != nil {
		return err
	}

	// Create the project
	_, err := storage.CreateProjectStructure(appContext.AppPaths.AppProjectsDirectoryPath, projectName)
	if err != nil {
		return err
	}

	err = storage.AddProjectToConfig(appContext.AppPaths.AppConfigPath, projectName)
	if err != nil {
		return err
	}

	err = storage.ChangeActiveProjectInConfig(appContext.AppPaths.AppConfigPath, projectName)
	if err != nil {
		return err
	}

	color.Green("New project %s created", projectName)

	return nil
}
