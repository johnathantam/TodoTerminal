package tasks

import (
	"fmt"

	"charm.land/huh/v2"
	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/app"
	"github.com/johnathantam/TodoTerminal/internal/storage"
	"github.com/johnathantam/TodoTerminal/internal/storage/models"
)

func selectTaskStatusByGUI() (models.TodoStatus, error) {
	var status models.TodoStatus

	err := huh.NewSelect[models.TodoStatus]().
		Title("Select task status").
		Options(
			huh.NewOption("Pending", models.TodoStatusPending),
			huh.NewOption("In Progress", models.TodoStatusInProgress),
			huh.NewOption("Completed", models.TodoStatusCompleted),
		).
		Value(&status).
		Run()

	if err != nil {
		return "", err
	}

	return status, nil
}

func SetTaskStatus(appContext app.AppContext, commandArguments []string) error {
	if len(commandArguments) < 1 || len(commandArguments) > 2 {
		return fmt.Errorf("usage: todo task status <task-id> [status]")
	}

	// Get task ID
	taskID := commandArguments[0]

	// Get the new task status
	var newStatus models.TodoStatus
	var err error
	if len(commandArguments) == 2 {
		newStatus = models.TodoStatus(commandArguments[1])
		if !newStatus.IsValid() {
			return fmt.Errorf("invalid task status %q", newStatus)
		}
	} else {
		newStatus, err = selectTaskStatusByGUI()
		if err != nil {
			return err
		}
	}

	projectName, err := storage.GetActiveProjectInConfig(appContext.AppPaths.AppConfigPath)
	if err != nil {
		return err
	}

	projectDirectoryPath, err := appContext.FindProjectDirectoryPath(projectName)
	if err != nil {
		return err
	}

	if err := storage.SetTaskStatusInProject(projectDirectoryPath, projectName, taskID, newStatus); err != nil {
		return err
	}

	color.Green("Task %q changed to %q\n", taskID, newStatus)

	return nil

}
