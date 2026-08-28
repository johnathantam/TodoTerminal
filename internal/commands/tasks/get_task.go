package tasks

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/app"
	"github.com/johnathantam/TodoTerminal/internal/storage"
)

func GetTask(appContext app.AppContext, commandArguments []string) error {
	if len(commandArguments) != 1 {
		return fmt.Errorf("usage: todo task get <taskID>")
	}

	// get the taskID
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

	// fetch the task
	task, err := storage.GetTaskInProject(projectDirectoryPath, projectName, taskID)
	if err != nil {
		return err
	}

	// view the task
	color.Green("ID: %s\n", task.ID)
	color.Green("Title: %s\n", task.Title)
	color.Green("Description: %s\n", task.Description)
	color.Green("Status: %s\n", task.Status)
	return nil
}
