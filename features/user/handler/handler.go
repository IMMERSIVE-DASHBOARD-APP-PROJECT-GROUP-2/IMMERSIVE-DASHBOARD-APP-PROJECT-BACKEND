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

func (handler *UserHandler) Login(c echo.Context) error {
	loginInput := AuthRequest{}
	errBind := c.Bind(&loginInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	userData, token, err := handler.userService.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "login failed") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error login,"+err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("login success", map[string]any{
		"token": token,
		"email": userData.Email,
		"id":    userData.Id,
		"role":  userData.Role,
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
