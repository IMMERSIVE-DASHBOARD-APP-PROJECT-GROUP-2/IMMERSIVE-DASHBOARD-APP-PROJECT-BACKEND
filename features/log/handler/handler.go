package handler

import (
	"net/http"
	"strconv"

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
func (handler *LogHandler) GetLogsByMenteeID(c echo.Context) error {
	// Get mentee ID from request path parameter
	menteeID := c.Param("menteeID")

	// Convert menteeID to uint
	menteeIDUint, err := strconv.ParseUint(menteeID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid mentee ID"))
	}

	// Get logs by mentee ID
	logs, err := handler.logService.GetLogsByMenteeID(uint(menteeIDUint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	// Convert logs to string slice
	logDescriptions := make([]string, len(logs))
	for i, log := range logs {
		logDescriptions[i] = log.Description
	}

	// Return success response with logs
	response := helper.SuccessResponse("Logs retrieved successfully")
	response["data"] = logDescriptions
	return c.JSON(http.StatusOK, response)
}
func NewLogHandler(logService log.LogServiceInterface) *LogHandler {
	return &LogHandler{
		logService: logService,
	}
}
