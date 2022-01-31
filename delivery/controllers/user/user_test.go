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
	"github.com/tidwall/gjson"

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

	userCtrl := NewUsersControllers(userRepo)

	usersGroup := ec.Group("/users")
	usersGroup.POST("/signup", userCtrl.PostUserCtrl())

	t.Run("회원가입 시도 -> 성공", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"account": "TestAccount2",
			"pwd": "TestPwd1",
		})
		uri := "/users/signup"
		req := httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		ec.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	cases := []struct {
		name     string
		account    string
		password string
		// expected
		code int
		msg  string
	}{
		{
			name:     "회원가입 시도 -> 실패(missing account)",
			password: "pass",
			code:     http.StatusBadRequest,
			msg:      "Field validation for 'Account' failed on the 'required' tag",
		}, {
			name:     "회원가입 시도 -> 실패(missing password)",
			account:  "user@gmail.com",
			code:     http.StatusBadRequest,
			msg:      "Field validation for 'Pwd' failed on the 'required' tag",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			uri := "/users/signup"
			res := httptest.NewRecorder()
			reqBody, _ := json.Marshal(map[string]string{
				"account": tc.account,
				"pwd": tc.password,
			})
	
			req, _ := http.NewRequest(http.MethodPost, uri,  bytes.NewBuffer(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			ec.ServeHTTP(res, req)

			assert.Equal(t, tc.code, res.Code)
			assert.Contains(t, gjson.Get(res.Body.String(), "errors.body").String(), tc.msg)
		})
	}

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
