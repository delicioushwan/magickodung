package question

type CreateQuestionRequestFormat struct {
	Title     string `json:"title" form:"title" validate:"required"`
	CategoryId uint64 `json:"category" form:"category" validate:"required"`
	Options []string `json:"options[]" form:"options[]" validate:"required"`
}


