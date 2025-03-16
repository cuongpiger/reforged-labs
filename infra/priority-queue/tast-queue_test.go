package priority_queue

import (
	lfmt "fmt"
	ltesting "testing"
	ltime "time"

	lstask "github.com/cuongpiger/reforged-labs/infra/task"
)

type Task struct {
	Id       string
	Priority int
	Content  string
	Index    int // Required for heap.Interface
}

func (t *Task) GetPriority() int {
	return t.Priority
}

func (t *Task) SetIndex(i int) {
	t.Index = i
}

func (t *Task) GetData() interface{} {
	return t.Content
}

func (t *Task) Clone() lstask.Task {
	return &Task{
		Id:       t.Id,
		Priority: t.Priority,
		Content:  t.Content,
		Index:    t.Index,
	}
}

func (t *Task) GetId() string {
	return t.Id
}

func (t *Task) Do() error {
	lfmt.Printf("Task ID: %s\n", t.Id)
	return nil
}

func TestNewTaskQueue(t *ltesting.T) {
	tq := NewTaskQueue()
	if tq == nil {
		t.Error("NewTaskQueue() should return a non-nil TaskQueue")
	}

	if tq.cond == nil {
		t.Error("TaskQueue.cond should not be nil")
	}

	if tq.pq == nil {
		t.Error("TaskQueue.pq should not be nil")
	}

	if len(tq.pq) != 0 {
		t.Error("TaskQueue.pq should be empty")
	}
}

func TestAddMultipleTasks(t *ltesting.T) {
	tq := NewTaskQueue()
	tasks := []lstask.Task{
		&Task{Priority: 3, Id: "task-1", Content: "Task 1"},
		&Task{Priority: 2, Id: "task-2", Content: "Task 2"},
		&Task{Priority: 1, Id: "task-3", Content: "Task 3"},
	}

	for _, task := range tasks {
		go tq.PushTask(task)
	}

	ltime.Sleep(3 * ltime.Second)

	if len(tq.pq) != len(tasks) {
		t.Errorf("TaskQueue.pq should have %d tasks, got %d", len(tasks), len(tq.pq))
	}

	for i := 0; i < 3; i++ {
		task := tq.PopTask()
		if task == nil {
			t.Error("TaskQueue.PopTask() should return a non-nil task")
		}

		lfmt.Printf("> Task ID: %s\n")
	}
}
