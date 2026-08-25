package commands

import (
	// system packages
	"fmt"

	// local packages
	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/storage"
)

func Init(commandArguments []string) error {
	if len(commandArguments) != 1 {
		return fmt.Errorf("usage: todo init <project-name>")
	}

	projectName := commandArguments[0]

	appPaths, err := storage.CreateAppStructure()
	if err != nil {
		return err
	}

	projectPath, err := storage.CreateProjectStructure(appPaths.AppProjectsDirectoryPath, projectName)
	if err != nil {
		return err
	}

	err = storage.AddProjectToConfig(appPaths.AppConfigPath, projectName)
	if err != nil {
		return err
	}

	err = storage.ChangeActiveProjectInConfig(appPaths.AppConfigPath, projectName)
	if err != nil {
		return err
	}

	color.Green("Project storage initialized at: %s", appPaths.AppDirectoryPath)
	color.Green("Project %q created at: %s", projectName, projectPath)

	return nil
}
