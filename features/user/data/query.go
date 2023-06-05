package data

import (
	"errors"

	"github.com/DASHBOARDAPP/app/middlewares"
	"github.com/DASHBOARDAPP/features/user"
	"github.com/DASHBOARDAPP/helper"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// GetAllUser implements user.UserDataInterface.
func (repo *userQuery) GetAllUser() ([]user.Core, error) {
	var userData []User
	tx := repo.db.Find(&userData) // select * from users
	if tx.Error != nil {
		return nil, tx.Error
	}

	// mapping dari struct gorm model ke struct entities core
	var usersCoreAll []user.Core
	for _, value := range userData {
		var userCore = user.Core{
			Id:        value.ID,
			Name:      value.Name,
			Phone:     value.Phone,
			Email:     value.Email,
			Password:  value.Password,
			Status:    user.UserStatus(value.Status),
			Team:      user.UserTeam(value.Team),
			Role:      user.UserRole(value.Role),
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
			DeletedAt: value.DeletedAt.Time,
		}
		usersCoreAll = append(usersCoreAll, userCore)
	}

	return usersCoreAll, nil
}

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(email string, password string) (user.Core, string, error) {
	var userData User
	tx := repo.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return user.Core{}, "", tx.Error
	}
	if tx.RowsAffected == 0 {
		return user.Core{}, "", errors.New("login failed, email salah")
	}
	checkPassword := helper.CheckPasswordHash(password, userData.Password)
	if !checkPassword {
		return user.Core{}, "", errors.New("login failed, password salah")
	}
	if userData.Status == NonActive {
		return user.Core{}, "", errors.New("hanya user dengan status aktif yang dapat melakukan login")
	}

	token, errToken := middlewares.CreateToken(int(userData.ID))
	if errToken != nil {
		return user.Core{}, "", errToken
	}

	dataCore := user.Core{
		Id:        userData.ID,
		Name:      userData.Name,
		Phone:     userData.Phone,
		Email:     userData.Email,
		Password:  userData.Password,
		Status:    user.UserStatus(userData.Status),
		Team:      user.UserTeam(userData.Team),
		Role:      user.UserRole(userData.Role),
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}

	return dataCore, token, nil
}

// Insert implements user.UserDataInterface.
// func (*userQuery) Insert(user *user.Core) error {
// 	panic("unimplemented")
// }

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}
