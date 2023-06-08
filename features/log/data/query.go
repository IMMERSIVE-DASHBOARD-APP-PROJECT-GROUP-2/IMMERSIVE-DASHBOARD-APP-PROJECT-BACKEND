package data

import (
	"github.com/DASHBOARDAPP/features/log"
	"gorm.io/gorm"
)

type logQuery struct {
	db *gorm.DB
}

// Createlog implements log.LogDataInterface.
func (repo *logQuery) Create(logInput log.Core) error {
	log := CoreToModel(logInput)

	err := repo.db.Create(&log).Error
	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func New(db *gorm.DB) log.LogDataInterface {
	return &logQuery{
		db: db,
	}
}
