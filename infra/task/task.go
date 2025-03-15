package task

type Task interface {
	GetPriority() int
	Do() error
	ID() string
	SetIndex(int)
	Clone() Task
}
