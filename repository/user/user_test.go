package user

import (
	"testing"

	"github.com/delicioushwan/magickodung/configs"
	"github.com/delicioushwan/magickodung/entities"
	"github.com/delicioushwan/magickodung/utils"

	"github.com/stretchr/testify/assert"
)

func TestUsersRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	userRepo := NewUsersRepo(db)
	var mockInserUser entities.User
	mockInserUser.Account = "TestName1"
	mockInserUser.Pwd = "TestPassword1"

	t.Run("Insert User into Database", func(t *testing.T) {

		res, err := userRepo.Create(mockInserUser)
		assert.Nil(t, err)
		assert.Equal(t, mockInserUser.Account, res.Account)
		assert.Equal(t, 1, int(res.UserId))
	})

	t.Run("Select User from Database", func(t *testing.T) {
		res, err := userRepo.Get(1)
		assert.Nil(t, err)
		//원하는 값을 가지고 있는지 체크하는 방법 찾아서 개선 필요
		assert.Equal(t,res.Account, mockInserUser.Account)
		assert.Equal(t,res.Pwd, mockInserUser.Pwd)
	})

	t.Run("Select User By account from Database", func(t *testing.T) {
		res, err := userRepo.GetByAccount(mockInserUser.Account)
		assert.Nil(t, err)
		//원하는 값을 가지고 있는지 체크하는 방법 찾아서 개선 필요
		assert.Equal(t,res.Account, mockInserUser.Account)
		assert.Equal(t,res.Pwd, mockInserUser.Pwd)
	})
}
