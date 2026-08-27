package tasks

import (
	"fmt"

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
	fmt.Printf("ID: %s\n", task.ID)
	fmt.Printf("Title: %s\n", task.Title)
	fmt.Printf("Description: %s\n", task.Description)
	fmt.Printf("Status: %s\n", task.Status)

	return nil
}
