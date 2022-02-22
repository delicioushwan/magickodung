package httpUtils

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewUnauthorized(msg string) error {
	if msg == "" {
		msg = "auth required"
	}
	return NewError(http.StatusUnauthorized, msg)
}

func NewStatusUnprocessableEntity(msg string) error {
	return NewError(http.StatusUnprocessableEntity, msg)
}

func NewNotFoundError(msg string) error {
	if msg == "" {
		msg = "resource not found"
	}
	return NewError(http.StatusNotFound, msg)
}

func NewInternalServerError(err error) error {
	msg := "interval server error occur"
	if err != nil {
		fmt.Println(err.Error())
	}
	return NewError(http.StatusInternalServerError, msg)
}

func NewBadRequest(err interface{}) error {
	msg := "bad request"
	if err != nil {
		switch v := err.(type) {
		case string:
			msg = v
		case error:
			msg = v.Error()
		}
	}

	return NewError(http.StatusBadRequest, msg)
}


func NewError(statusCode int, msg string) error {
	fmt.Println(msg)
	return echo.NewHTTPError(statusCode, &Error{
		Errors: map[string]interface{}{
			"body": msg,
		},
	})
}
