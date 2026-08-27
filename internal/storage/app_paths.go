package storage

// AppPaths holds the resolved filesystem locations that make up the
// TodoTerminal app's on-disk structure.
type AppPaths struct {
	AppDirectoryPath         string
	AppProjectsDirectoryPath string
	AppConfigPath            string
}
