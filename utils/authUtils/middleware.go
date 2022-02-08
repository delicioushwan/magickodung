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
			_, err := c.Cookie("1P_AS")
			
			if err != nil {
				fmt.Println("anonymousToken Middleware Error : ",err)
				cookie := new(http.Cookie)
				cookie.Name = "1P_AS"
				cookie.Value = strconv.FormatInt(time.Now().UnixMilli(), 10)
				cookie.HttpOnly = true
				cookie.Path= "/"
				cookie.Expires = time.Now().Add(24 * 365 * time.Hour)
				c.SetCookie(cookie)
			}

			return next(c)
		}
	}
	
}

	
