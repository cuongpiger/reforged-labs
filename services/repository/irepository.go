package repository

import (
	lsync "sync"

	lgorm "gorm.io/gorm"

	lsadsrepo "github.com/cuongpiger/reforged-labs/services/repository/postgres"
)

func NewRepository(db *lgorm.DB) IRepository {
	return &repository{
		db: db,
	}
}

var (
	advertisementsRepo     lsadsrepo.IAdvertisementRepository
	advertisementsRepoOnce lsync.Once
)

type IRepository interface {
	NewAdvertisementRepo() lsadsrepo.IAdvertisementRepository
}

type repository struct {
	db *lgorm.DB
}

func (s *repository) NewAdvertisementRepo() lsadsrepo.IAdvertisementRepository {
	advertisementsRepoOnce.Do(func() {
		advertisementsRepo = lsadsrepo.NewAdvertisementRepository(s.db)
	})

	return advertisementsRepo
}
