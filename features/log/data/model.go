package data

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Description string
	MenteeID    uint //ID mentee yang memiliki log
	UserID      uint // ID user yang memberi Log
}
