package cli

import (
	"fmt"

	"github.com/johnathantam/TodoTerminal/internal/app"
	"github.com/johnathantam/TodoTerminal/internal/commands/projects"
	"github.com/johnathantam/TodoTerminal/internal/commands/tasks"
)

type AppCommand func(app.AppContext, []string) error

var projectCommands = map[string]AppCommand{
	"add":     projects.AddProject,
	"remove":  projects.RemoveProject,
	"switch":  projects.SwitchProject,
	"current": projects.GetCurrentProject,
	"list":    projects.GetProjectList,
}

var taskCommands = map[string]AppCommand{
	"add":     tasks.AddTask,
	"remove":  tasks.RemoveTask,
	"get":     tasks.GetTask,
	"list":    tasks.GetTaskList,
	"details": tasks.SetTaskDetails,
	"status":  tasks.SetTaskStatus,
	"clear":   tasks.ClearTasks,
}

func AppCommandDispatch(appContext app.AppContext, commandArguments []string) error {
	if len(commandArguments) < 2 {
		return fmt.Errorf("usage: todo <project|task> <command> [arguments]")
	}

	commandGroup := commandArguments[0]
	commandName := commandArguments[1]
	commandArgs := commandArguments[2:]

	var commands map[string]AppCommand

	switch commandGroup {
	case "project":
		commands = projectCommands
	case "task":
		commands = taskCommands
	default:
		return fmt.Errorf("unknown command group: %q", commandGroup)
	}

	command, exists := commands[commandName]
	if !exists {
		return fmt.Errorf("unknown %s command: %q", commandGroup, commandName)
	}

	return command(appContext, commandArgs)
}
