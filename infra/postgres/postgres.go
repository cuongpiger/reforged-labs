package postgres

import (
	lzap "go.uber.org/zap"
	lpostgres "gorm.io/driver/postgres"
	lgorm "gorm.io/gorm"
)

func InitPostgreSQL(puri string) (*lgorm.DB, error) {
	dbClient, err := lgorm.Open(lpostgres.Open(puri), &lgorm.Config{})

	if err != nil {
		lzap.L().Error("Failed to create PostgreSQL client", lzap.Error(err))
		return nil, err
	}

	return dbClient, nil
}
