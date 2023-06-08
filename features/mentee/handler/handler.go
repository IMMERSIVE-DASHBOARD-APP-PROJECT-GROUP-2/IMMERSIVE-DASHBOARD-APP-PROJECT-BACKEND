package handler

import (
	"net/http"
	"strings"

	"github.com/DASHBOARDAPP/features/mentee"
	"github.com/DASHBOARDAPP/helper"
	"github.com/labstack/echo/v4"
)

type MenteeHandler struct {
	menteeService mentee.MenteeServiceInterface
}

func New(service mentee.MenteeServiceInterface) *MenteeHandler {
	return &MenteeHandler{
		menteeService: service,
	}
}

func (handler *MenteeHandler) CreateMentee(c echo.Context) error {

	menteeInput := MenteeRequest{}
	// bind, membaca data yg dikirimkan client via request body
	errBind := c.Bind(&menteeInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	// mapping input ke datacore
	menteeCore := mentee.Core{
		// ClassId
		ClassID: menteeInput.ClassID,
		// MainData
		Name:        menteeInput.Name,
		Phone:       menteeInput.Phone,
		Email:       menteeInput.Email,
		Address:     menteeInput.Address,
		HomeAddress: menteeInput.HomeAddress,
		Telegram:    menteeInput.Telegram,
		Gender:      mentee.MenteeGender(menteeInput.Gender),
		// EducationData
		Category:  mentee.MenteeCategory(menteeInput.Category),
		Major:     menteeInput.Major,
		Graduated: menteeInput.Graduated,
		// EmergencyData
		EmergencyName:   menteeInput.EmergencyName,
		EmergencyStatus: mentee.EmergencyStatus(menteeInput.EmergencyStatus),
		EmergencyPhone:  menteeInput.EmergencyPhone,
		// Status
		Status: mentee.MenteeStatus(menteeInput.Status),
	}

	err := handler.menteeService.CreateMentee(menteeCore)
	if err != nil {
		if strings.Contains(err.Error(), "failed inserted data mentee") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("error inserted data mentee, row affected = 0"))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error inserted data mentee"+err.Error()))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success inserted data mentee"))
}

func (handler *MenteeHandler) GetAllMentee(c echo.Context) error {
	//Memanggil function di Service logic via interface
	results, err := handler.menteeService.GetAllMentee()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data mentee"))
	}
	var menteeResponse []MenteeResponse
	for _, value := range results {
		menteeResponse = append(menteeResponse, MenteeResponse{
			Id:       value.Id,
			ClassID:  value.ClassID,
			Name:     value.Name,
			Gender:   MenteeGender(value.Gender),
			Category: MenteeCategory(value.Category),
			Status:   MenteeStatus(value.Status),
		})
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data class", menteeResponse))
}
