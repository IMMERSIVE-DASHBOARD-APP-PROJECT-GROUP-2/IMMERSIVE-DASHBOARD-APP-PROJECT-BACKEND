package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/DASHBOARDAPP/app/middlewares"
	"github.com/DASHBOARDAPP/features/user"
	"github.com/DASHBOARDAPP/helper"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(handler user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: handler,
	}
}

func (handler *UserHandler) Login(c echo.Context) error {
	// Memeriksa apakah email dan password inputan dapat di bind
	loginInput := AuthRequest{}
	errBind := c.Bind(&loginInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}
	// Memeriksa apakah email & password telah diinputkan di database
	userData, token, err := handler.userService.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "login failed") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
			// Memeriksa validasi di database dan validasi lainnya
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("login success", map[string]any{
		"user_id": userData.Id,
		"email":   userData.Email,
		"role":    userData.Role,
		"token":   token,
	}))
}

func (handler *UserHandler) GetAllUser(c echo.Context) error {
	//Memanggil function di Service logic via interface
	results, err := handler.userService.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data user"))
	}
	var userResponse []UserResponse
	for _, value := range results {
		userResponse = append(userResponse, UserResponse{
			Id:     value.Id,
			Name:   value.Name,
			Email:  value.Email,
			Team:   UserTeam(value.Team),
			Role:   UserRole(value.Role),
			Status: UserStatus(value.Status),
		})
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data user", userResponse))
}

func (handler *UserHandler) CreateUser(c echo.Context) error {
	// Mendapatkan data pengguna dari permintaan
	userInput := UserRequest{}
	err := c.Bind(&userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	// Mendapatkan ID pengguna yang login
	loggedInUserID, err := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("gagal mendapatkan ID pengguna"))
	}

	// Membuat data pengguna
	user := user.Core{
		Name:     userInput.Name,
		Phone:    userInput.Phone,
		Email:    userInput.Email,
		Password: userInput.Password,
		Role:     user.UserRole(userInput.Role),
		Status:   user.UserStatus(userInput.Status),
		Team:     user.UserTeam(userInput.Team),
	}

	// Memanggil service untuk menambahkan pengguna
	err = handler.userService.Create(user, loggedInUserID)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("gagal insert data"+err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("sukses insert data"))
}

func (handler *UserHandler) UpdateUser(c echo.Context) error {
	// Get the ID of the user to be updated
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("ID pengguna tidak valid"))
	}

	// Get the updated user data from the request
	userInput := UserRequest{}
	err = c.Bind(&userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Gagal mengambil data pengguna"))
	}
	// Mendapatkan ID pengguna yang login
	loggedInUserID, err := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("gagal mendapatkan ID pengguna"))
	}

	// Membuat objek user.Core dari userInput
	updatedUser := user.Core{
		Name:     userInput.Name,
		Phone:    userInput.Phone,
		Email:    userInput.Email,
		Password: userInput.Password,
		Role:     user.UserRole(userInput.Role),
		Status:   user.UserStatus(userInput.Status),
		Team:     user.UserTeam(userInput.Team),
	}

	// Memperbarui pengguna
	err = handler.userService.Update(userID, updatedUser, loggedInUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Gagal memperbarui pengguna"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Berhasil memperbarui pengguna"))
}

func (handler *UserHandler) DeleteUser(c echo.Context) error {
	// Mendapatkan ID pengguna yang akan dihapus
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("ID pengguna tidak valid"))
	}

	// Mendapatkan ID pengguna yang login
	loggedInUserID, err := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Gagal mendapatkan ID pengguna"))
	}

	// Memanggil service untuk menghapus pengguna
	err = handler.userService.Delete(userID, loggedInUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Gagal menghapus pengguna"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Pengguna berhasil dihapus"))
}

func (handler *UserHandler) UpdateUserById(c echo.Context) error {
	// Mendapatkan nilai ID dari parameter di URL
	id := c.Param("id")

	// Bind data pengguna yang baru dari request body
	userInput := UserRequest{}
	// bind, membaca data yg dikirimkan client via request body
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data user"))
	}
	// Mendapatkan ID pengguna yang sedang login
	loggedInUserID, err := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Gagal mendapatkan ID pengguna"))
	}

	// Jika ID pengguna yang sedang login tidak sama dengan ID pengguna yang ingin diubah
	if strconv.Itoa(loggedInUserID) != id {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Anda hanya dapat mengubah profil sendiri"))
	}

	// mapping dari request ke core
	userCore := user.Core{
		Name:     userInput.Name,
		Phone:    userInput.Phone,
		Email:    userInput.Email,
		Password: userInput.Password,
	}
	err = handler.userService.UpdateUserById(id, userCore)
	if err != nil {
		if strings.Contains(err.Error(), "failed updated data user") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("error updated data user, row affected = 0"))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success updated data user"))
}
