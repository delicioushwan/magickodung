package authUtils

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/delicioushwan/magickodung/utils/httpUtils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func NewJWTMiddleware(secret string) echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			accessTokenCookie, Uerr := c.Cookie("mg_at")
			if Uerr != nil {
				return httpUtils.NewUnauthorized("로그인 필요!")
			}

			userId, parseErr := parseToken(accessTokenCookie.Value)
			if parseErr != nil {
				return httpUtils.NewForbiden(parseErr)
			}
			c.Set("user", userId)
		return next(c)
	}
	}}

func NewTokenMiddleware(secret string) echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, Cerr := c.Cookie("1P_AS")
			if Cerr != nil {
				fmt.Println("anonymousToken Middleware Error : ", Cerr)
				tokenValue := strconv.FormatInt(time.Now().UnixMilli(), 10)
				encryptoValue := EncryptAES([]byte(secret), tokenValue)
		
				cookie := new(http.Cookie)
				cookie.Name = "1P_AS"
				cookie.Value = encryptoValue
				cookie.HttpOnly = true
				cookie.Path= "/"
				cookie.Expires = time.Now().Add(24 * 3 * time.Hour)
				c.SetCookie(cookie)
			}
	
			accessTokenCookie, Uerr := c.Cookie("mg_at")
			if Uerr == nil {
				userId, parseErr := parseToken(accessTokenCookie.Value)
				if parseErr != nil {
					cookie := new(http.Cookie)
					cookie.Name = "mg_at"
					cookie.Value =  ""
					cookie.HttpOnly = true
					cookie.Path= "/"
					cookie.MaxAge = 0
					c.SetCookie(cookie)
					return httpUtils.NewForbiden(parseErr)
				}
				c.Set("user", userId)
				return next(c)
			}
		return next(c)
		}	
	}
}

func parseToken(auth string) (uint64, error) {
	token, Perr := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, httpUtils.NewForbiden("unexpected signing method")
		}
		return []byte(config.Secret), nil
	})
	if Perr != nil {
		return 0, httpUtils.NewForbiden(Perr)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := uint64(claims["UserID"].(float64))
		return userId, nil

	}
	return 0, httpUtils.NewForbiden(nil)
}
	
