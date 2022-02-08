package routes

import (
	"net/http"

	"github.com/delicioushwan/magickodung/delivery/controllers/user"

	"github.com/labstack/echo/v4"
)

func Session (c echo.Context) error {
		return c.JSON(http.StatusOK, "nice")
	}

func RegisterPath(
	e *echo.Echo,
	uctrl *user.UsersController,
) {

	e.GET("/session", Session)

	usersGroup := e.Group("/users")
	usersGroup.POST("/signup", uctrl.Signup())
	usersGroup.POST("/login", uctrl.Login())
}
