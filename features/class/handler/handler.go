package handler

import (
	"net/http"
	"strings"

	"github.com/DASHBOARDAPP/app/middlewares"
	"github.com/DASHBOARDAPP/features/class"
	"github.com/DASHBOARDAPP/helper"
	"github.com/labstack/echo/v4"
)

type ClassHandler struct {
	classService class.ClassServiceInterface
}

func New(service class.ClassServiceInterface) *ClassHandler {
	return &ClassHandler{
		classService: service,
	}
}

func (handler *ClassHandler) CreateClass(c echo.Context) error {

	// Mendapatkan ID pengguna yang sedang login
	loggedInUserID, err := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Gagal mendapatkan ID pengguna"))
	}

	classInput := ClassRequest{}
	// bind, membaca data yg dikirimkan client via request body
	errBind := c.Bind(&classInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	// mapping input class dan id login ke core
	classCore := class.Core{
		Name:   classInput.Name,
		UserID: uint(loggedInUserID),
	}

	err = handler.classService.CreateClass(classCore)
	if err != nil {
		if strings.Contains(err.Error(), "failed inserted data class") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("error inserted data class, row affected = 0"))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error inserted data class"+err.Error()))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success inserted data class"))
}
