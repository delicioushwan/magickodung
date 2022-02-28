package answer

type CreateAnswerRequestFormat struct {
	QuestionId uint64 `json:"questionId" form:"questionId" validate:"required"`
	OptionId uint64 `json:"optionId" form:"optionId" validate:"required"`
}


