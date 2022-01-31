package user

import "github.com/delicioushwan/magickodung/entities"

type RegisterUserResponseFormat struct {
	Message string          `json:"message"`
	Token    string `json:"token"`
}

type LoginUserResponseFormat struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type GetUsersResponseFormat struct {
	Message string          `json:"message"`
	Data    []entities.User `json:"data"`
}

type GetUserResponseFormat struct {
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
}

type PutUserResponseFormat struct {
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
}

type DeleteUserResponseFormat struct {
	Message string `json:"message"`
}
