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

	//라우터도 각 도메인별로 쪼개서 해당 모듈불러오면 되도록 수정 필요
	usersGroup := ec.Group("/users")
	usersGroup.POST("/signup", userCtrl.Signup())
	usersGroup.POST("/login", userCtrl.Login())

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

	signUpcases := []struct {
		name     string
		account    string
		password string
		// expected
		code int
		msg  string
	}{
		{
			name:     "회원가입 시도 -> 실패(missing password)",
			account:  "user@gmail.com",
			password: "",
			code:     http.StatusBadRequest,
			msg:      "Field validation for 'Pwd' failed on the 'required' tag",
		}, {
			name:     "회원가입 시도 -> 실패(missing account)",
			account:  "",
			password: "123",
			code:     http.StatusBadRequest,
			msg:      "Field validation for 'Account' failed on the 'required' tag",
		},
	}

	for _, tc := range signUpcases {
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

	t.Run("로그인 시도 -> 성공", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"account": "TestAccount2",
			"pwd": "TestPwd1",
		})
		uri := "/users/login"
		req := httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		ec.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	loginCases := []struct {
		name     string
		account    string
		password string
		// expected
		code int
		msg  string
	}{
		{
			name:     "로그인 시도 -> 실패(missing password)",
			account:  "user@gmail.com",
			password: "",
			code:     http.StatusBadRequest,
			msg:      "Field validation for 'Pwd' failed on the 'required' tag",
		}, {
			name:     "로그인 시도 -> 실패(missing account)",
			account:  "",
			password: "123",
			code:     http.StatusBadRequest,
			msg:      "Field validation for 'Account' failed on the 'required' tag",
		}, {
			name:     "로그인 시도 -> 실패(unmatch password)",
			account:  "TestAccount2",
			password: "123",
			code:     http.StatusBadRequest,
			msg:      "아이디와 비밀번호를 확인해 주세요.",
		}, {
			name:     "로그인 시도 -> 실패(account does not exist)",
			account:  "test",
			password: "123",
			code:     http.StatusBadRequest,
			msg:      "존재하지 않는 회원입니다. \n 아이디와 비밀번호를 확인해 주세요.",
		},
		
	}

	for _, tc := range loginCases {
		t.Run(tc.name, func(t *testing.T) {
			uri := "/users/login"
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

}
