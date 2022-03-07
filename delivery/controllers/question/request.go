package question

type CreateQuestionRequestFormat struct {
	Title     string `json:"title" form:"title" validate:"required"`
	CategoryId *uint64 `json:"category" form:"category" validate:"required"` //*uint64 0을 통과시키기 위하여 포인터 사용
	Options []string `json:"options" form:"options" validate:"required"`
}

