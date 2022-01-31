package auth

import (
	"net/http"
	"time"

	"github.com/delicioushwan/magickodung/repository/auth"
	"github.com/delicioushwan/magickodung/utils/httpUtils"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Repo auth.AuthInterface
}

func NewAuthControllers(aurepo auth.AuthInterface) *AuthController {
	return &AuthController{
		Repo: aurepo,
	}
}

func (authcon AuthController) LoginAuthCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := LoginRequest{}
		if err := c.Bind(&req); err != nil {
			return httpUtils.NewBadRequest(err)
		}
		checkedUser, err := authcon.Repo.Login(req.Account, req.Pwd)

		if err != nil {
			return httpUtils.NewInternalServerError(err)
		}

		if err != nil || checkedUser.UserId != 0 {
			if req.Account != "" && req.Pwd != "" {
				token, err := CreateTokenAuth(checkedUser.UserId)
				if err != nil {
					return httpUtils.NewInternalServerError(err)
				}
				return c.JSON(
					http.StatusOK, map[string]interface{}{
						"message": "Successful Operation",
						"token":   token,
					},
				)
			}
			return httpUtils.NewBadRequest(err)
		} else {
			return echo.ErrNotFound
		}

	}
}

func CreateTokenAuth(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("RAHASIA"))
}
