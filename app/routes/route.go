package routes

import (
	"github.com/DASHBOARDAPP/app/middlewares"
	_userData "github.com/DASHBOARDAPP/features/user/data"
	_userHandler "github.com/DASHBOARDAPP/features/user/handler"
	_userService "github.com/DASHBOARDAPP/features/user/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(e *echo.Echo, db *gorm.DB) {
	userData := _userData.New(db)

	userService := _userService.New(userData)

	userHandlerAPI := _userHandler.New(userService)

	// Register middleware
	// jwtMiddleware := middlewares.JWTMiddleware()

	// User Routes

	e.POST("/users", userHandlerAPI.CreateUser)
	e.POST("/login", userHandlerAPI.Login)
	e.GET("/users", userHandlerAPI.GetAllUser, middlewares.JWTMiddleware())

}
