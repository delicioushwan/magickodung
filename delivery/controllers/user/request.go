package user

type UserCommonRequestFormat struct {
	Account     string `json:"account" form:"account" validate:"required"`
	Pwd string `json:"pwd" form:"pwd" validate:"required"`
}

