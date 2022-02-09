package authUtils

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewJWTMiddleware(secret string) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(
		middleware.JWTConfig{
			Claims:      &JWTClaims{},
			SigningKey:  []byte(secret),
		},
	)
}

func NewAnonymousTokenMiddleware(secret string) echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			val, err := c.Cookie("1P_AS")
			if err != nil {
				fmt.Println("anonymousToken Middleware Error : ",err)
				tokenValue := strconv.FormatInt(time.Now().UnixMilli(), 10)
				encryptoValue := EncryptAES([]byte(secret), tokenValue)

				fmt.Println(encryptoValue)
				cookie := new(http.Cookie)
				cookie.Name = "1P_AS"
				cookie.Value = encryptoValue
				cookie.HttpOnly = true
				cookie.Path= "/"
				cookie.Expires = time.Now().Add(24 * 365 * time.Hour)
				c.SetCookie(cookie)
			} else {
				c.Set("1P_AS",DecryptAES([]byte(secret), val.Value))
			}

			return next(c)
		}
	}
	
}

	
