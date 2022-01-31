package user

import (
	"github.com/delicioushwan/magickodung/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Get(userId int) (entities.User, error) {
	user := entities.User{}
	ur.db.Find(&user, userId)
	return user, nil
}

func (ur *UserRepository) Create(newUser entities.User) (entities.User, error) {
	ur.db.Save(&newUser)
	return newUser, nil
}
