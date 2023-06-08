package handler

import (
	"net/http"

	"github.com/DASHBOARDAPP/features/log"
	"github.com/DASHBOARDAPP/helper"
	"github.com/labstack/echo/v4"
)

type LogHandler struct {
	logService log.LogServiceInterface
}

func New(service log.LogServiceInterface) *LogHandler {
	return &LogHandler{
		logService: service,
	}
}
func (handler *LogHandler) CreateLog(c echo.Context) error {
	// Bind request body to LogRequest struct
	logInput := log.Core{}
	if err := c.Bind(&logInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed to parse request body"))
	}

	// Create log
	err := handler.logService.Insert(logInput)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Log created successfully"))
}
