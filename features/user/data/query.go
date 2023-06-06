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

// UpdateUserById implements user.UserDataInterface.
func (repo *userQuery) UpdateUserById(id string, userInput user.Core) error {
	// Mencari pengguna berdasarkan ID
	var userData User
	tx := repo.db.First(&userData, id)
	// Mengupdate data pengguna berdasarkan ID dari userInputGorm
	px := repo.db.Model(&userData).Updates(CoreToModel(userInput))
	if tx.Error != nil {
		return tx.Error
	} else if px.Error != nil {
		return px.Error
	}

	// Menyimpan perubahan data pengguna dari Input ke database
	tx = repo.db.Save(&userData)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("Updated Failed, row affected = 0")
	}
	return nil
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

// GetAllUser implements user.UserDataInterface.
func (repo *userQuery) GetAllUser() ([]user.Core, error) {
	var userData []User
	// Mencari data user di database
	tx := repo.db.Find(&userData)
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

	// Mencocokkan data inputan email dengan email di database
	tx := repo.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return user.Core{}, "", tx.Error
	}
	if tx.RowsAffected == 0 {
		return user.Core{}, "", errors.New("login failed, email salah")
	}
	// Mencocokkan data inputan password dengan password yang telah di hashing di database
	checkPassword := helper.CheckPasswordHash(userData.Password, password)
	if !checkPassword {
		return user.Core{}, "", errors.New("login failed, password salah")
	}
	// Memastikan status pengguna Active
	if userData.Status == NonActive {
		return user.Core{}, "", errors.New("Hanya user dengan status aktif yang dapat melakukan login")
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
