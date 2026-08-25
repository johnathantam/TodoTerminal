package main

import (
	// system packages
	"fmt"
	"os"

	// local packages
	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/app"
	"github.com/johnathantam/TodoTerminal/internal/commands"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		color.New(color.FgRed).Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: todo <command>")
	}

	commandName := args[0]
	commandArguments := args[1:]

	if commandName == "init" {
		return commands.Init(commandArguments)
	}

	appContext, err := app.LoadAppContext()
	if err != nil {
		return err
	}

	switch commandName {
	// case "add-task":
	//     return commands.AddTask(appContext, commandArgs)

	case "add-project":
		return commands.AddProject(appContext, commandArguments)
	case "remove-project":
		return commands.RemoveProject(appContext, commandArguments)

	default:
		return fmt.Errorf("unknown command: %q", commandName)
	}
}
