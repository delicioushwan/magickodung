package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/delicioushwan/magickodung/configs"
	"github.com/delicioushwan/magickodung/entities"
	userRepo "github.com/delicioushwan/magickodung/repository/user"
	"github.com/delicioushwan/magickodung/utils"
	"github.com/delicioushwan/magickodung/utils/httpUtils"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {

	config := configs.GetConfig()
	db := utils.InitTtestDB(config)
	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

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

	t.Run("회원가입 시도 -> 성공", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"account": "TestAccount2",
			"pwd": "TestPwd1",
		})

		req := httptest.NewRequest(http.MethodPost, "/users/signup", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		context := ec.NewContext(req, res)

		userCon := NewUsersControllers(userRepo)

		if assert.NoError(t, userCon.PostUserCtrl()(context)) {
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})

	// t.Run("회원가입 시도 -> 파라미터 없어서 실패", func(t *testing.T){
	// 	reqBody, _ := json.Marshal(map[string]string{
	// 		"name":     "TestAccount1",
	// 		"password": "TestPwd1",
	// 	})


	// })
// 	t.Run("POST /users/login", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"name":     "TestAccount1",
// 			"password": "TestPwd1",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		context := ec.NewContext(req, res)
// 		context.SetPath("/users/login")

// 		authCon := auth.NewAuthControllers(mockAuthRepository{})
// 		authCon.LoginAuthCtrl()(context)

// 		responses := LoginUserResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, "Successful Operation", responses.Message)
// 		assert.Equal(t, 200, res.Code)

// 	})

}
