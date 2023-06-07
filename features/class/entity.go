package class

import "time"

// struct class
type Core struct {
	Id        uint
	Name      string `json:"name" form:"name"`
	UserID    uint   `json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ClassDataInterface interface {
	CreateClass(classInput Core) error
	GetAllClass() ([]Core, error)
}

type ClassServiceInterface interface {
	CreateClass(classInput Core) error
	GetAllClass() ([]Core, error)
}
