package data

import (
	"github.com/DASHBOARDAPP/features/class"
	menteeGorm "github.com/DASHBOARDAPP/features/mentee/data"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name    string
	UserID  uint                // ID pengguna yang memiliki kelas ini
	Mentees []menteeGorm.Mentee // Relasi One-to-Many dengan model Mentee
}

// mapping dari core ke gorm
func CoreToModel(dataCore class.Core) Class {
	return Class{
		Name:   dataCore.Name,
		UserID: dataCore.UserID,
	}
}
