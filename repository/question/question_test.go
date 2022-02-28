package question

import (
	"testing"

	"github.com/delicioushwan/magickodung/configs"
	"github.com/delicioushwan/magickodung/entities"
	"github.com/delicioushwan/magickodung/utils"
	"github.com/stretchr/testify/assert"
)

func create(x uint64) *uint64 {
	return &x
}

func TestQuestionRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitTtestDB(config)

	db.Migrator().DropTable(&entities.Question{})
	db.AutoMigrate(&entities.Question{})



	questionRepo := NewQuestionsRepo(db)
	var mockQuestion entities.Question
	mockQuestion.Title = "test"
	mockQuestion.CategoryId = create(1)

	t.Run("Insert Question into Database", func(t *testing.T){
		res, err := questionRepo.Create(mockQuestion)
		assert.Nil(t, err)
		assert.Equal(t, mockQuestion.Title, res.Title)
		assert.Equal(t, mockQuestion.CategoryId, res.CategoryId)
		assert.Equal(t, 1, int(res.QuestionId))
	})
}