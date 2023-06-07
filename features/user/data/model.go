package data

import (
	classGorm "github.com/DASHBOARDAPP/features/class/data"
	logGorm "github.com/DASHBOARDAPP/features/log/data"
	"github.com/DASHBOARDAPP/features/user"
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
	Name     string            `json:"name" form:"name"`
	Phone    string            `gorm:"unique" json:"phone" form:"phone"`
	Email    string            `gorm:"unique" json:"email" form:"email"`
	Password string            `json:"password" form:"password"`
	Status   UserStatus        `gorm:"type:ENUM('active', 'non_active')"`
	Team     UserTeam          `gorm:"type:ENUM('manager', 'mentor', 'team_placement', 'team_people_skill')"`
	Role     UserRole          `gorm:"type:ENUM('admin', 'karyawan')"`
	Logs     []logGorm.Log     // Relasi One-to-Many dengan model Log
	Classes  []classGorm.Class `gorm:"constraint:OnUpdate:CASCADE"` // Relasi One-to-Many dengan model Class
}

// mapping dari core ke gorm
func ModelToCore(dataCore *user.Core) *User {
	return &User{
		Name:     dataCore.Name,
		Phone:    dataCore.Phone,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Status:   UserStatus(dataCore.Status),
		Team:     UserTeam(dataCore.Team),
		Role:     UserRole(dataCore.Role),
	}
}

func CoreToModel(dataUser User) user.Core {
	return user.Core{
		Name:     dataUser.Name,
		Phone:    dataUser.Phone,
		Email:    dataUser.Email,
		Password: dataUser.Password,
		Status:   user.UserStatus(dataUser.Status),
		Team:     user.UserTeam(dataUser.Team),
		Role:     user.UserRole(dataUser.Role),
	}
}

// mapping dari core ke gorm
func CoreToUpdateModel(dataCore user.Core) User {
	return User{
		Name:     dataCore.Name,
		Phone:    dataCore.Phone,
		Email:    dataCore.Email,
		Password: dataCore.Password,
	}
}
