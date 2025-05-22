package review

import (
	"amazing_review/internal/adapter/handler/dto/form_answer"
	domain "github.com/MathieuRocher/amazing_domain"
)

type ReviewOutput struct {
	ID                 uint                           `json:"id"`
	UserID             uint                           `json:"user"`
	FormID             uint                           `json:"form"`
	CourseAssignmentID uint                           `json:"course_assignment"`
	Answers            []form_answer.FormAnswerOutput `json:"answers"`
}

func FromDomain(r *domain.Review) *ReviewOutput {
	return &ReviewOutput{
		ID:                 r.ID,
		UserID:             r.UserID,
		FormID:             r.FormID,
		CourseAssignmentID: r.CourseAssignmentID,
		Answers:            mapAnswers(r.FormAnswers),
	}
}

func mapAnswers(domainAnswers []domain.FormAnswer) []form_answer.FormAnswerOutput {
	var out []form_answer.FormAnswerOutput
	for _, a := range domainAnswers {
		out = append(out, form_answer.FormAnswerOutput{
			QuestionID: a.FormQuestionID,
			Answer:     a.Answer,
			Rating:     a.Rating,
			Option:     a.Option,
		})
	}
	return out
}
