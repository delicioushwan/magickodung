package entities

type Option struct {
	OptionId  uint64 `gorm:"primaryKey; autoIncrement"`
	QuestionId		uint64 `gorm:"notNull;"`
	Option		string `gorm:"notNull"`
	Quantity uint64 `gorm:"notNull"`
	State string `gorm:"notNull"`
	Common
}
