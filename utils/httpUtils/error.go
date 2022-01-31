package httpUtils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewUnauthorized() error {
	return NewError(http.StatusUnauthorized, "auth required")
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
		msg = err.Error()
	}
	return NewError(http.StatusInternalServerError, msg)
}

func NewBadRequest(err error) error {
	msg := "bad request"
	if err != nil {
		msg = err.Error()
	}

	return NewError(http.StatusBadRequest, msg)
}


func NewError(statusCode int, msg string) error {
	return echo.NewHTTPError(statusCode, &Error{
		Errors: map[string]interface{}{
			"body": msg,
		},
	})
}
