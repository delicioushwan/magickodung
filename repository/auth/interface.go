package auth

import "github.com/delicioushwan/magickodung/entities"

type AuthInterface interface {
	Login(account, pwd string) (entities.User, error)
}
