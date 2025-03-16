package usecase

import (
	lctx "context"
	lfmt "fmt"
	lsmdl "github.com/cuongpiger/reforged-labs/models"
	lrand "math/rand"
	ltime "time"

	lzap "go.uber.org/zap"

	lstask "github.com/cuongpiger/reforged-labs/infra/task"
	lsrepo "github.com/cuongpiger/reforged-labs/services/repository"
)

type TaskChain interface {
	Next(pctx lctx.Context, ptask lstask.Task) error
}

type InQueueTaskChain struct {
	NextChain  TaskChain
	Repository lsrepo.IRepository
}

type ProcessingTaskChain struct {
	NextChain  TaskChain
	Repository lsrepo.IRepository
}

type CompletedTaskChain struct {
	Repository lsrepo.IRepository
}

func (s *InQueueTaskChain) Next(pctx lctx.Context, ptask lstask.Task) error {
	lzap.L().Info("Get the advertisement based on its ID", lzap.String("advertisementId", ptask.GetId()))

	advertisement, ok := ptask.GetData().(*lsmdl.Advertisement)
	if !ok {
		lzap.L().Error("Failed to cast task to advertisement task")
		return lfmt.Errorf(lfmt.Sprintf("Failed to cast task to advertisement task with ID %s", ptask.GetId()))
	}

	advertisement.Status = "queued"
	err := s.Repository.NewAdvertisementRepo().UpdateAdvertisement(pctx, advertisement)
	if err != nil {
		lzap.L().Error("Failed to update advertisement", lzap.Error(err))
		return err
	}

	if s.NextChain != nil {
		return s.NextChain.Next(pctx, NewAdvertisementTask(advertisement))
	}

	return nil
}

func (s *ProcessingTaskChain) Next(pctx lctx.Context, ptask lstask.Task) error {
	lzap.L().Info("Get the advertisement based on its ID", lzap.String("advertisementId", ptask.GetId()))

	advertisement, ok := ptask.GetData().(*lsmdl.Advertisement)
	if !ok {
		lzap.L().Error("Failed to cast task to advertisement task")
		return lfmt.Errorf(lfmt.Sprintf("Failed to cast task to advertisement task with ID %s", ptask.GetId()))
	}

	// Process the logic here
	advertisement.Status = "processing"
	err := s.Repository.NewAdvertisementRepo().UpdateAdvertisement(pctx, advertisement)
	if err != nil {
		lzap.L().Error("Failed to update advertisement", lzap.Error(err))
		return err
	}

	ltime.Sleep(ltime.Duration(lrand.Intn(10-3+1)+3) * ltime.Second)

	if s.NextChain != nil {
		return s.NextChain.Next(pctx, NewAdvertisementTask(advertisement))
	}

	return nil
}

func (s *CompletedTaskChain) Next(pctx lctx.Context, ptask lstask.Task) error {
	lzap.L().Info("Get the advertisement based on its ID", lzap.String("advertisementId", ptask.GetId()))

	advertisement, ok := ptask.GetData().(*lsmdl.Advertisement)
	if !ok {
		lzap.L().Error("Failed to cast task to advertisement task")
		return lfmt.Errorf(lfmt.Sprintf("Failed to cast task to advertisement task with ID %s", ptask.GetId()))
	}

	now := ltime.Now()
	advertisement.CompleteAt = &now
	advertisement.Status = "completed"
	err := s.Repository.NewAdvertisementRepo().UpdateAdvertisement(pctx, advertisement)
	if err != nil {
		lzap.L().Error("Failed to update advertisement", lzap.Error(err))
		return err
	}

	return nil
}
