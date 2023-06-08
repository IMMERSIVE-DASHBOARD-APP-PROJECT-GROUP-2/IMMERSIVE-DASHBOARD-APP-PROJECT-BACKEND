package handler

type LogRequest struct {
	Description string `json:"description" validate:"required"`
	MenteeID    uint   `json:"mentee_id" validate:"required"`
	UserID      uint   `json:"user_id" validate:"required"`
}
