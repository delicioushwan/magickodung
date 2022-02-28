package entities

type Answer struct {
	AnswerId  uint64 `gorm:"primaryKey; autoIncrement"`
	UserId		uint64 `gorm:"notNull;"`
	QuestionId uint64 `gorm:"notNull;"`
	OptionId uint64 `gorm:"notNull;"`
}
