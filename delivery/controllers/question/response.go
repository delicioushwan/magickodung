package question

import "github.com/delicioushwan/magickodung/entities"

type UserResponse struct {
	User struct {
		Account    string `json:"account"`
		Token    string `json:"token"`
	} `json:"user"`
}

func ToUserResponse(u *entities.User, token string) *UserResponse {
	user := new(UserResponse)
	user.User.Account = u.Account
	user.User.Token = token
	return user
}