package handler

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

type MenteeRequest struct {
	// ClassId
	ClassID uint `json:"class_id" form:"class_id"`
	// MainData
	Name        string       `json:"name" form:"name"`
	Phone       string       `json:"phone" form:"phone"`
	Email       string       `json:"email" form:"email"`
	Address     string       `json:"current_address" form:"current_address"`
	HomeAddress string       `json:"home_address" form:"home_address"`
	Telegram    string       `json:"telegram" form:"telegram"`
	Gender      MenteeGender `json:"gender" form:"gender"`
	// EducationData
	Category  MenteeCategory `json:"education_type" form:"education_type"`
	Major     string         `json:"major" form:"major"`
	Graduated string         `json:"graduated" form:"graduated"`
	// EmergencyData
	EmergencyName   string          `json:"emergency_name" form:"emergency_name"`
	EmergencyStatus EmergencyStatus `json:"emergency_status" form:"emergency_status"`
	EmergencyPhone  string          `json:"emergency_phone" form:"emergency_phone"`
	Status          MenteeStatus    `json:"status" form:"status"`
}
