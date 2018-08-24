package enum

// Type represents the type of a specific task
type Type string

func (t Type) String() string {
	return string(t)
}

const (
	TaskCreate = Type("TaskCreate")
)

type State string

func (s State) String() string {
	return string(s)
}

const (
	TaskPending = State("task-pending")
	TaskDoing   = State("task-doing")
	TaskDone    = State("task-done")
)
