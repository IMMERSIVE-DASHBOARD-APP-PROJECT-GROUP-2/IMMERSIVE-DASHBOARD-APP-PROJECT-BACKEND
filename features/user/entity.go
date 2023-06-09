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
	Name      string `json:"name" form:"name" validate:"required"`
	Phone     string `json:"phone" form:"phone" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required,email"`
	Password  string `json:"password" form:"password" validate:"required"`
	Status    UserStatus
	Team      UserTeam
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type LoginInput struct {
	Email    string `json:"name" form:"name" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UpdatedInput struct {
	Name     string `json:"name" form:"name"`
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserDataInterface interface {
	Insert(user Core) error
	Login(email, password string) (Core, string, error)
	GetAllUser(keyword string) ([]Core, error)
	GetUserByID(userID uint) (*Core, error)
	GetRoleByID(userID int) (UserRole, error)
	Update(userID int, updatedUser Core) error
	Delete(userID int) error
	UpdateUserById(id string, userInput Core) error
}

type UserServiceInterface interface {
	GetRoleByID(userID int) (UserRole, error)
	Create(user Core, loggedInUserID int) error
	Login(email, password string) (Core, string, error)
	GetAllUser(keyword string) ([]Core, error)
	Update(userID int, updatedUser Core, loggedInUserID int) error
	Delete(userID int, loggedInUserID int) error
	UpdateUserById(id string, userInput Core) error
}
