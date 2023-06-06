package handler

import (
	"net/http"
	"strings"

	"github.com/DASHBOARDAPP/features/user"
	"github.com/DASHBOARDAPP/helper"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}
func (handler *UserHandler) CreateUser(c echo.Context) error {
	userInput := UserRequest{}

	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	// Pengecekan peran pengguna
	if userInput.Role != "admin" && userInput.Role != "manager" {
		return c.JSON(http.StatusForbidden, helper.FailedResponse("Only admin and manager can add users"))
	}

	userCore := &user.Core{
		Name:     userInput.Name,
		Phone:    userInput.Phone,
		Email:    userInput.Email,
		Password: userInput.Password,
		Role:     user.UserRole(userInput.Role),
		Status:   user.UserStatus(userInput.Status), Team: user.UserTeam(userInput.Team),
	}

	err := handler.userService.Create(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data"+err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success insert data"))
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
	// mapping dari request ke core
	userCore := user.Core{
		Name:     userInput.Name,
		Phone:    userInput.Phone,
		Email:    userInput.Email,
		Password: userInput.Password,
	}
	err := handler.userService.UpdateUserById(id, userCore)
	if err != nil {
		if strings.Contains(err.Error(), "failed updated data user") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("error updated data user, row affected = 0"))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success updated data user"))
}
