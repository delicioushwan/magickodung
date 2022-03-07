package authUtils

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewJWTMiddleware(secret string) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(
		middleware.JWTConfig{
			TokenLookup: "cookie:access_token",
			Claims:      &JWTClaims{},
			SigningKey:  []byte(secret),
			ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
				keyFunc := func(t *jwt.Token) (interface{}, error) {
					fmt.Println("23423423RWEREWRWER@#$@#$WERWERWER",t.Method.Alg())
					if t.Method.Alg() != "HS256" {
						return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
					}
					return []byte(secret), nil
				}
		
				fmt.Println(auth,"authauthauthauthauthauthauthauthauthauth")
				token, err := jwt.Parse(auth, keyFunc)
				if err != nil {
					return nil, err
				}
				if !token.Valid {
					return nil, errors.New("invalid token")
				}
				return token, nil
			},		},
	)
}

func NewAnonymousTokenMiddleware(secret string) echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, err := c.Cookie("1P_AS")
			if err != nil {
				fmt.Println("anonymousToken Middleware Error : ",err)
				tokenValue := strconv.FormatInt(time.Now().UnixMilli(), 10)
				encryptoValue := EncryptAES([]byte(secret), tokenValue)

				cookie := new(http.Cookie)
				cookie.Name = "1P_AS"
				cookie.Value = encryptoValue
				cookie.HttpOnly = true
				cookie.Path= "/"
				cookie.Expires = time.Now().Add(24 * 365 * time.Hour)
				c.SetCookie(cookie)
			} 
			return next(c)
		}
	}
	
}

	
