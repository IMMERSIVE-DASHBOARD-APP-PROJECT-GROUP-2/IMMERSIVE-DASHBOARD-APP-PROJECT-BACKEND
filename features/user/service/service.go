package service

import (
	"github.com/DASHBOARDAPP/features/user"
	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

// Create implements user.UserServiceInterface.
func (*userService) Create(user *user.Core) error {
	panic("unimplemented")
}

// Login implements user.UserServiceInterface.
func (*userService) Login(email string, password string) (user.Core, string, error) {
	panic("unimplemented")
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}
