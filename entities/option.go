package entities

type Option struct {
	OptionId  uint64 `gorm:"primaryKey; autoIncrement"`
	QuetionId		uint64 `gorm:"notNull;"`
	Option		string `gorm:"notNull"`
	Quatity uint64
	State string `gorm:"notNull"`
	Common
}
