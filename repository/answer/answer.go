package answer

import (
	"github.com/delicioushwan/magickodung/entities"

	"gorm.io/gorm"
)

type AnswerRepository struct {
	db *gorm.DB
}

func NewAnswersRepo(db *gorm.DB) *AnswerRepository {
	return &AnswerRepository{db: db}
}

func (op *AnswerRepository) Create(newAnswer entities.Answer) (int64, error) {
	res := op.db.Create(&newAnswer)
	return res.RowsAffected, res.Error
}
