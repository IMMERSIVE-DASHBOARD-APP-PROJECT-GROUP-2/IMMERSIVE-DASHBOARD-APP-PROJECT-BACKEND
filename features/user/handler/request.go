package handler

type AuthRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
	Status   string `json:"status" form:"status"`
	Role     string `json:"role" form:"role"`
	Team     string `json:"team" form:"team"`
}
