package mentee

import (
	"time"
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

type Core struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	// MainData
	Name        string       `json:"name" form:"name"`
	Address     string       `json:"address" form:"address"`
	HomeAddress string       `json:"current_address" form:"current_address"`
	Email       string       `json:"email" form:"email"`
	Gender      MenteeGender `json:"gender" form:"gender"`
	Telegram    string       `json:"telegram" form:"telegram"`
	Phone       string       `json:"phone" form:"phone"`
	Status      MenteeStatus `json:"status" form:"status"`
	// EmergencyData
	EmergencyName   string          `json:"emergency_name" form:"emergency_name"`
	EmergencyStatus EmergencyStatus `json:"emergency_status" form:"emergency_status"`
	EmergencyPhone  string          `json:"emergency_phone" form:"emergency_phone"`
	// EducationData
	Category  MenteeCategory `json:"education_type" form:"education_type"`
	Major     string         `json:"major" form:"major" validate:"required"`
	Graduated string         `json:"graduated" form:"graduated"`
	ClassID   uint           `json:"class_id" form:"class_id"`
}

type MenteeDataInterface interface {
	CreateMentee(menteeInput Core) error
	GetAllMentee(keyword string) ([]Core, error)
}

type MenteeServiceInterface interface {
	CreateMentee(menteeInput Core) error
	GetAllMentee(keyword string) ([]Core, error)
}
