package bot

import "errors"

var (
	ErrorAbortTask     = errors.New("error Abort Task")
	ErrorAddTask       = errors.New("error Add Task")
	ErrorStatusTask    = errors.New("error Status Task")
	ErrorStatusAllTask = errors.New("error Status All Task")
	ErrorListUsers = errors.New("error List Users")

)
