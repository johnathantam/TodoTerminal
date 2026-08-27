package commands

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	titleColor   = color.New(color.Bold)
	sectionColor = color.New(color.Bold)
	commandColor = color.New(color.FgCyan)
)

func printCommand(command, description string) {
	const commandWidth = 35

	fmt.Printf("  ")
	commandColor.Printf("%-*s", commandWidth, command)
	fmt.Printf("%s\n", description)
}

func Help() error {
	titleColor.Println("TodoTerminal")

	fmt.Println()
	fmt.Println("Usage:")
	fmt.Printf("  ")
	commandColor.Printf("todo")
	fmt.Println(" <command> [arguments]")

	fmt.Println()
	sectionColor.Println("Commands:")

	printCommand(
		"init <project-name>",
		"Initialize a new project",
	)

	fmt.Println()
	sectionColor.Println("Project commands:")

	printCommand(
		"project add <project-name>",
		"Add a project",
	)

	printCommand(
		"project remove <project-name>",
		"Remove a project",
	)

	printCommand(
		"project switch <project-name>",
		"Switch the active project",
	)

	printCommand(
		"project list",
		"List all projects",
	)

	fmt.Println()
	sectionColor.Println("Task commands:")

	printCommand(
		"task add <title> [description]",
		"Add a task",
	)

	printCommand(
		"task remove <task-id>",
		"Remove a task",
	)

	printCommand(
		"task get <task-id>",
		"Get a task",
	)

	printCommand(
		"task list",
		"List all tasks",
	)

	printCommand(
		"task details <task-id>",
		"Show task details",
	)

	printCommand(
		"task status <task-id> [status]",
		"Change task status",
	)

	fmt.Println()

	return nil
}
