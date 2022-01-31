package user

import (
	"net/http"
	"strconv"

	"github.com/delicioushwan/magickodung/entities"
	"github.com/delicioushwan/magickodung/repository/user"
	"github.com/delicioushwan/magickodung/utils/auth"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

type UsersController struct {
	Repo user.UserInterface
}

func NewUsersControllers(usrep user.UserInterface) *UsersController {
	return &UsersController{Repo: usrep}
}

// POST /user/register
func (uscon UsersController) PostUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		newUserReq := UserCommonRequestFormat{}

		if err := c.Bind(&newUserReq); err != nil {
			return echo.ErrBadRequest
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(newUserReq.Pwd), 14)
		newUser := entities.User{
			Account:     newUserReq.Account,
			Pwd: string(hash),
		}

		u, err := uscon.Repo.Create(newUser)
		if err != nil {
			return echo.ErrInternalServerError
		}
		
		token, err := auth.MakeJWTToken(u.UserId)
		if err != nil {
			return echo.ErrInternalServerError
		}
	

		return c.JSON(http.StatusOK, token)
	}

}

// GET /users/:id
func (uscon UsersController) GetUserCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return echo.ErrBadRequest
		}

		user, err := uscon.Repo.Get(id)
		if err != nil {
			return echo.ErrNotFound
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    user,
		})
	}

}

