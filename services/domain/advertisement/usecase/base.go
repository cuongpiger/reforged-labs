package usecase

import lsrepo "github.com/cuongpiger/reforged-labs/services/repository"

type IAdvertisementUseCase interface {
}

type advertisementUseCase struct {
	repo lsrepo.IRepository
}

func NewAdvertisementUseCase(repo lsrepo.IRepository) IAdvertisementUseCase {
	return &advertisementUseCase{repo: repo}
}
