package service

import (
	"errors"
	"fmt"
	"regexp"
	"unicode"

	"github.com/DASHBOARDAPP/features/user"
	"github.com/DASHBOARDAPP/helper"
	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

func (service *userService) Create(user *user.Core) error {
	// Lakukan validasi jika hanya admin atau manager yang dapat menambahkan pengguna
	if user.Role != helper.NewUserRole("admin") && user.Team != helper.NewUserTeam("manager") {
		return fmt.Errorf("only admin and manager can add users")
	}
	// Insert the user data into the database
	err := service.userData.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

// GetAllUser implements user.UserServiceInterface.
func (service *userService) GetAllUser() ([]user.Core, error) {
	data, err := service.userData.GetAllUser()
	return data, err
}

// Login implements user.UserServiceInterface.
func (service *userService) Login(email string, password string) (user.Core, string, error) {
	//validasi email tidak boleh kosong
	if email == "" || password == "" {
		return user.Core{}, "", errors.New("error validation: email, password harus diisi")
	}
	// Validasi email harus format email
	emailformat := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailformat.MatchString(email) {
		return user.Core{}, "", errors.New("error validation: format email tidak valid")
	}
	// Validasi panjang password minimal 8 kata
	if len(password) < 8 {
		return user.Core{}, "", errors.New("error validation: password harus memiliki panjang minimal 8 karakter")
	}
	// Validasi password kombinasi Huruf Besar, Huruf Kecil, Angka,
	hasUppercase := false
	hasLowercase := false
	hasDigit := false
	for _, ch := range password {
		if unicode.IsUpper(ch) {
			hasUppercase = true
		} else if unicode.IsLower(ch) {
			hasLowercase = true
		} else if unicode.IsDigit(ch) {
			hasDigit = true
		}
	}
	if !hasUppercase || !hasLowercase || !hasDigit {
		return user.Core{}, "", errors.New("error validation: password harus kombinasi huruf besar, huruf kecil, dan angka")
	}
	dataLogin, token, err := service.userData.Login(email, password)
	return dataLogin, token, err
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}
