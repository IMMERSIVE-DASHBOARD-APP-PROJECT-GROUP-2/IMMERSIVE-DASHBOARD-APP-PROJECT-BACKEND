package data

import (
	logGorm "github.com/DASHBOARDAPP/features/log/data"
	"gorm.io/gorm"
)

type MenteeCategory string
type MenteeStatus string

const (
	IT    MenteeCategory = "it"
	NonIT MenteeCategory = "non_it"
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

type Mentee struct {
	gorm.Model
	Name     string
	Address  string
	Phone    string
	Category MenteeCategory `gorm:"type:ENUM('it','non_it')"`
	Status   MenteeStatus   `gorm:"type:ENUM('interview', 'join_class', 'unit1', 'unit2', 'unit3', 'repeat_unit1', 'repeat_unit2', 'repeat_unit3', 'placement', 'eliminated', 'graduate')"`
	ClassID  uint
	Logs     []logGorm.Log // Relasi One-to-Many dengan model Class
}
