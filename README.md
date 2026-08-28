# TodoTerminal

A fast, no-frills todo and project manager for your terminal — organize tasks by project, switch contexts like git branches, and never leave the command line.

## Features

- 📁 **Project-based organization** — group tasks into separate projects, switch between them instantly
- ✅ **Full task lifecycle** — add, edit, list, filter by status, and clear tasks in bulk
- 🔍 **Git-style discovery** — automatically finds your `.todoterminal` config by walking up from your current directory, just like `.git`
- 🎨 **Colored terminal output** — clean, readable output with highlighted active projects
- 📦 **Single static binary** — no runtime dependencies, install in seconds

## Installation

### macOS / Linux

```bash
curl -fsSL https://raw.githubusercontent.com/johnathantam/TodoTerminal/main/install.sh | sh
```

### Windows (PowerShell)

```powershell
irm https://raw.githubusercontent.com/johnathantam/TodoTerminal/main/install.ps1 | iex
```

### Via Go

If you already have Go installed:

```bash
go install github.com/johnathantam/TodoTerminal/cmd/todoterminal@latest
```

### Manual download

Prebuilt binaries for macOS, Linux, and Windows (amd64/arm64) are available on the [Releases page](https://github.com/johnathantam/TodoTerminal/releases).

## Uninstall

**macOS / Linux**

```bash
curl -fsSL https://raw.githubusercontent.com/johnathantam/TodoTerminal/main/uninstall.sh | sh
```

**Windows (PowerShell)**

```powershell
irm https://raw.githubusercontent.com/johnathantam/TodoTerminal/main/uninstall.ps1 | iex
```

## Quick Start

```bash
# Initialize a new project in the current directory
todo init my-first-project

# Add a task
todo task add "Write the README" "Make it look nice"

# List all tasks
todo task list

# Mark a task as in_progress
todo task status <task-id> in_progress

# See all your projects
todo project list
```

## Usage

```
todo <command> [arguments]
```

### Commands

| Command | Description |
|---|---|
| `init <project-name>` | Initialize a new project |
| `destroy` (alias: `cleanup`) | Delete the `.todoterminal` directory and all its projects/tasks (prompts for confirmation) |

### Project Commands

| Command | Description |
|---|---|
| `project add <project-name>` | Add a project |
| `project remove <project-name>` | Remove a project |
| `project switch <project-name>` | Switch the active project |
| `project current` | See the current active project |
| `project list` | List all projects |

### Task Commands

| Command | Description |
|---|---|
| `task add <title> [description]` | Add a task |
| `task remove <task-id>` | Remove a task |
| `task get <task-id>` | Get a task |
| `task list` | List all tasks |
| `task details <task-id> <new-title> [new-description]` | Update a task's title and description |
| `task status <task-id> [status]` | Change task status |
| `task clear [status]` | Clear all tasks, or only those of a specific status |

Run `todo help` at any time to see this list from the CLI itself.

## How it works

TodoTerminal stores its data in a `.todoterminal/` directory, similar to how git uses `.git/`. When you run a command, it searches the current directory and walks upward through parent directories until it finds one — so you can run `todo` commands from any subdirectory of your project.

```
.todoterminal/
├── config.json          # tracks all registered projects + the active one
└── projects/
    └── my-project/
        ├── my-project.json              # project metadata
        └── my-project-todo-list.json    # tasks
```

## Building from source

```bash
git clone https://github.com/johnathantam/TodoTerminal.git
cd TodoTerminal
go build -o todo ./cmd/todoterminal
```

## Releasing

Releases are built and published automatically via [GoReleaser](https://goreleaser.com/) and GitHub Actions whenever a version tag is pushed:

```bash
git tag v1.0.1
git push origin v1.0.1
```
