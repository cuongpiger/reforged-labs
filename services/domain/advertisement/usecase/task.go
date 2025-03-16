package usecase

import (
	lstask "github.com/cuongpiger/reforged-labs/infra/task"
	lsmdl "github.com/cuongpiger/reforged-labs/models"
)

type AdvertisementTask struct {
	Advertisement *lsmdl.Advertisement
	Index         int // Required for heap.Interface
}

func (s *AdvertisementTask) GetPriority() int {
	return s.Advertisement.Priority
}

func (s *AdvertisementTask) SetIndex(i int) {
	s.Index = i
}

func (s *AdvertisementTask) GetId() string {
	return s.Advertisement.Id
}

func (s *AdvertisementTask) GetData() interface{} {
	return s.Advertisement
}

func NewAdvertisementTask(advertisement *lsmdl.Advertisement) lstask.Task {
	return &AdvertisementTask{
		Advertisement: advertisement,
	}
}
