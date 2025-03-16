package worker_pool

import (
	ltime "time"

	lzap "go.uber.org/zap"
)

type (
	Dispatcher interface {
		LaunchWorker(w WorkerLauncher)
		MakeRequest(r Request)
		Stop()
	}

	dispatcher struct {
		inChan chan Request
	}
)

func NewDispatcher(psize int) Dispatcher {
	return &dispatcher{
		inChan: make(chan Request, psize),
	}
}

func (s *dispatcher) LaunchWorker(pworker WorkerLauncher) {
	pworker.LaunchWorker(s.inChan)
}

func (s *dispatcher) MakeRequest(preq Request) {
	select {
	case s.inChan <- preq:
		lzap.L().Info("Request accepted")
	case <-ltime.After(ltime.Second * 5):
		lzap.L().Info("Request rejected")
		return
	}
}

func (s *dispatcher) Stop() {
	close(s.inChan)
}
