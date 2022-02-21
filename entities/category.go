package entities

type Category struct {
	CategoryId  uint64 `gorm:"primaryKey; autoIncrement"`
	Category		string `gorm:"notNull;"`
}
