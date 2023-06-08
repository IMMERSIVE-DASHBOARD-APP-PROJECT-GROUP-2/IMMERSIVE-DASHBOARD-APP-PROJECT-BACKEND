package service

import (
	"errors"
	"fmt"
	"regexp"
	"unicode"

	"github.com/DASHBOARDAPP/features/user"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

// Create implements user.UserServiceInterface.
func (service *userService) Create(user user.Core) error {
	errValidate := service.validate.Struct(user)
	if errValidate != nil {
		return errValidate
	}

	// Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Set the hashed password to the user input
	user.Password = string(hashedPassword)

	errInsert := service.userData.Insert(user)
	return errInsert
}

// Create implements user.UserServiceInterface.

// UpdateUserById implements user.UserServiceInterface.
func (service *userService) UpdateUserById(id string, userInput user.Core) error {
	// Mengatur ulang validator
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

	// Cek apakah pengguna mengirimkan data kosong untuk semua bidang
	if userInput.Name == "" && userInput.Phone == "" && userInput.Email == "" && userInput.Password == "" {
		return errors.New("error validation: Data tidak boleh kosong")
	}

	// Validasi email harus format email
	emailFormat := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if updatedInput.Email != "" && !emailFormat.MatchString(updatedInput.Email) {
		return errors.New("error validation: Format email tidak valid")
	}

	// Validasi panjang password minimal 8 karakter
	if updatedInput.Password != "" && len(updatedInput.Password) < 8 {
		return errors.New("error validation: Password harus memiliki panjang minimal 8 karakter")
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
			return errors.New("error validation: Password harus kombinasi huruf besar, huruf kecil, dan angka")
		}
	}

	errUpdate := service.userData.UpdateUserById(id, userInput)
	return errUpdate
}

// Delete implements user.UserServiceInterface.
func (service *userService) Delete(userID int, loggedInUserID int) error {
	// Periksa peran pengguna yang sedang login
	role, err := service.userData.GetRoleByID(loggedInUserID)
	if err != nil {
		return errors.New("Gagal mendapatkan peran pengguna yang sedang login")
	}

	// Hanya izinkan pengguna dengan peran "admin" untuk menghapus pengguna
	if role != user.Admin {
		return errors.New("Hanya admin yang dapat menghapus pengguna")
	}
	// Panggil metode Delete di repo
	err = service.userData.Delete(userID)
	if err != nil {
		return err
	}

	return nil
}

func (service *userService) Update(userID int, updatedUser user.Core, loggedInUserID int) error {
	role, err := service.userData.GetRoleByID(loggedInUserID)
	if err != nil {
		return errors.New("Gagal mendapatkan peran pengguna yang sedang login")
	}

	if role != user.Admin {
		return errors.New("Hanya admin yang dapat memperbarui pengguna")
	}

	if updatedUser.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.New("Gagal menghash password")
		}
		updatedUser.Password = string(hashedPassword)
	}

	// Update the user in the repository
	err = service.userData.Update(userID, updatedUser)
	if err != nil {
		return err
	}

	return nil
}

func (service *userService) GetRoleByID(userID int) (user.UserRole, error) {
	role, err := service.userData.GetRoleByID(userID)
	if err != nil {
		return "", err
	}

	return user.UserRole(role), nil
}

// func (service *userService) Create(user user.Core, loggedInUserID int) error {
// 	role, err := service.userData.GetRoleByID(loggedInUserID)
// 	if err != nil {
// 		return errors.New("Gagal mendapatkan role pengguna yang masuk")
// 	}

// 	if role != "admin" && user.Team != "manager" {
// 		return errors.New("Hanya admin yang dapat membuat pengguna")
// 	}

// 	// Generate hashed password
// 	hashedPassword, err := helper.HashPassword(user.Password)
// 	if err != nil {
// 		return err
// 	}

// 	// Set hashed password to user
// 	user.Password = hashedPassword

// 	// Call repository to insert user
// 	err = service.userData.Insert(user)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// GetAllUser implements user.UserServiceInterface.
func (service *userService) GetAllUser(keyword string) ([]user.Core, error) {
	data, err := service.userData.GetAllUser(keyword)
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
