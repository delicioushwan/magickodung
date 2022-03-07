package entities

type Question struct {
	QuestionId  uint64 `gorm:"primaryKey; autoIncrement"`
	Title		string `gorm:"type:varchar(50); notNull;"`
	CategoryId		*uint64 `gorm:"notNull"`
	UserId uint64 `gorm:"notNull"`
	State string `gorm:"notNull"`
	Common
}

type GetCommonQuestionsResponse struct {
	QuestionId  uint64 
	Title		string
	OptionId  uint64
	Option		string
	Quantity uint64
}