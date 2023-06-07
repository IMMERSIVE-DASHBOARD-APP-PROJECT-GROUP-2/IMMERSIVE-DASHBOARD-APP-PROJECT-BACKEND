package routes

import (
	"github.com/DASHBOARDAPP/app/middlewares"
	_classData "github.com/DASHBOARDAPP/features/class/data"
	_classHandler "github.com/DASHBOARDAPP/features/class/handler"
	_classService "github.com/DASHBOARDAPP/features/class/service"
	_userData "github.com/DASHBOARDAPP/features/user/data"
	_userHandler "github.com/DASHBOARDAPP/features/user/handler"
	_userService "github.com/DASHBOARDAPP/features/user/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(e *echo.Echo, db *gorm.DB) {
	userData := _userData.New(db)
	classData := _classData.New(db)

	userService := _userService.New(userData)
	classService := _classService.New(classData)

	userHandlerAPI := _userHandler.New(userService)
	classHandlerAPI := _classHandler.New(classService)
	// // Register middleware
	jwtMiddleware := middlewares.JWTMiddleware()

	// // User Routes
	e.POST("/login", userHandlerAPI.Login)
	e.GET("/users", userHandlerAPI.GetAllUser, middlewares.JWTMiddleware())
	e.PUT("/users/:id", userHandlerAPI.UpdateUserById, middlewares.JWTMiddleware())
	e.POST("/users/role", userHandlerAPI.CreateUser, jwtMiddleware)
	e.PUT("/users/role/:id", userHandlerAPI.UpdateUser, jwtMiddleware)
	e.DELETE("/users/role/:id", userHandlerAPI.DeleteUser, jwtMiddleware)

	//EndPointBook
	e.POST("/classes", classHandlerAPI.CreateClass, middlewares.JWTMiddleware())
	e.DELETE("/classes/:id", classHandlerAPI.DeleteClass, middlewares.JWTMiddleware())

}
