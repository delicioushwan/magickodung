package question

import (
	"github.com/delicioushwan/magickodung/entities"

	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionsRepo(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (qt *QuestionRepository) Create(newQuestion entities.Question) (entities.Question, error) {
	err := qt.db.Create(&newQuestion).Error
	return newQuestion, err
}
