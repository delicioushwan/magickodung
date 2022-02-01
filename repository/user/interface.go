package user

import "github.com/delicioushwan/magickodung/entities"

type UserInterface interface {
	Get(userId int) (entities.User, error)
	Create(newUser entities.User) (entities.User, error)
	GetByAccount(account string) (entities.User, error)
}
