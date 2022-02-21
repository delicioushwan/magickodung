package question

import "github.com/delicioushwan/magickodung/entities"

type QuestionInterface interface {
	Create(newQuestion entities.Question) (entities.Question, error)
}
