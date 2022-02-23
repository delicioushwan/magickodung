package option

import "github.com/delicioushwan/magickodung/entities"

type OptionInterface interface {
	Create(newOption []entities.Option) (int64, error)
}