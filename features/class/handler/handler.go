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

func (handler *ClassHandler) GetAllClass(c echo.Context) error {
	//Memanggil function di Service logic via interface
	results, err := handler.classService.GetAllClass()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data class"))
	}
	var classResponse []ClassResponse
	for _, value := range results {
		classResponse = append(classResponse, ClassResponse{
			Id:     value.Id,
			Name:   value.Name,
			UserID: value.UserID,
		})
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data class", classResponse))
}

func (handler *ClassHandler) UpdateClassById(c echo.Context) error {
	// Mendapatkan nilai ID dari parameter di URL
	id := c.Param("id")
	// Mendapatkan ID pengguna yang sedang login
	loggedInUserID, err := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Gagal mendapatkan ID pengguna"))
	}
	// Bind data pengguna yang baru dari request body
	classInput := ClassRequest{}
	// bind, membaca data yg dikirimkan client via request body
	errBind := c.Bind(&classInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data class"))
	}

	// mapping input class dan id login ke core
	classCore := class.Core{
		Name:   classInput.Name,
		UserID: uint(loggedInUserID),
	}

	err = handler.classService.UpdateClassById(id, classCore)
	if err != nil {
		if strings.Contains(err.Error(), "failed updated data class") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("error updated data class, row affected = 0"))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success updated data Class"))
}
