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

func (tq *TaskQueue) PopTask() lstask.Task {
	tq.mu.Lock()
	defer tq.mu.Unlock()

	for len(tq.pq) == 0 {
		tq.cond.Wait() // Wait until a task is available
	}

	task := lheap.Pop(&tq.pq).(lstask.Task)
	return task
}
