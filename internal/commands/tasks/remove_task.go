package tasks

import (
	// System packages
	"fmt"

	// Local packages
	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/app"
	"github.com/johnathantam/TodoTerminal/internal/storage"
)

func RemoveTask(appContext app.AppContext, commandArguments []string) error {
	if len(commandArguments) != 1 {
		return fmt.Errorf("usage: todo task remove [taskID]")
	}

	// Grab task info
	taskID := commandArguments[0]

	// Grab the current project
	projectName, err := storage.GetActiveProjectInConfig(appContext.AppPaths.AppConfigPath)
	if err != nil {
		return err
	}

	// Grab the project directory
	projectDirectoryPath, err := appContext.FindProjectDirectoryPath(projectName)
	if err != nil {
		return err
	}

	// Remove the task
	err = storage.RemoveTaskStructure(projectDirectoryPath, projectName, taskID)
	if err != nil {
		return err
	}

	color.Green("Removed task '%s' was created in '%s'", taskID, projectName)

	return nil
}
