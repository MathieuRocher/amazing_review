package form_question

type FormQuestionInput struct {
	ID         uint   `json:"id"`
	Question   string `json:"question"`
	Type       string `json:"type"` // ou enum
	Options    string `json:"options"`
	IsRequired bool   `json:"is_required"`
}
