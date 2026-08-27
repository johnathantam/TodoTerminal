package commands

import "fmt"

func Help() error {
	fmt.Println(`TodoTerminal commands:

  init <project-name>              Initialize the app
  add-project <project-name>       Create a project
  remove-project <project-name>    Remove a project
  switch-project <project-name>    Switch active project
  get-projects                     List projects

  add-task <title> [description]   Add a task
  remove-task <task-id>            Remove a task
  get-task <task-id>               Show a task
  get-tasks                        List tasks
  set-task-details <task-id>       Update task details
  set-task-status <task-id> [status] Update task status
	`)

	return nil
}
