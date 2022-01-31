package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/delicioushwan/magickodung/configs"
	"github.com/delicioushwan/magickodung/delivery/controllers/auth"
	"github.com/delicioushwan/magickodung/entities"
	"github.com/delicioushwan/magickodung/repository/user"
	"github.com/delicioushwan/magickodung/utils"

	"github.com/labstack/echo/v4"
	"github.com/magiconair/properties/assert"
)

func TestUsers(t *testing.T) {

	config := configs.GetConfig()
	db := utils.InitTtestDB(config)
	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	var dummyUser entities.User
	dummyUser.Account = "TestAccount1"
	dummyUser.Pwd = "TestPwd1"

	useRep := user.NewUsersRepo(db)
	_, err := useRep.Create(dummyUser)
	if err != nil {
		fmt.Println(err)
	}
	ec := echo.New()

	t.Run("POST /users/signup", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "TestAccount1",
			"password": "TestPwd1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := ec.NewContext(req, res)
		context.SetPath("/users/signup")

		userCon := NewUsersControllers(mockUserRepository{})
		userCon.PostUserCtrl()(context)

		responses := RegisterUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, res.Code)

	})
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

func TestFalseUsers(t *testing.T) {
	e := echo.New()

	t.Run("POST /users/signup", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/signup")

		userCon := NewUsersControllers(mockFalseUserRepository{})
		userCon.PostUserCtrl()(context)

		responses := RegisterUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, responses.Message, "Internal Server Error")
		assert.Equal(t, res.Code, 500)
	})
	t.Run("POST /users/signup", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]int{
			"name": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/signup")

		userCon := NewUsersControllers(mockFalseUserRepository{})
		userCon.PostUserCtrl()(context)

		responses := RegisterUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, responses.Message, "Bad Request")
		assert.Equal(t, res.Code, 400)
	})
	t.Run("POST /users/login", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		authCon := auth.NewAuthControllers(mockFalseAuthRepository{})
		authCon.LoginAuthCtrl()(context)

		responses := LoginUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, responses.Message, "Internal Server Error")
		assert.Equal(t, res.Code, 500)
	})

}

type mockAuthRepository struct{}

func (mua mockAuthRepository) Login(name, password string) (entities.User, error) {
	return entities.User{UserId: 1, Account: "TestAccount1", Pwd: "TestPwd1"}, nil
}

type mockUserRepository struct{}

func (mur mockUserRepository) GetAll() ([]entities.User, error) {
	return []entities.User{
		{Account: "TestAccount1", Pwd: "TestPwd1"},
	}, nil
}
func (mur mockUserRepository) Get(userId int) (entities.User, error) {
	return entities.User{Account: "TestAccount1", Pwd: "TestPwd1"}, nil
}
func (mur mockUserRepository) Create(newUser entities.User) (entities.User, error) {
	return entities.User{Account: "TestAccount1", Pwd: "TestPwd1"}, nil
}
func (mur mockUserRepository) Update(updateUser entities.User, userId int) (entities.User, error) {
	return entities.User{Account: "TestAccount1", Pwd: "TestPwd1"}, nil
}
func (mur mockUserRepository) Delete(userId int) (entities.User, error) {
	return entities.User{UserId: 1, Account: "TestAccount1", Pwd: "TestPwd1"}, nil
}

// FALSE SECTION
type mockFalseAuthRepository struct{}

func (mua mockFalseAuthRepository) Login(account, pwd string) (entities.User, error) {
	return entities.User{UserId: 1, Account: "TestAccount1", Pwd: "TestPwd1"}, errors.New("Bad Request")
}

type mockFalseUserRepository struct{}

func (mur mockFalseUserRepository) GetAll() ([]entities.User, error) {
	return []entities.User{
		{Account: "TestAccount1", Pwd: "TestPwd1"},
	}, errors.New("Bad Request")
}
func (mur mockFalseUserRepository) Get(userId int) (entities.User, error) {
	return entities.User{Account: "TestAccount1", Pwd: "TestPwd1"}, errors.New("Bad Request")
}
func (mur mockFalseUserRepository) Create(newUser entities.User) (entities.User, error) {
	return entities.User{Account: "TestAccount1", Pwd: "TestPwd1"}, errors.New("Bad Request")
}
func (mur mockFalseUserRepository) Update(updateUser entities.User, userId int) (entities.User, error) {
	return entities.User{Account: "TestAccount1", Pwd: "TestPwd1"}, errors.New("Bad Request")
}
func (mur mockFalseUserRepository) Delete(userId int) (entities.User, error) {
	return entities.User{Account: "TestAccount1", Pwd: "TestPwd1"}, errors.New("Bad Request")
}
