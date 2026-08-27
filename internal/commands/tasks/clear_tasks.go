package tasks

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/app"
	"github.com/johnathantam/TodoTerminal/internal/storage"
	"github.com/johnathantam/TodoTerminal/internal/storage/models"
)

func ClearTasks(appContext app.AppContext, commandArguments []string) error {
	if len(commandArguments) > 1 {
		return fmt.Errorf("usage: todo task clear [status]")
	}

	var status *models.TodoStatus
	if len(commandArguments) == 1 {
		taskStatus := models.TodoStatus(commandArguments[0])

		if !taskStatus.IsValid() {
			return fmt.Errorf("invalid task status %q. Use either pending, in_progress, or completed", taskStatus)
		}

		status = &taskStatus
	}

	projectName, err := storage.GetActiveProjectInConfig(appContext.AppPaths.AppConfigPath)
	if err != nil {
		return err
	}

	projectDirectoryPath, err := appContext.FindProjectDirectoryPath(projectName)
	if err != nil {
		return err
	}

	if err := storage.ClearTasksInProject(projectDirectoryPath, projectName, status); err != nil {
		return err
	}

	color.Green("Tasks cleared")

	return nil
}
