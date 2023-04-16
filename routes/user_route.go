package routes

import (
	"echo-mongo-api/controllers"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	e.POST("/users", controllers.CreateUser)
	e.GET("/users/:userId", controllers.GetAUser)
	e.PUT("/users/:userId", controllers.EditAUser)
	e.DELETE("/users/:userId", controllers.DeleteAUser)
	e.GET("/users", controllers.GetAllUsers)
}
