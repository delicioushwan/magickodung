package option

import (
	"testing"

	"github.com/delicioushwan/magickodung/configs"
	"github.com/delicioushwan/magickodung/entities"
	"github.com/delicioushwan/magickodung/utils"
	"github.com/stretchr/testify/assert"
)

func TestOptionRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitTtestDB(config)

	db.Migrator().DropTable(&entities.Option{})
	db.AutoMigrate(&entities.Option{})

	optionRepo := NewOptionsRepo(db)
	var mockOptions []entities.Option
	mockOptions = append(mockOptions,
		entities.Option{QuestionId: 1, Option: "test1"},
		entities.Option{QuestionId: 1, Option: "test2"},
	)

	t.Run("Insert Options into Database", func(t *testing.T){
		rowsAffected, err := optionRepo.Create(mockOptions)
		assert.Nil(t, err)
		assert.Equal(t, 2, int(rowsAffected))
	})
}