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

func (op *OptionRepository) Create(newOption []entities.Option) (int64, error) {
	res := op.db.Create(&newOption)
	return res.RowsAffected, res.Error
}
