package projects

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/app"
	"github.com/johnathantam/TodoTerminal/internal/storage"
)

func SwitchProject(appContext app.AppContext, commandArguments []string) error {
	if len(commandArguments) != 1 {
		return fmt.Errorf("usage: todo project switch <project-name>")
	}

	// Switch active project
	projectName := commandArguments[0]
	err := storage.ChangeActiveProjectInConfig(appContext.AppPaths.AppConfigPath, projectName)
	if err != nil {
		return err
	}

	color.Green("Current project has been switched to %s", projectName)

	return nil
}
