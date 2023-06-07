package data

import (
	"errors"

	"github.com/DASHBOARDAPP/features/class"
	"gorm.io/gorm"
)

type classQuery struct {
	db *gorm.DB
}

// UpdateClassById implements class.ClassDataInterface.
func (repo *classQuery) UpdateClassById(id string, classInput class.Core) error {
	// Mencari pengguna berdasarkan ID
	var classData Class
	tx := repo.db.First(&classData, id)

	// Mengupdate data pengguna berdasarkan ID dari userInputGorm
	px := repo.db.Model(&classData).Updates(CoreToModel(classInput))
	if tx.Error != nil {
		return tx.Error
	} else if px.Error != nil {
		return px.Error
	}

	// Menyimpan perubahan data pengguna dari Input ke database
	tx = repo.db.Save(&classData)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("Updated Failed, row affected = 0")
	}
	return nil
}

// GetAllClass implements class.ClassDataInterface.
func (repo *classQuery) GetAllClass() ([]class.Core, error) {
	var classData []Class
	// Mencari data user di database
	tx := repo.db.Find(&classData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// mapping dari struct gorm model ke struct entities core
	var classCoreAll []class.Core
	for _, value := range classData {
		var classCore = class.Core{
			Id:        value.ID,
			Name:      value.Name,
			UserID:    value.UserID,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
			DeletedAt: value.DeletedAt.Time,
		}
		classCoreAll = append(classCoreAll, classCore)
	}

	return classCoreAll, nil
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
