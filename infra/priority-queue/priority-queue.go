package priority_queue

import (
	ltask "github.com/cuongpiger/reforged-labs/infra/task"
)

type (
	priorityQueue []ltask.Task
)

func (s priorityQueue) Len() int {
	return len(s)
}

func (s priorityQueue) Less(i, j int) bool {
	return s[i].GetPriority() < s[j].GetPriority() // Lower value = higher priority
}

func (s priorityQueue) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	s[i].SetIndex(i)
	s[j].SetIndex(j)
}

func (s *priorityQueue) Push(ptask interface{}) {
	n := len(*s)
	item := ptask.(ltask.Task)
	item.SetIndex(n)
	*s = append(*s, item)
}

func (s *priorityQueue) Pop() interface{} {
	old := *s
	n := len(old)
	task := old[n-1]
	*s = old[0 : n-1]
	return task
}
