package storage

import (
	// system packages
	"errors"
	"os"
	"path/filepath"
)

// FindAppPaths searches from the current working directory upward
// until it finds a valid TodoTerminal application structure.
func FindLocationOfAppPaths() (AppPaths, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return AppPaths{}, err
	}

	currentPath := cwd
	for {
		candidatePath := filepath.Join(currentPath, ".todoterminal")

		isValid, err := IsAppStructureValid(candidatePath)
		if err != nil {
			return AppPaths{}, err
		}
		if isValid {
			return AppPaths{
				AppDirectoryPath:         candidatePath,
				AppProjectsDirectoryPath: filepath.Join(candidatePath, "projects"),
				AppConfigPath:            filepath.Join(candidatePath, "config.json"),
			}, nil
		}

		parentPath := filepath.Dir(currentPath)
		// We have reached the filesystem root.
		if parentPath == currentPath {
			break
		}
		currentPath = parentPath
	}

	return AppPaths{}, errors.New("TodoTerminal project not found")
}

// IsAppStructureValid reports whether the expected TodoTerminal directory
// structure (root dir, projects dir, config.json) exists on disk, without
// creating or modifying anything.
func IsAppStructureValid(rootPath string) (bool, error) {
	// Check root directory.
	rootInfo, err := os.Stat(rootPath)

	// make sure root directory exists
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	// make sure root directory is a directory and not a file
	if !rootInfo.IsDir() {
		return false, nil
	}

	// Check projects directory.
	projectsPath := filepath.Join(rootPath, "projects")
	projectsInfo, err := os.Stat(projectsPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	if !projectsInfo.IsDir() {
		return false, nil
	}

	// Check config file.
	configPath := filepath.Join(rootPath, "config.json")
	configInfo, err := os.Stat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	if configInfo.IsDir() {
		return false, nil
	}

	return true, nil
}

// ProjectExists reports whether a project directory exists inside the given
// projects directory.
func ProjectExists(projectsDirectoryPath, projectName string) (bool, error) {
	projectDirectoryPath := filepath.Join(projectsDirectoryPath, projectName)

	projectInfo, err := os.Stat(projectDirectoryPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return projectInfo.IsDir(), nil
}
