package question

import "github.com/delicioushwan/magickodung/entities"

type QuestionInterface interface {
	Create(newQuestion entities.Question) (entities.Question, error)
	GetRandom(userId uint64, categoryId *uint64) ([]entities.GetCommonQuestionsResponse, error)
	GetQuestionsByUserId(userId uint64, offset int, limit int) ([]entities.GetCommonQuestionsResponse, error)
}
