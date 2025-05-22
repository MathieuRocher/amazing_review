package review

import (
	"amazing_review/internal/adapter/handler/dto/form_answer"
)

type ReviewInput struct {
	FormId             uint                          `json:"form_id"`
	CourseAssignmentId uint                          `json:"course_assignment_id"`
	Answers            []form_answer.FormAnswerInput `json:"answers"`
}
