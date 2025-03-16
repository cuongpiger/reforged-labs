package postgres

import (
	lctx "context"
	lsutil "github.com/cuongpiger/reforged-labs/utils"
	lzap "go.uber.org/zap"

	lgorm "gorm.io/gorm"

	lsmdl "github.com/cuongpiger/reforged-labs/models"
)

func NewAdvertisementRepository(client *lgorm.DB) IAdvertisementRepository {
	return &advertisementRepository{
		client: client,
	}
}

type IAdvertisementRepository interface {
	CreateAdvertisement(pctx lctx.Context, pads *lsmdl.Advertisement) error
}

type advertisementRepository struct {
	client *lgorm.DB
}

func (s *advertisementRepository) CreateAdvertisement(pctx lctx.Context, pads *lsmdl.Advertisement) error {
	var (
		log = lsutil.GetLogger(pctx)
	)

	log.Info("Creating advertisement")
	res := s.client.Create(pads)
	if res.Error != nil {
		log.Error("Failed to create advertisement", lzap.Error(res.Error))
		return res.Error
	}

	log.Info("Successfully created advertisement")
	return nil
}
