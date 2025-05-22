package form_answer

type FormAnswerOutput struct {
	QuestionID uint   `json:"question"`
	Answer     string `json:"answer,omitempty"`
	Rating     uint   `json:"rating,omitempty"`
	Option     uint   `json:"option,omitempty"`
}
