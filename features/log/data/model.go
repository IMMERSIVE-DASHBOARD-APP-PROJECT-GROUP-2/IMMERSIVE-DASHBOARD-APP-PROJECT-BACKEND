package data

import (
	"github.com/DASHBOARDAPP/features/log"
	"github.com/DASHBOARDAPP/features/mentee"
	"github.com/DASHBOARDAPP/features/user"
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Description string
	MenteeID    uint         //ID mentee yang memiliki log
	UserID      uint         // ID user yang memberi Log
	User        *user.Core   // Pengguna yang memberi Log
	Mentee      *mentee.Core // Mentee yang memiliki log

}

// mapping dari core ke gorm
func CoreToModel(dataCore log.Core) Log {
	return Log{
		Description: dataCore.Description,
		MenteeID:    dataCore.MenteeID,
		UserID:      dataCore.UserID,
	}
}
