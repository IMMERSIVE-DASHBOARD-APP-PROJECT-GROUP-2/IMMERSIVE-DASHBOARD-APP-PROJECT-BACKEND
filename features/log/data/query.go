package data

import (
	"github.com/DASHBOARDAPP/features/log"
	"gorm.io/gorm"
)

type logQuery struct {
	db *gorm.DB
}

// GetLogsByID implements log.LogDataInterface.
func (*logQuery) GetLogsByID(logID uint) ([]log.Core, error) {
	panic("unimplemented")
}

func (repo *logQuery) GetLogsByMenteeID(menteeID uint) ([]log.Core, error) {
	var logs []log.Core
	err := repo.db.Where("mentee_id = ?", menteeID).Find(&logs).Error
	if err != nil {
		// Handle error
		return nil, err
	}

	return logs, nil
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
