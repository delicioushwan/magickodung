package authUtils

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"github.com/delicioushwan/magickodung/configs"
)

type JWTClaims struct {
	UserID uint
	jwt.StandardClaims
}

var config = configs.GetConfig()

const expires time.Duration = 240 * time.Hour

func MakeJWTToken(userID uint) (string, error) {
	c := &JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expires).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(config.Secret))
}

// CurrentUser returns current user id which stored at echo.Context if exist, otherwise returns 0.
func CurrentUser(ctx echo.Context) uint {
	token, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return 0
	}
	return token.Claims.(*JWTClaims).UserID
}

func SetAuthCookie(c echo.Context, token string) {
	cookie := new(http.Cookie)
	cookie.Name = "access_token"
	cookie.Value = token
	cookie.HttpOnly = true
	cookie.Path= "/"
	cookie.Expires = time.Now().Add(24 * 30 * time.Hour)
	c.SetCookie(cookie)
}
