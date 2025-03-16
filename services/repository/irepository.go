package repository

import (
	lgorm "gorm.io/gorm"
)

func NewRepository(db *lgorm.DB) IRepository {
	return &repository{
		db: db,
	}
}

type IRepository interface {
}

type repository struct {
	db *lgorm.DB
}
