package storage

import (
	// system packages
	"os"
	"path/filepath"
)

func isCurrentWorkingDirectoryInitialized() (bool, error) {
	// grab working directory
	cwd, err := os.Getwd()
	if err != nil {
		return false, err
	}

	// check whether there is a folder named TodoTerminal in the current working directory
	rootPath := filepath.Join(cwd, "TodoTerminal")
	pathInfo, err := os.Stat(rootPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	// check whether there is a config.json file in the TodoTerminal folder
	configPath := filepath.Join(rootPath, "config.json")
	if _, err := os.Stat(configPath); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return pathInfo.IsDir(), nil
}
