package data

import (
	"github.com/DASHBOARDAPP/features/log"
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Description string
	MenteeID    uint //ID mentee yang memiliki log
	UserID      uint // ID user yang memberi Log
}

// mapping dari core ke gorm
func CoreToModel(dataCore log.Core) Log {
	return Log{
		Description: dataCore.Description,
		MenteeID:    dataCore.MenteeID,
		UserID:      dataCore.UserID,
	}
}
