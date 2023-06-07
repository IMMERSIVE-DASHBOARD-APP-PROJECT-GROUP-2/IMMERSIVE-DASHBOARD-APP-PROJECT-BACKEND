package handler

type ClassResponse struct {
	Id     uint
	Name   string `json:"name" form:"name"`
	UserID uint   `json:"user_id"`
}
