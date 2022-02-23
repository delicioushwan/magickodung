package category

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/delicioushwan/magickodung/utils/httpUtils"
	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
)

func TestQuestion(t *testing.T) {
	ec := echo.New()
	ec.Validator = httpUtils.NewValidator()

	ec.GET("/category", GetCategory())

	t.Run("카테고리 불러오기 시도 -> 성공", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{})

		uri := "/category"
		req := httptest.NewRequest(http.MethodGet, uri, bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()

		ec.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})
}