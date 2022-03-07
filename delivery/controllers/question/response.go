package question

type Option struct {
	OptionId  uint64 `json:"optionid"`
	Option		string `json:"option"`
	Quantity uint64 `json:"quantity"`
}
type CommonQuestionsResponse struct {
	QuestionId  uint64  `json:"questionId"`
	Title		string		`json:"title"`
	Options [] Option `json:"options"`
}
