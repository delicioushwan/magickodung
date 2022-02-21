package option

import (
	"github.com/delicioushwan/magickodung/entities"

	"gorm.io/gorm"
)

type OptionRepository struct {
	db *gorm.DB
}

func NewOptionsRepo(db *gorm.DB) *OptionRepository {
	return &OptionRepository{db: db}
}

func (op *OptionRepository) Create(newOption entities.Option) (entities.Option, error) {
	err := op.db.Create(&newOption).Error
	return newOption, err
}
