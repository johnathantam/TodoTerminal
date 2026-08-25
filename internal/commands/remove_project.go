package commands

import (
	// system packages
	"fmt"

	// local packages

	"github.com/johnathantam/TodoTerminal/internal/app"
)

func RemoveProject(appContext app.AppContext, commandArguments []string) error {
	if len(commandArguments) != 1 {
		return fmt.Errorf("usage: todo remove-project <project-name>")
	}

	return nil
}
