package user

import "time"

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

type Core struct {
	Id        uint
	Name      string `json:"name" form:"name"`
	Phone     string `json:"phone" form:"phone"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Status    UserStatus
	Team      UserTeam
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDataInterface interface {
	Insert(user *Core) error
	Login(email, password string) (Core, string, error)
}

type UserServiceInterface interface {
	Create(user *Core) error
	Login(email, password string) (Core, string, error)
}
