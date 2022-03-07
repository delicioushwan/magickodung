package authUtils

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/delicioushwan/magickodung/configs"
	"github.com/delicioushwan/magickodung/utils/httpUtils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JWTClaims struct {
	UserID uint64
	jwt.StandardClaims
}

const (
	AuthScheme = "Bearer"
)

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
	fmt.Println(ok,"OKOKOKOKOKOKOKOKO")
	if !ok {
		return 0
	}
	return token.Claims.(*JWTClaims).UserID
}

func CurrentVisitor(ctx echo.Context) (uint64, error) {
	cookie, err := ctx.Cookie("1P_AS")
	if err != nil {
		return 0, httpUtils.NewInternalServerError(err)
	}
	visitorToken, err := strconv.ParseUint(DecryptAES([]byte(config.Secret), cookie.Value), 10, 64)
	if err != nil {
		return 0, httpUtils.NewInternalServerError(err)
	}
	return visitorToken, nil
}

func CurrentUserId(ctx echo.Context) (uint64) {
	var userId uint64
	if userId = CurrentUser(ctx); userId == 0 {
		fmt.Println("token??", userId)
		userId, _ = CurrentVisitor(ctx)
	}
	fmt.Println("what ??????", userId)
	return userId
}

func CurrentAuthUserId(ctx echo.Context) (uint64, error) {
	var userId uint64
	if userId = CurrentUser(ctx); userId == 0 {
		return 0, httpUtils.NewUnauthorized("회원만 이용가능")
	}

	return userId, nil
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
