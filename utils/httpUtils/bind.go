package httpUtils

import "github.com/labstack/echo/v4"

func BindAndValidate(ctx echo.Context, v interface{}) error {
	if err := ctx.Bind(v); err != nil {
		return err
	}
	if err := ctx.Validate(v); err != nil {
		return err
	}
	return nil
}
