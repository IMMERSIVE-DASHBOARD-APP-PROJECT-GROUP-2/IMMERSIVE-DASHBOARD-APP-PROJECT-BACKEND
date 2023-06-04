package data

import (
	"github.com/DASHBOARDAPP/features/user"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// Insert implements user.UserDataInterface.
func (*userQuery) Insert(user *user.Core) error {
	panic("unimplemented")
}

// Login implements user.UserDataInterface.
func (*userQuery) Login(email string, password string) (user.Core, string, error) {
	panic("unimplemented")
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}
