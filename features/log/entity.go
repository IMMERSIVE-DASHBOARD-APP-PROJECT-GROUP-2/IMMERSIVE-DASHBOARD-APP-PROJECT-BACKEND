package log

// struct class
type Core struct {
	ID          uint
	Description string `json:"description" validate:"required"`
	MenteeID    uint   `json:"mentee_id" validate:"required"`
	UserID      uint   `json:"user_id" validate:"required"`
}

type LogDataInterface interface {
	Create(logInput Core) error
}

type LogServiceInterface interface {
	Insert(logInput Core) error
}
