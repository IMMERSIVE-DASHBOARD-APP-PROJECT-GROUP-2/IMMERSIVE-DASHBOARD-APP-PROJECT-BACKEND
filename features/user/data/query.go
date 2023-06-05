package data

import (
	"errors"
	"fmt"

	"github.com/DASHBOARDAPP/app/middlewares"
	"github.com/DASHBOARDAPP/features/user"
	"github.com/DASHBOARDAPP/helper"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func (repo *userQuery) Insert(user *user.Core) error {
	if user.Role != helper.NewUserRole("admin") && user.Team != helper.NewUserTeam("manager") {
		return fmt.Errorf("only admin and manager can add users")
	}

	// Create a new database model from the user core data
	userData := ModelToCore(user)
	fmt.Println(user.Password)
	// Hash password sebelum disimpan
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}

	// Mengganti password dengan hashed password
	userData.Password = hashedPassword

	// Insert the user data into the database
	if err := repo.db.Create(&userData).Error; err != nil {
		return err
	}

	return nil
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
	fmt.Println(email)
	fmt.Println(password)
	fmt.Println(userData.Password)
	checkPassword := helper.CheckPasswordHash(userData.Password, password)
	if !checkPassword {
		return user.Core{}, "", errors.New("login failed, password salah")
	}
	fmt.Println(checkPassword)
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

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}
