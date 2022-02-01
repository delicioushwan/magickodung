package user

import (
	"net/http"
	"strconv"

	"github.com/delicioushwan/magickodung/entities"
	"github.com/delicioushwan/magickodung/repository/user"
	"github.com/delicioushwan/magickodung/utils/authUtils"
	"github.com/delicioushwan/magickodung/utils/httpUtils"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

type UsersController struct {
	Repo user.UserInterface
}

func NewUsersControllers(usrep user.UserInterface) *UsersController {
	return &UsersController{Repo: usrep}
}

func (ctrl UsersController) Signup() echo.HandlerFunc {

	return func(c echo.Context) error {
		req := UserCommonRequestFormat{}

		if err := httpUtils.BindAndValidate(c, &req); err != nil {
			return httpUtils.NewBadRequest(err)
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(req.Pwd), 14)
		newUser := entities.User{
			Account: req.Account,
			Pwd: string(hash),
		}

		u, err := ctrl.Repo.Create(newUser)
		if err != nil {
			return httpUtils.NewInternalServerError(err)
		}

		token, err := authUtils.MakeJWTToken(u.UserId)
		if err != nil {
			return httpUtils.NewInternalServerError(err)
		}
	
		authUtils.SetAuthToken(c.Request(), token)

		return c.JSON(http.StatusOK, token)
	}

}


func (ctrl UsersController) GetUserCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return httpUtils.NewBadRequest(err)
		}

		user, err := ctrl.Repo.Get(id)
		if err != nil {
			return httpUtils.NewInternalServerError(err)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    user,
		})
	}

}

func (ctrl UsersController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := UserCommonRequestFormat{}
		if err := httpUtils.BindAndValidate(c, &req); err != nil {
			return httpUtils.NewBadRequest(err)
		}

		user, err := ctrl.Repo.GetByAccount(req.Account)
		if err != nil {
			return httpUtils.NewBadRequest("존재하지 않는 회원입니다. \n 아이디와 비밀번호를 확인해 주세요.")
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(req.Pwd)); err != nil {
			return httpUtils.NewBadRequest("아이디와 비밀번호를 확인해 주세요.")
		}

		token, err := authUtils.MakeJWTToken(user.UserId)
		if err != nil {
			return httpUtils.NewInternalServerError(err)
		}

		authUtils.SetAuthToken(c.Request(), token)

		return c.JSON(http.StatusOK, token)
	}
}


