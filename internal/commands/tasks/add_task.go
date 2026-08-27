package tasks

import (
	// System packages
	"fmt"

	// Local packages
	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/app"
	"github.com/johnathantam/TodoTerminal/internal/storage"
)

func AddTask(appContext app.AppContext, commandArguments []string) error {
	if len(commandArguments) < 1 || len(commandArguments) > 2 {
		return fmt.Errorf("usage: todo task add <title> [description]")
	}

	// Grab task info
	taskTitle := commandArguments[0]
	taskDescription := ""
	if len(commandArguments) == 2 {
		taskDescription = commandArguments[1]
	}

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

	// Add the task
	err = storage.CreateTaskStructure(projectDirectoryPath, projectName, taskTitle, taskDescription)
	if err != nil {
		return err
	}

	color.Green("New task '%s' was created in '%s'", taskTitle, projectName)

	return nil
}
