package handler

type UserRole string
type UserTeam string
type UserStatus string

const (
	Admin    UserRole = "admin"
	Karyawan UserRole = "karyawan"
)

const (
	Active    UserStatus = "active"
	NonActive UserStatus = "non_active"
)

const (
	Manager         UserTeam = "manager"
	Mentor          UserTeam = "mentor"
	TeamPlacement   UserTeam = "team_placement"
	TeamPeopleSkill UserTeam = "team_people_skill"
)

type UserResponse struct {
	Id     uint
	Name   string `json:"name" form:"name"`
	Email  string `json:"email" form:"email"`
	Team   UserTeam
	Role   UserRole
	Status UserStatus
}
