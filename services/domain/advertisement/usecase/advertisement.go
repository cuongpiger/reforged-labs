package usecase

import (
	lctx "context"
	lzap "go.uber.org/zap"
	ltime "time"

	luuid "github.com/google/uuid"

	lsdto "github.com/cuongpiger/reforged-labs/dto"
	lsmdl "github.com/cuongpiger/reforged-labs/models"
	lsutil "github.com/cuongpiger/reforged-labs/utils"
)

func (s *advertisementUseCase) CreateAdvertisement(ctx lctx.Context, preq *lsdto.CreateAdvertisementRequestDTO) (*lsmdl.Advertisement, error) {
	var (
		log = lsutil.GetLogger(ctx)
	)

	log.Info("Create advertisement")
	uuid := "ads-" + luuid.New().String()
	adv := &lsmdl.Advertisement{
		Id:       uuid,
		Status:   "submitted",
		Priority: preq.Priority,
		CreateAt: ltime.Now(),
	}

	if err := s.repo.NewAdvertisementRepo().CreateAdvertisement(ctx, adv); err != nil {
		log.Error("Create advertisement failed", lzap.Error(err))
		return nil, err
	}

	log.Info("Create advertisement success")
	return adv, nil
}

func (s *advertisementUseCase) GetAdvertisement(ctx lctx.Context, pid string) (*lsmdl.Advertisement, error) {
	var (
		log = lsutil.GetLogger(ctx)
	)

	log.Info("Get advertisement")
	adv, err := s.repo.NewAdvertisementRepo().GetAdvertisementById(ctx, pid)
	if err != nil {
		log.Error("Get advertisement failed", lzap.Error(err))
		return nil, err
	}

	log.Info("Get advertisement success")
	return adv, nil
}
