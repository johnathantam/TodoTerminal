package tasks

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/app"
	"github.com/johnathantam/TodoTerminal/internal/storage"
)

func SetTaskDetails(appContext app.AppContext, commandArguments []string) error {
	if len(commandArguments) != 3 {
		return fmt.Errorf("usage: todo task details <taskID> [newTaskTitle] (newTaskDescription)")
	}

	// Grab the parameters
	taskID := commandArguments[0]
	newTaskTitle := commandArguments[1]
	newTaskDescription := commandArguments[2]

	// Grab the current project
	currentProjectName, err := storage.GetActiveProjectInConfig(appContext.AppPaths.AppConfigPath)
	if err != nil {
		return err
	}

	currentProjectDirectoryPath, err := appContext.FindProjectDirectoryPath(currentProjectName)
	if err != nil {
		return err
	}

	// Update the task
	err = storage.SetTaskDetailsInProject(currentProjectDirectoryPath, currentProjectName, taskID, newTaskTitle, newTaskDescription)
	if err != nil {
		return err
	}

	color.Green("Task '%s' has been updated in %s with new details", taskID, currentProjectName)

	return nil
}
