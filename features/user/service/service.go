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

// UpdateUserById implements user.UserServiceInterface.
func (service *userService) UpdateUserById(id string, userInput user.Core) error {
	// Mengatur validator
	validate := validator.New()
	updatedInput := user.UpdatedInput{
		Name:     userInput.Name,
		Phone:    userInput.Phone,
		Email:    userInput.Email,
		Password: userInput.Password,
	}
	fmt.Println(updatedInput)
	errValidate := validate.Struct(updatedInput)
	if errValidate != nil {
		return errValidate
	}

	// Validasi email harus format email
	emailFormat := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if updatedInput.Email != "" && !emailFormat.MatchString(updatedInput.Email) {
		return errors.New("error validation: Format email tidak valid")
	}

	// Validasi panjang password minimal 8 karakter
	if updatedInput.Password != "" && len(updatedInput.Password) < 8 {
		return errors.New("error validation: password harus memiliki panjang minimal 8 karakter")
	}

	// Validasi password kombinasi huruf besar, huruf kecil, dan angka
	if updatedInput.Password != "" {
		hasUppercase := false
		hasLowercase := false
		hasDigit := false
		for _, ch := range updatedInput.Password {
			if unicode.IsUpper(ch) {
				hasUppercase = true
			} else if unicode.IsLower(ch) {
				hasLowercase = true
			} else if unicode.IsDigit(ch) {
				hasDigit = true
			}
		}
		if !hasUppercase || !hasLowercase || !hasDigit {
			return errors.New("error validation: password harus kombinasi huruf besar, huruf kecil, dan angka")
		}
	}

	errUpdate := service.userData.UpdateUserById(id, userInput)
	return errUpdate

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
	// Mengatur validator
	validate := validator.New()
	loginInput := user.LoginInput{
		Email:    email,
		Password: password,
	}
	errValidate := validate.Struct(loginInput)
	if errValidate != nil {
		return user.Core{}, "", errValidate
	}

	// Validasi email harus format email
	emailformat := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailformat.MatchString(loginInput.Email) {
		return user.Core{}, "", errors.New("error validation: Format email tidak valid")
	}

	// Validasi panjang password minimal 8 kata
	if len(loginInput.Password) < 8 {
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

	// Lanjutkan dengan proses login
	dataLogin, token, errValidate := service.userData.Login(email, password)
	return dataLogin, token, errValidate
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}
