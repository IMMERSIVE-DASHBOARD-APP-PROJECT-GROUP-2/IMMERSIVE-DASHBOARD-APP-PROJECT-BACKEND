package data

import (
	"errors"
	"fmt"

	"github.com/DASHBOARDAPP/features/mentee"
	"gorm.io/gorm"
)

type menteeQuery struct {
	db *gorm.DB
}

// GetMenteeByID implements mentee.MenteeDataInterface.
func (repo *menteeQuery) GetMenteeByID(menteeID uint) (*mentee.Core, error) {
	menteeData := &mentee.Core{}
	err := repo.db.Preload("Logs").First(menteeData, menteeID).Error
	if err != nil {
		return nil, err
	}
	return menteeData, nil
}

func (query *menteeQuery) DeleteMentee(menteeID uint) error {
	var mentee Mentee
	result := query.db.Where("id = ?", menteeID).First(&mentee)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("record not found")
		}
		return fmt.Errorf("failed to find mentee: %v", result.Error)
	}

	// Soft delete the mentee by setting the "deleted_at" field
	result = query.db.Delete(&mentee)
	if result.Error != nil {
		return fmt.Errorf("failed to delete mentee: %v", result.Error)
	}

	return nil
}

// UpdateMentee implements mentee.MenteeDataInterface.
func (repo *menteeQuery) UpdateMentee(menteeInput mentee.Core) error {
	menteeToUpdate := Mentee{}
	// Cari entri mentee berdasarkan ID
	if err := repo.db.First(&menteeToUpdate, menteeInput.Id).Error; err != nil {
		return err
	}

	// Mapping data dari menteeInput ke menteeToUpdate menggunakan fungsi CoreToModel
	updatedMentee := CoreToModel(&menteeInput)

	// Perbarui data menteeToUpdate dengan data baru dari updatedMentee
	menteeToUpdate.Name = updatedMentee.Name
	menteeToUpdate.Address = updatedMentee.Address
	menteeToUpdate.HomeAddress = updatedMentee.HomeAddress
	menteeToUpdate.Email = updatedMentee.Email
	menteeToUpdate.Gender = updatedMentee.Gender
	menteeToUpdate.Telegram = updatedMentee.Telegram
	menteeToUpdate.Phone = updatedMentee.Phone
	menteeToUpdate.Status = updatedMentee.Status
	menteeToUpdate.EmergencyName = updatedMentee.EmergencyName
	menteeToUpdate.EmergencyStatus = updatedMentee.EmergencyStatus
	menteeToUpdate.EmergencyPhone = updatedMentee.EmergencyPhone
	menteeToUpdate.Category = updatedMentee.Category
	menteeToUpdate.Major = updatedMentee.Major
	menteeToUpdate.Graduated = updatedMentee.Graduated
	menteeToUpdate.ClassID = updatedMentee.ClassID

	// Simpan perubahan ke database
	if err := repo.db.Save(&menteeToUpdate).Error; err != nil {
		return err
	}

	return nil
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
func (repo *menteeQuery) GetAllMentee(keyword string) ([]mentee.Core, error) {
	var menteeData []Mentee
	// Mencari data mentee di database
	tx := repo.db
	if keyword != "" {
		tx = tx.Where("id LIKE ?", "%"+keyword+"%").
			Or("name LIKE ?", "%"+keyword+"%").
			Or("email LIKE ?", "%"+keyword+"%").
			Or("class_id LIKE ?", "%"+keyword+"%").
			Or("status LIKE ?", "%"+keyword+"%").
			Or("category LIKE ?", "%"+keyword+"%").
			Or("gender LIKE ?", "%"+keyword+"%")
	}
	tx = tx.Find(&menteeData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// mapping dari struct gorm model ke struct entities core
	var menteeCoreAll []mentee.Core
	for _, value := range menteeData {
		var menteeCore = mentee.Core{
			Id:              value.ID,
			Name:            value.Name,
			Address:         value.Address,
			HomeAddress:     value.HomeAddress,
			Email:           value.Email,
			Gender:          mentee.MenteeGender(value.Gender),
			Telegram:        value.Telegram,
			Phone:           value.Phone,
			Status:          mentee.MenteeStatus(value.Status),
			EmergencyName:   value.EmergencyName,
			EmergencyStatus: mentee.EmergencyStatus(value.EmergencyStatus),
			EmergencyPhone:  value.EmergencyPhone,
			Category:        mentee.MenteeCategory(value.Category),
			Major:           value.Major,
			Graduated:       value.Graduated,
			ClassID:         value.ClassID,
		}
		menteeCoreAll = append(menteeCoreAll, menteeCore)
	}

	return menteeCoreAll, nil
}

func New(db *gorm.DB) mentee.MenteeDataInterface {
	return &menteeQuery{
		db: db,
	}
}
