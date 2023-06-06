package service

import (
	"errors"

	"github.com/DASHBOARDAPP/features/user"
	"github.com/DASHBOARDAPP/features/user/handler"
	"github.com/DASHBOARDAPP/helper"
	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

func (service *userService) GetRoleByID(userID int) (user.UserRole, error) {
	role, err := service.userData.GetRoleByID(userID)
	if err != nil {
		return "", err
	}

	return user.UserRole(role), nil
}

func (service *userService) Create(user user.Core, loggedInUserID int) error {
	role, err := service.userData.GetRoleByID(loggedInUserID)
	if err != nil {
		return errors.New("Gagal mendapatkan role pengguna yang masuk")
	}

	if role != "admin" {
		return errors.New("Hanya admin yang dapat membuat pengguna")
	}

	// Generate hashed password
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}

	// Set hashed password to user
	user.Password = hashedPassword

	// Call repository to insert user
	err = service.userData.Insert(user)
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

	loginInput := handler.AuthRequest{
		Email:    email,
		Password: password,
	}

	errValidate := service.validate.Struct(loginInput)
	if errValidate != nil {
		return user.Core{}, "", errors.New(errValidate.Error())
	}

	// Melakukan login
	dataLogin, token, err := service.userData.Login(email, password)
	return dataLogin, token, err
	//validasi email dan password tidak boleh kosong
	// if email == "" || password == "" {
	// 	return user.Core{}, "", errors.New("error validation: email, password harus diisi")
	// }
	// Validasi email harus format email
	// emailformat := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	// if !emailformat.MatchString(email) {
	// 	return user.Core{}, "", errors.New("error validation: format email tidak valid")
	// }
	// Validasi panjang password minimal 8 kata
	// if len(password) < 8 {
	// 	return user.Core{}, "", errors.New("error validation: password harus memiliki panjang minimal 8 karakter")
	// }
	// Validasi password kombinasi Huruf Besar, Huruf Kecil, Angka,
	// hasUppercase := false
	// hasLowercase := false
	// hasDigit := false
	// for _, ch := range password {
	// 	if unicode.IsUpper(ch) {
	// 		hasUppercase = true
	// 	} else if unicode.IsLower(ch) {
	// 		hasLowercase = true
	// 	} else if unicode.IsDigit(ch) {
	// 		hasDigit = true
	// 	}
	// }
	// if !hasUppercase || !hasLowercase || !hasDigit {
	// 	return user.Core{}, "", errors.New("error validation: password harus kombinasi huruf besar, huruf kecil, dan angka")
	// }
	// dataLogin, token, err := service.userData.Login(email, password)
	// return dataLogin, token, err
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}
