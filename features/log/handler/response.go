package handler

type logResponse struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	MenteeID    uint   `json:"mentee_id"`
	UserID      uint   `json:"user_id"`
	MenteeName  string `json:"mentee_name"`
	UserName    string `json:"user_name"`
}
