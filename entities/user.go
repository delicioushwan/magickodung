package entities

type User struct {
	Common
	//USERID AUTO GENERATE
	UserId  uint `gorm:"primaryKey"`
	Account		string
	Pwd		string
}
