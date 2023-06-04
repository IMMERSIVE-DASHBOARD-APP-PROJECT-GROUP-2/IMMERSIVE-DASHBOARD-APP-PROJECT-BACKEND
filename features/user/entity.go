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
	Name      string `validate:"required"`
	Phone     string
	Email     string `validate:"required,email"`
	Password  string
	Status    UserStatus
	Team      UserTeam
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDataInterface interface {
	Login(email, password string) (Core, string, error)
	Insert(user *Core) error
}

type UserServiceInterface interface {
	Create(user *Core) error
	Login(email, password string) (Core, string, error)
}
