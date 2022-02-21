package authUtils

import (
	"net/http"
	"time"

	"github.com/delicioushwan/magickodung/configs"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JWTClaims struct {
	UserID uint64
	jwt.StandardClaims
}

var config = configs.GetConfig()

const expires time.Duration = 24 * 30 * time.Hour

func MakeJWTToken(userID uint64) (string, error) {
	c := &JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expires).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(config.Secret))
}

func CurrentUser(ctx echo.Context) uint64 {
	token, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return 0
	}
	return token.Claims.(*JWTClaims).UserID
}

type visitor struct {
	visitorId uint64
}
func CurrentVisitor(ctx echo.Context) uint64 {
	visitor := ctx.Get("1P_AS").(visitor)

	return visitor.visitorId

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
