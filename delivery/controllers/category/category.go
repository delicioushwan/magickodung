package category

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


func GetCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		category := []Category{}
		category = append(category,
			Category{Key: 0, Value: "패션"},
			Category{Key: 1, Value: "고민"},
			Category{Key: 2, Value: "심심"},
			Category{Key: 3, Value: "연애"},
			Category{Key: 4, Value: "암거나"},
		)
		return c.JSON(http.StatusOK, category)
	}
}