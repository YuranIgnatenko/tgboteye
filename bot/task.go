package bot

import (
	"fmt"
)

type Task struct {
	TargetUser   string
	Status       bool
	abortChannel chan struct{}
}

func NewTask(targeUser string) *Task {
	abort_ch := make(chan struct{})
	return &Task{
		TargetUser:   targeUser,
		Status:       false,
		abortChannel: abort_ch,
	}
}

func (t *Task) Start() {
	t.Status = true
	for {
		select {
		case <-t.abortChannel:
			t.Status = false
			return
		default:
			LineGet(fmt.Sprintf("python3 cmd/check_inline.py %s", t.TargetUser))
		}
	}

}

func (t *Task) Abort() {
	t.abortChannel <- struct{}{}
}
