package projects

import (
	// System packages
	"fmt"

	// Local packages
	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/app"
	"github.com/johnathantam/TodoTerminal/internal/storage"
)

func GetProjectList(appContext app.AppContext, commandArguments []string) error {
	if len(commandArguments) != 0 {
		return fmt.Errorf("usage: todo add-project")
	}

	projectNames, err := storage.GetProjectsInConfig(appContext.AppPaths.AppConfigPath)
	if err != nil {
		return err
	}

	currentActiveProjectName, err := storage.GetActiveProjectInConfig(appContext.AppPaths.AppConfigPath)
	if err != nil {
		return err
	}

	for _, projectName := range projectNames {
		if projectName == currentActiveProjectName {
			color.Green("* %s", projectName)
		} else {
			fmt.Printf("  %s\n", projectName)
		}
	}

	return nil
}
