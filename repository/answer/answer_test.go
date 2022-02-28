package answer

import (
	"testing"

	"github.com/delicioushwan/magickodung/configs"
	"github.com/delicioushwan/magickodung/entities"
	"github.com/delicioushwan/magickodung/utils"
	"github.com/stretchr/testify/assert"
)

func TestAnswerRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitTtestDB(config)

	db.Migrator().DropTable(&entities.Answer{})
	db.AutoMigrate(&entities.Answer{})

	answerRepo := NewAnswersRepo(db)
	var mockAnswer entities.Answer
	mockAnswer.QuestionId = 1
	mockAnswer.OptionId = 1
	mockAnswer.UserId = 1

	t.Run("Insert Answers into Database", func(t *testing.T) {
		rowsAffected, err := answerRepo.Create(mockAnswer)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(rowsAffected))
	})
}