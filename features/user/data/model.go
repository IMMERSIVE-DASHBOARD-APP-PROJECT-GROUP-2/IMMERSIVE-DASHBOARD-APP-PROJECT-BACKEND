package data

import (
	classGorm "github.com/DASHBOARDAPP/features/class/data"
	logGorm "github.com/DASHBOARDAPP/features/log/data"
	"gorm.io/gorm"
)

type UserRole string
type UserTeam string
type UserStatus string

const (
	Admin    UserRole = "admin"
	Karyawan UserRole = "karyawan"
)
const (
	Active    UserStatus = "active"
	NonActive UserStatus = "non_active"
)
const (
	Manager         UserTeam = "manager"
	Mentor          UserTeam = "mentor"
	TeamPlacement   UserTeam = "team_placement"
	TeamPeopleSkill UserTeam = "team_people_skill"
)

type User struct {
	gorm.Model
	Name     string
	Phone    string
	Email    string `gorm:"unique" `
	Password string
	Status   UserStatus        `gorm:"type:ENUM('active', 'non_active')"`
	Team     UserTeam          `gorm:"type:ENUM('manager', 'mentor', 'team_placement', 'team_people_skill')"`
	Role     UserRole          `gorm:"type:ENUM('admin', 'karyawan')"`
	Logs     []logGorm.Log     // Relasi One-to-Many dengan model Log
	Classes  []classGorm.Class // Relasi One-to-Many dengan model Class
}
