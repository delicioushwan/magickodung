package routes

import (
	"github.com/delicioushwan/magickodung/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(
	e *echo.Echo,
	uctrl *user.UsersController,
) {

	e.Use(middleware.RemoveTrailingSlash())

	usersGroup := e.Group("/users")
	usersGroup.POST("/signup", uctrl.Signup())
	usersGroup.POST("/login", uctrl.Login())
}
