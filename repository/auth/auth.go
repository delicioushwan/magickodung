package auth

import (
	"github.com/delicioushwan/magickodung/entities"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) LoginUser(name, password string) (entities.User, error) {
	var user entities.User
	getPass := entities.User{}
	ar.db.Select("password").Where("Name = ?", name).Find(&getPass)
	bcrypt.CompareHashAndPassword([]byte(getPass.Pwd), []byte(password))
	ar.db.Where("Name = ?", name).Find(&user)

	return user, nil
}
