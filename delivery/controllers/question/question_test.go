package question

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/delicioushwan/magickodung/configs"
	"github.com/delicioushwan/magickodung/entities"
	optionRepo "github.com/delicioushwan/magickodung/repository/option"
	questionRepo "github.com/delicioushwan/magickodung/repository/question"
	userRepo "github.com/delicioushwan/magickodung/repository/user"
	"github.com/delicioushwan/magickodung/utils"
	"github.com/delicioushwan/magickodung/utils/httpUtils"
	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
)

func TestQuestion(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitTtestDB(config)
	db.Migrator().DropTable(&entities.User{}, &entities.Question{}, &entities.Option{})
	db.AutoMigrate(&entities.User{}, &entities.Question{}, &entities.Option{})

	var dummyUser entities.User
	dummyUser.Account = "TestAccount1"
	dummyUser.Pwd = "TestPwd1"
	
	userRepo := userRepo.NewUsersRepo(db)
	_, err := userRepo.Create(dummyUser)
	if err != nil {
		fmt.Println(err)
	}
	ec := echo.New()
	ec.Validator = httpUtils.NewValidator()

	questionRepo := questionRepo.NewQuestionsRepo(db)
	optionRepo := optionRepo.NewOptionsRepo(db)
	questionCtrl := NewQuestionsControllers(questionRepo,optionRepo)

	questionGroup := ec.Group("/questions")
	questionGroup.POST("/", questionCtrl.CreateQuestion())

	t.Run("질문하기 시도 -> 성공", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"title": "testTitle",
			"category": 1,
			"options": []string{"test", "test1"},
		})

		uri := "/questions/"
		req := httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()

		ec.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})
}