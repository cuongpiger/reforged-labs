package worker_pool

import lzap "go.uber.org/zap"

type (
	WorkerLauncher interface {
		LaunchWorker(in chan Request)
	}
)

type AdvertisementWorker struct {
	workerId int
}

func NewAdvertisementWorker(workerId int) WorkerLauncher {
	return &AdvertisementWorker{workerId: workerId}
}

func (s *AdvertisementWorker) LaunchWorker(pin chan Request) {
	go func() {
		for msg := range pin {
			err := msg.Handler()
			if err != nil {
				lzap.L().Error("Failed to process the task", lzap.Error(err), lzap.String("taskID", msg.Task.GetId()))
			}
		}
	}()
}
