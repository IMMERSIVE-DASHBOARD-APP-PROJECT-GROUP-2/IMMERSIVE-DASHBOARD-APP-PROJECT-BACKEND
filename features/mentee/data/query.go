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
	var menteeData []Mentee
	// Mencari data user di database
	tx := repo.db.Preload("Class").Find(&menteeData)
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
