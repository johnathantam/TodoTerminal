package main

import (
	// system packages
	"fmt"
	"os"

	// local packages
	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: todo <command>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "init":
		if err := commands.Init(os.Args[2:]); err != nil {
			color.New(color.FgRed).Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
