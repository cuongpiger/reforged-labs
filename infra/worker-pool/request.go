package worker_pool

import (
	lstask "github.com/cuongpiger/reforged-labs/infra/task"
)

type (
	Request struct {
		Task    lstask.Task
		Handler RequestHandler
	}

	RequestHandler func() error
)
