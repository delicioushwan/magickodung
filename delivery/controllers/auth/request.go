package auth

type LoginRequest struct {
	Account     string `json:"account" form:"account"`
	Pwd string `json:"pwd" form:"pwd"`
}
