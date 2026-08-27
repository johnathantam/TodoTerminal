package main

import (
	// system packages
	"fmt"
	"os"

	// local packages
	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/app"
	"github.com/johnathantam/TodoTerminal/internal/cli"
	"github.com/johnathantam/TodoTerminal/internal/commands"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		color.New(color.FgRed).Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// func run(args []string) error {
// 	if len(args) < 1 {
// 		return fmt.Errorf("usage: todo <command>")
// 	}

// 	commandName := args[0]
// 	commandArguments := args[1:]

// 	if commandName == "help" {
// 		return commands.Help()
// 	}

// 	if commandName == "init" {
// 		return commands.Init(commandArguments)
// 	}

// 	appContext, err := app.LoadAppContext()
// 	if err != nil {
// 		return err
// 	}

// 	switch commandName {

// 	// Commands related to projects
// 	case "add-project":
// 		return projectsCommands.AddProject(appContext, commandArguments)
// 	case "remove-project":
// 		return projectsCommands.RemoveProject(appContext, commandArguments)
// 	case "switch-project":
// 		return projectsCommands.SwitchProject(appContext, commandArguments)
// 	case "get-projects":
// 		return projectsCommands.GetProjectList(appContext, commandArguments)

// 	// Commands related to tasks
// 	case "add-task":
// 		return tasksCommands.AddTask(appContext, commandArguments)
// 	case "remove-task":
// 		return tasksCommands.RemoveTask(appContext, commandArguments)
// 	case "get-task":
// 		return tasksCommands.GetTask(appContext, commandArguments)
// 	case "get-tasks":
// 		return tasksCommands.GetTaskList(appContext, commandArguments)
// 	case "set-task-details":
// 		return tasksCommands.SetTaskDetails(appContext, commandArguments)
// 	case "set-task-status":
// 		return tasksCommands.SetTaskStatus(appContext, commandArguments)

// 	default:
// 		return fmt.Errorf("unknown command: %q", commandName)
// 	}
// }

func run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: todo <command>")
	}

	commandName := args[0]
	commandArguments := args[1:]

	// Check for essential commands first like help and init which
	// - Init sets a lot of crucial infrastructure so run it first
	switch commandName {
	case "help", "-h", "--help":
		return commands.Help()
	case "init":
		return commands.Init(commandArguments)
	}

	// Route all other commands
	appContext, err := app.LoadAppContext()
	if err != nil {
		return err
	}

	return cli.AppCommandDispatch(appContext, args)
}
