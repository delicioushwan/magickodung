package answer

import "github.com/delicioushwan/magickodung/entities"

type AnswerInterface interface {
	Create(newAnswer entities.Answer) (int64, error)
}