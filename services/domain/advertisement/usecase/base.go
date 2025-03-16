package usecase

import (
	lctx "context"

	lsdto "github.com/cuongpiger/reforged-labs/dto"
	lsqueue "github.com/cuongpiger/reforged-labs/infra/priority-queue"
	lsmdl "github.com/cuongpiger/reforged-labs/models"
	lsrepo "github.com/cuongpiger/reforged-labs/services/repository"
)

type IAdvertisementUseCase interface {
	CreateAdvertisement(ctx lctx.Context, preq *lsdto.CreateAdvertisementRequestDTO, ptaskQueue *lsqueue.TaskQueue) (*lsmdl.Advertisement, error)
	GetAdvertisement(ctx lctx.Context, pid string) (*lsmdl.Advertisement, error)
}

type advertisementUseCase struct {
	repo lsrepo.IRepository
}

func NewAdvertisementUseCase(repo lsrepo.IRepository) IAdvertisementUseCase {
	return &advertisementUseCase{repo: repo}
}
