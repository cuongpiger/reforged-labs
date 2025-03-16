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
	GetAdvertisementById(pctx lctx.Context, pid string) (*lsmdl.Advertisement, error)
	UpdateAdvertisement(pctx lctx.Context, pads *lsmdl.Advertisement) error
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

func (s *advertisementRepository) GetAdvertisementById(pctx lctx.Context, pid string) (*lsmdl.Advertisement, error) {
	var (
		log = lsutil.GetLogger(pctx).With(lzap.String("advertisementId", pid))
	)

	log.Info("Getting advertisement")
	adv := &lsmdl.Advertisement{}
	res := s.client.Where("id = ?", pid).First(adv)
	if res.Error != nil {
		log.Error("Failed to get advertisement", lzap.Error(res.Error))
		return nil, res.Error
	}

	log.Info("Successfully get advertisement")
	return adv, nil
}

func (s *advertisementRepository) UpdateAdvertisement(pctx lctx.Context, pads *lsmdl.Advertisement) error {
	var (
		log = lsutil.GetLogger(pctx).With(lzap.String("advertisementId", pads.Id))
	)

	log.Info("Updating advertisement")
	res := s.client.Save(pads)
	if res.Error != nil {
		log.Error("Failed to update advertisement", lzap.Error(res.Error))
		return res.Error
	}

	log.Info("Successfully updated advertisement")
	return nil
}
