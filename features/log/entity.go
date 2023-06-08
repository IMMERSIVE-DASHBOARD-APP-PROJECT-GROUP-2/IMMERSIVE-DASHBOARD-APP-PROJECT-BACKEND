package log

import (
	"github.com/DASHBOARDAPP/features/mentee"
	"github.com/DASHBOARDAPP/features/user"
)

// struct class
type Core struct {
	ID          uint
	Description string       `json:"description" validate:"required"`
	MenteeID    uint         `json:"mentee_id" validate:"required"`
	UserID      uint         `json:"user_id" validate:"required"`
	User        *user.Core   // Pengguna yang memberi Log
	Mentee      *mentee.Core // Mentee yang memiliki log
	MenteeName  string       // Mentee's name
	UserName    string
}

type LogDataInterface interface {
	Create(logInput Core) error
	GetLogsByMenteeID(menteeID uint) ([]Core, error)
	GetLogsByID(logID uint) ([]Core, error)
}

type LogServiceInterface interface {
	Insert(logInput Core) error
	GetLogsByMenteeID(menteeID uint) ([]Core, error)
	GetLogsByID(logID uint) ([]Core, error)
}
