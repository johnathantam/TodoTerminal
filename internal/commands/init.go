package commands

import (
	// system packages
	"fmt"

	// local packages
	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/storage"
)

func Init(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: todo init <project-name>")
	}

	projectName := args[0]

	rootPath, err := storage.InitializeRoot()
	if err != nil {
		return fmt.Errorf("initializing project storage: %w", err)
	}

	projectPath, err := storage.CreateProject(rootPath, projectName)
	if err != nil {
		return fmt.Errorf("project %q was not created: %w", projectName, err)
	}

	color.Green("Project storage initialized at: %s", rootPath)
	color.Green("Project %q created at: %s", projectName, projectPath)

	return nil
}
