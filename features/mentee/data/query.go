package data

import (
	"errors"

	"github.com/DASHBOARDAPP/features/mentee"
	"gorm.io/gorm"
)

type menteeQuery struct {
	db *gorm.DB
}

// CreateMentee implements mentee.MenteeDataInterface.
func (repo *menteeQuery) CreateMentee(menteeInput mentee.Core) error {
	// mapping dari struct entities core ke gorm model
	menteeInputGorm := CoreToModel(&menteeInput)
	tx := repo.db.Create(&menteeInputGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("Insert Failed, row affected = 0")
	}
	return nil
}

// GetAllMentee implements mentee.MenteeDataInterface.
func (repo *menteeQuery) GetAllMentee() ([]mentee.Core, error) {
	panic("unimplemented")
}

func New(db *gorm.DB) mentee.MenteeDataInterface {
	return &menteeQuery{
		db: db,
	}
}
