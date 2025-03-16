package priority_queue

import (
	lheap "container/heap"
	lsync "sync"

	lstask "github.com/cuongpiger/reforged-labs/infra/task"
)

type TaskQueue struct {
	mu   lsync.Mutex
	pq   priorityQueue
	cond *lsync.Cond
}

func NewTaskQueue() *TaskQueue {
	pq := make(priorityQueue, 0)

	lheap.Init(&pq)
	tq := &TaskQueue{pq: pq}
	tq.cond = lsync.NewCond(&tq.mu) // Condition variable to signal workers
	return tq
}

func (s *TaskQueue) PushTask(task interface{}) {
	s.mu.Lock()
	lheap.Push(&s.pq, task)
	s.cond.Signal() // Wake up a waiting worker
	s.mu.Unlock()
}

func (s *TaskQueue) PopTask() lstask.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	for len(s.pq) == 0 {
		s.cond.Wait() // Wait until a task is available
	}

	task := lheap.Pop(&s.pq).(lstask.Task)
	return task
}
