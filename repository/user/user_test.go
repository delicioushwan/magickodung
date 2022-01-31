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

	t.Run("Insert User into Database", func(t *testing.T) {
		var mockInserUser entities.User
		mockInserUser.Account = "TestName1"
		mockInserUser.Pwd = "TestPassword1"

		res, err := userRepo.Create(mockInserUser)
		assert.Nil(t, err)
		assert.Equal(t, mockInserUser.Account, res.Account)
		assert.Equal(t, 1, int(res.UserId))
	})

	t.Run("Select User from Database", func(t *testing.T) {
		res, err := userRepo.Get(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

}
