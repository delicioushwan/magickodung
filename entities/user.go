package entities

import "time"

type User struct {
	//USERID AUTO GENERATE
	UserId  uint `gorm:"primaryKey; autoIncrement"`
	Account		string `gorm:"type:varchar(50); unique; notNull;"`
	Pwd		string `gorm:"notNull"`
	Last_login time.Time `gorm:"autoCreateTime"`
	Common
}
