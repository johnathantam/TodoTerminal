package tasks

import (
	// System packages
	"fmt"

	// Local packages
	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/app"
	"github.com/johnathantam/TodoTerminal/internal/storage"
	"github.com/johnathantam/TodoTerminal/internal/storage/models"
)

var (
	pendingColor    = color.New(color.FgYellow)
	inProgressColor = color.New(color.FgCyan)
	completedColor  = color.New(color.FgGreen)

	pendingHeader    = color.New(color.FgYellow, color.Bold)
	inProgressHeader = color.New(color.FgCyan, color.Bold)
	completedHeader  = color.New(color.FgGreen, color.Bold)
)

func printTodoSection(
	title string,
	todos []models.TodoItem,
	itemColor *color.Color,
	headerColor *color.Color,
	symbol string,
) {
	fmt.Printf("\n%s\n", headerColor.Sprint(title))
	fmt.Println("────────────────────────")

	if len(todos) == 0 {
		fmt.Println("No tasks here")
		return
	}

	for _, task := range todos {
		fmt.Printf("%s %s: %s\n",
			itemColor.Sprint(symbol),
			task.ID,
			task.Title,
		)
	}
}

func GetTaskList(appContext app.AppContext) error {
	// Get active project
	currentActiveProjectName, err := storage.GetActiveProjectInConfig(appContext.AppPaths.AppConfigPath)
	if err != nil {
		return err
	}

	// Grab the project directory
	projectDirectoryPath, err := appContext.FindProjectDirectoryPath(currentActiveProjectName)
	if err != nil {
		return err
	}

	tasks, err := storage.GetTasksInProject(projectDirectoryPath, currentActiveProjectName)
	if err != nil {
		return err
	}

	pending := []models.TodoItem{}
	inProgress := []models.TodoItem{}
	completed := []models.TodoItem{}
	for _, task := range tasks.Todos {
		switch task.Status {
		case models.TodoStatusPending:
			pending = append(pending, task)

		case models.TodoStatusInProgress:
			inProgress = append(inProgress, task)

		case models.TodoStatusCompleted:
			completed = append(completed, task)
		}
	}

	printTodoSection(
		"Pending",
		pending,
		pendingColor,
		pendingHeader,
		"[ ]",
	)

	printTodoSection(
		"In Progress",
		inProgress,
		inProgressColor,
		inProgressHeader,
		"[>]",
	)

	printTodoSection(
		"Completed",
		completed,
		completedColor,
		completedHeader,
		"[x]",
	)

	fmt.Println("\n- Add tasks using the `todo add-task <taskTitle> [taskDescription]` command")

	return nil
}
