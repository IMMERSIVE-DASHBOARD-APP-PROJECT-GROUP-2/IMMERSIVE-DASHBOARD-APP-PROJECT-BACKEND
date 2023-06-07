package data

import (
	logGorm "github.com/DASHBOARDAPP/features/log/data"
	"github.com/DASHBOARDAPP/features/mentee"
	"gorm.io/gorm"
)

type MenteeCategory string
type MenteeStatus string
type MenteeGender string
type EmergencyStatus string

const (
	IT    MenteeCategory = "IT"
	NonIT MenteeCategory = "Non_IT"
)
const (
	Interview   MenteeStatus = "interview"
	JoinClass   MenteeStatus = "join_class"
	Unit1       MenteeStatus = "unit1"
	Unit2       MenteeStatus = "unit2"
	Unit3       MenteeStatus = "unit3"
	RepeatUnit1 MenteeStatus = "repeat_unit1"
	RepeatUnit2 MenteeStatus = "repeat_unit2"
	RepeatUnit3 MenteeStatus = "repeat_unit3"
	Placement   MenteeStatus = "placement"
	Eliminated  MenteeStatus = "eliminated"
	Graduate    MenteeStatus = "graduate"
)
const (
	Male   MenteeGender = "male"
	Female MenteeGender = "female"
)
const (
	Parent  EmergencyStatus = "orang_tua"
	Sibling EmergencyStatus = "saudara_kandung"
	Grandpa EmergencyStatus = "kakek"
	Grandma EmergencyStatus = "nenek"
)

type Mentee struct {
	gorm.Model
	// MainData
	Name        string       `json:"name" form:"name"`
	Address     string       `json:"address" form:"address"`
	HomeAddress string       `json:"current_address" form:"current_address"`
	Email       string       `gorm:"unique" json:"email" form:"email"`
	Gender      MenteeGender `gorm:"type:ENUM('male','female')"`
	Telegram    string       `gorm:"unique" json:"telegram" form:"telegram"`
	Phone       string       `gorm:"unique" json:"phone" form:"phone"`
	Status      MenteeStatus `gorm:"type:ENUM('interview', 'join_class', 'unit1', 'unit2', 'unit3', 'repeat_unit1', 'repeat_unit2', 'repeat_unit3', 'placement', 'eliminated', 'graduate')"`
	// EmergencyData
	EmergencyName   string          `json:"emergency_name" form:"emergency_name"`
	EmergencyStatus EmergencyStatus `gorm:"type:ENUM('orang_tua', 'saudara_kandung', 'kakek', 'nenek')"`
	EmergencyPhone  string          `json:"emergency_phone" form:"emergency_phone"`
	// EducationData
	Category  MenteeCategory `gorm:"type:ENUM('IT','Non_IT')"`
	Major     string         `json:"major" form:"major"`
	Graduated string         `json:"graduated" form:"graduated"`
	ClassID   uint           `json:"class_id" form:"class_id"`
	Logs      []logGorm.Log  // Relasi One-to-Many dengan model Class
}

// mapping dari core ke gorm
func CoreToModel(dataMentee *mentee.Core) *Mentee {
	return &Mentee{
		Name:            dataMentee.Name,
		Address:         dataMentee.Address,
		HomeAddress:     dataMentee.HomeAddress,
		Email:           dataMentee.Email,
		Gender:          MenteeGender(dataMentee.Gender),
		Telegram:        dataMentee.Telegram,
		Phone:           dataMentee.Phone,
		Status:          MenteeStatus(dataMentee.Status),
		EmergencyName:   dataMentee.EmergencyName,
		EmergencyStatus: EmergencyStatus(dataMentee.EmergencyStatus),
		EmergencyPhone:  dataMentee.EmergencyPhone,
		Category:        MenteeCategory(dataMentee.Category),
		Major:           dataMentee.Major,
		Graduated:       dataMentee.Graduated,
		ClassID:         dataMentee.ClassID,
	}
}
