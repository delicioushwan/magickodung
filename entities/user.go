package entities

import "time"

type User struct {
	//USERID AUTO GENERATE
	UserId  uint `gorm:"primaryKey; autoIncrement"`
	Account		string `gorm:"type:varchar(50); uniqueIndex; not null"`
	Pwd		string
	Last_login time.Time `gorm:"autoCreateTime; not null"`
	Common
}
