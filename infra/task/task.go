package task

type Task interface {
	GetPriority() int
	SetIndex(int)
	GetId() string
	GetData() interface{}
}
