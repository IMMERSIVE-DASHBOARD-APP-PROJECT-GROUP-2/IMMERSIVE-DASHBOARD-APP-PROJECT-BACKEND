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

// GetUserByID implements user.UserDataInterface.
func (repo *userQuery) GetUserByID(userID uint) (*user.Core, error) {
	var userData user.Core
	err := repo.db.First(&userData, userID).Error
	if err != nil {
		return nil, err
	}
	return &userData, nil
}

// UpdateUserById implements user.UserDataInterface.
func (repo *userQuery) UpdateUserById(id string, userInput user.Core) error {
	// Mencari pengguna berdasarkan ID
	var userData User
	tx := repo.db.First(&userData, id)
	// Hash password sebelum disimpan
	hashedPassword, err := helper.HashPassword(userInput.Password)
	if err != nil {
		return err
	}
	// Mengganti password dengan hashed password
	userInput.Password = hashedPassword
	// Mengupdate data pengguna berdasarkan ID dari userInputGorm
	px := repo.db.Model(&userData).Updates(CoreToUpdateModel(userInput))
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

// Delete implements user.UserDataInterface.
func (repo *userQuery) Delete(userID int) error {
	var user User

	if err := repo.db.First(&user, userID).Error; err != nil {
		return err
	}

	// Ubah status pengguna menjadi "non_active"
	user.Status = NonActive
	// Simpan perubahan status pengguna
	if err := repo.db.Save(&user).Error; err != nil {
		return err
	}

	// Hapus pengguna dari basis data
	if err := repo.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

// Update implements user.UserDataInterface.
func (repo *userQuery) Update(userID int, updatedUser user.Core) error {
	userData := ModelToCore(&updatedUser)
	userData.ID = uint(userID)

	tx := repo.db.Model(&User{}).Where("id = ?", userID).Updates(userData)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("Gagal memperbarui data, tidak ada baris yang terpengaruh")
	}

	return nil
}

// GetRoleByID implements user.UserDataInterface.
func (repo *userQuery) GetRoleByID(userID int) (user.UserRole, error) {
	var u User
	err := repo.db.Where("id = ?", userID).First(&u).Error
	if err != nil {
		return "", err
	}

	return user.UserRole(u.Role), nil
}

// Insert implements user.UserDataInterface.
func (repo *userQuery) Insert(user user.Core) error {
	// Konversi data pengguna dari Core ke Model
	userModel := ModelToCore(&user)

	// Lakukan operasi penambahan pengguna ke basis data menggunakan GORM
	err := repo.db.Create(userModel).Error
	if err != nil {
		// Error saat menambahkan pengguna ke basis data
		return err
	}

	return nil
}

// GetAllUser implements user.UserDataInterface.
func (repo *userQuery) GetAllUser(keyword string) ([]user.Core, error) {
	var userData []User
	tx := repo.db
	if keyword != "" {
		tx = tx.Where("id LIKE ?", "%"+keyword+"%").
			Or("name LIKE ?", "%"+keyword+"%").
			Or("email LIKE ?", "%"+keyword+"%").
			Or("team LIKE ?", "%"+keyword+"%").
			Or("role LIKE ?", "%"+keyword+"%").
			Or("status LIKE ?", "%"+keyword+"%")
	}
	tx = tx.Find(&userData)
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
