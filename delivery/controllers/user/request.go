package user

type UserCommonRequestFormat struct {
	Account     string `json:"accoung" form:"account"`
	Pwd string `json:"pwd" form:"pwd"`
}

