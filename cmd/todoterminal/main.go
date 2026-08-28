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
	case "cleanup", "destroy":
		return commands.Destroy(commandArguments)
	}

	// Route all other commands
	appContext, err := app.LoadAppContext()
	if err != nil {
		return err
	}

	return cli.AppCommandDispatch(appContext, args)
}
