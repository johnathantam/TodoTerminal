package storage

import (
	// System packages
	"fmt"
	"time"

	"github.com/johnathantam/TodoTerminal/internal/storage/models"
)

func GetTaskInProject(projectDirectoryPath string, projectName string, taskID string) (models.TodoItem, error) {
	todoList, err := ReadProjectTodoList(projectDirectoryPath, projectName)
	if err != nil {
		return models.TodoItem{}, err
	}

	// loop through todoList to find the given task
	for _, taskInProject := range todoList.Todos {
		if taskInProject.ID == taskID {
			return taskInProject, nil
		}
	}

	return models.TodoItem{}, fmt.Errorf("Could not find task %s in %s", taskID, projectName)
}

func GetTasksInProject(projectDirectoryPath string, projectName string) (models.TodoList, error) {
	todoList, err := ReadProjectTodoList(projectDirectoryPath, projectName)
	if err != nil {
		return models.TodoList{}, err
	}

	return todoList, nil
}

func SetTaskDetailsInProject(projectDirectoryPath string, projectName string, taskID string, newTitle string, newDescription string) error {
	// Grab the todolist
	todoList, err := ReadProjectTodoList(projectDirectoryPath, projectName)
	if err != nil {
		return err
	}

	// Find and update the task details
	for index := range todoList.Todos {
		if todoList.Todos[index].ID != taskID {
			continue
		}

		todoList.Todos[index].Title = newTitle
		todoList.Todos[index].Description = newDescription

		if err := WriteProjectTodoList(projectDirectoryPath, projectName, todoList); err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("task %q does not exist", taskID)
}

func SetTaskStatusInProject(projectDirectoryPath string, projectName string, taskID string, newStatus models.TodoStatus) error {
	// Check if the status is one of the set statuses
	if !newStatus.IsValid() {
		return fmt.Errorf("Can't change task status to %s. Must be pending, in progress, or completed", newStatus)
	}

	// Grab the todolist
	todoList, err := ReadProjectTodoList(projectDirectoryPath, projectName)
	if err != nil {
		return err
	}

	// Find and update the task status
	for index := range todoList.Todos {
		if todoList.Todos[index].ID != taskID {
			continue
		}

		todoList.Todos[index].Status = newStatus

		if err := WriteProjectTodoList(projectDirectoryPath, projectName, todoList); err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("task %q does not exist", taskID)
}

func ClearTasksInProject(projectDirectoryPath string, projectName string, status *models.TodoStatus) error {
	if status != nil && !status.IsValid() {
		return fmt.Errorf("invalid task status %q", *status)
	}

	// find todo list
	todoList, err := ReadProjectTodoList(projectDirectoryPath, projectName)
	if err != nil {
		return err
	}

	// Clear every task in status or all of them
	if status == nil {
		todoList.Todos = []models.TodoItem{}
	} else {
		remainingTasks := make([]models.TodoItem, 0)
		for _, task := range todoList.Todos {
			if task.Status != *status {
				remainingTasks = append(remainingTasks, task)
			}
		}

		todoList.Todos = remainingTasks
	}

	// Write to json
	err = WriteProjectTodoList(projectDirectoryPath, projectName, todoList)
	if err != nil {
		return err
	}

	return nil
}

func CreateTaskStructure(projectDirectoryPath string, projectName string, taskTitle string, taskDescription string) error {
	// find todo list
	todoList, err := ReadProjectTodoList(projectDirectoryPath, projectName)
	if err != nil {
		return err
	}

	// create new task
	newTask := models.TodoItem{
		ID:          fmt.Sprintf("%d", time.Now().UnixNano()),
		Title:       taskTitle,
		Description: taskDescription,
		Status:      models.TodoStatusPending,
	}

	// add task to the todolist
	todoList.Todos = append(todoList.Todos, newTask)

	// update the task list
	err = WriteProjectTodoList(projectDirectoryPath, projectName, todoList)
	if err != nil {
		return err
	}

	return nil
}

func RemoveTaskStructure(projectDirectoryPath string, projectName string, taskId string) error {
	// find todo list
	todoList, err := ReadProjectTodoList(projectDirectoryPath, projectName)
	if err != nil {
		return err
	}

	// loop through todo list and check for given id
	found := false
	updatedTasks := make([]models.TodoItem, 0, len(todoList.Todos))
	for _, task := range todoList.Todos {
		if task.ID == taskId {
			found = true
			continue
		}
		updatedTasks = append(updatedTasks, task)
	}
	if !found {
		return fmt.Errorf("task %q does not exist", taskId)
	}

	todoList.Todos = updatedTasks

	// write back to json to update
	err = WriteProjectTodoList(projectDirectoryPath, projectName, todoList)
	if err != nil {
		return err
	}

	return nil
}
