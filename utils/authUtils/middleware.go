package authUtils

import (
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
