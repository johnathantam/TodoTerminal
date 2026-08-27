package projects

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/app"
	"github.com/johnathantam/TodoTerminal/internal/storage"
)

func GetCurrentProject(appContext app.AppContext, commandArguments []string) error {
	if len(commandArguments) != 0 {
		return fmt.Errorf("useage: todo project current")
	}

	// Grab the current active project
	currentActiveProjectName, err := storage.GetActiveProjectInConfig(appContext.AppPaths.AppConfigPath)
	if err != nil {
		return err
	}

	// Print out the project
	color.Green("%s\n", currentActiveProjectName)

	return nil
}
