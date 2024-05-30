package enum

type TaskStatus int

const (
	_ TaskStatus = iota
	TaskInit
	TaskRun
	TaskFinish
	TaskStop
	TaskPause
)
