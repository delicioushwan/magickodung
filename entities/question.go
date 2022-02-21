package entities

type Question struct {
	QuestionId  uint64 `gorm:"primaryKey; autoIncrement"`
	Title		string `gorm:"type:varchar(50); notNull;"`
	CategoryId		uint64 `gorm:"notNull"`
	UserId uint64 `gorm:"notNull"`
	State string `gorm:"notNull"`
	Common
}
