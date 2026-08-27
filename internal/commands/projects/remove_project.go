package projects

import (
	// system packages
	"fmt"

	// local packages
	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/app"
	"github.com/johnathantam/TodoTerminal/internal/storage"
)

func RemoveProject(appContext app.AppContext, commandArguments []string) error {
	if len(commandArguments) != 1 {
		return fmt.Errorf("usage: todo remove-project <project-name>")
	}

	projectName := commandArguments[0]

	// Make sure the project to remove isn't the current active project
	currentActiveProjectName, err := storage.GetActiveProjectInConfig(appContext.AppPaths.AppConfigPath)
	if err != nil {
		return err
	}
	if currentActiveProjectName == projectName {
		return fmt.Errorf("Cannot remove the active project")
	}

	// Remove the project from the projects folder
	err = storage.RemoveProjectStructure(appContext.AppPaths.AppProjectsDirectoryPath, projectName)
	if err != nil {
		return err
	}

	// Remove the project from the config
	err = storage.RemoveProjectFromConfig(appContext.AppPaths.AppConfigPath, projectName)
	if err != nil {
		return err
	}

	color.Green("Project %s has been deleted", projectName)

	return nil
}
