package form_answer

import (
	"amazing_review/internal/adapter/handler/dto/form_question"
	"fmt"
	domain "github.com/MathieuRocher/amazing_domain"
)

type FormAnswerInput struct {
	QuestionID uint   `form:"question_id" validate:"required"`
	Answer     string `form:"answer,omitempty"`
	Rating     uint   `form:"rating,omitempty"`
	Option     uint   `form:"option,omitempty"`
}

func ToDomainFormAnswer(input FormAnswerInput, question form_question.FormQuestionInput) (*domain.FormAnswer, error) {
	// Validation
	switch question.Type {
	case "Field":
		if question.IsRequired && input.Answer == "" {
			return nil, fmt.Errorf("text answer required")
		}
	case "Rating":
		if question.IsRequired && input.Rating == 0 {
			return nil, fmt.Errorf("rating required")
		}
	case "Select", "Radio":
		if question.IsRequired && input.Option == 0 {
			return nil, fmt.Errorf("option selection required")
		}
	}

	// Construction du domain
	return &domain.FormAnswer{
		FormQuestionID: question.ID,
		Answer:         input.Answer,
		Rating:         input.Rating,
		Option:         input.Option,
	}, nil
}
