package data

import (
	"errors"

	"github.com/DASHBOARDAPP/features/class"
	"gorm.io/gorm"
)

type classQuery struct {
	db *gorm.DB
}

// CreateClass implements class.ClassDataInterface.
func (repo *classQuery) CreateClass(classInput class.Core) error {
	// mapping dari struct entities core ke gorm model
	classInputGorm := CoreToModel(classInput)
	tx := repo.db.Create(&classInputGorm) //Ini query insert ke database
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("Insert Failed, row affected = 0")
	}
	return nil
}

func New(db *gorm.DB) class.ClassDataInterface {
	return &classQuery{
		db: db,
	}
}
