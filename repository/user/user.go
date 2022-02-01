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
	err := ur.db.Find(&user, userId).Error
	return user, err
}

func (ur *UserRepository) GetByAccount(account string) (entities.User, error) {
	user := entities.User{}
	err := ur.db.Where("account = ?", account).First(&user).Error
	return user, err
}

func (ur *UserRepository) Create(newUser entities.User) (entities.User, error) {
	err := ur.db.Create(&newUser).Error
	return newUser, err
}
