package application

import (
	"amazing_review/internal/adapter/handler/dto/review"
	domain "github.com/MathieuRocher/amazing_domain"
)

type ReviewUseCaseInterface interface {
	FindAll() ([]domain.Review, error)
	FindByID(id uint) (*domain.Review, error)
	Create(input *review.ReviewInput, userID uint) error
	Update(q *domain.Review) error
	Delete(id uint) error
}

type ReviewRepositoryInterface interface {
	FindAll() ([]domain.Review, error)
	FindByID(id uint) (*domain.Review, error)
	Create(review *domain.Review) error
	Update(review *domain.Review) error
	Delete(id uint) error
}

type UserRepositoryInterface interface {
	FindByID(id uint) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}

type FormRepositoryInterface interface {
	FindByID(id uint) (*domain.Form, error)
}

type ReviewUseCase struct {
	reviewRepo     ReviewRepositoryInterface
	formAnswerRepo FormAnswerRepositoryInterface
}

func NewReviewUseCase(
	reviewRepo ReviewRepositoryInterface,
	formAnswerRepo FormAnswerRepositoryInterface,
) *ReviewUseCase {
	return &ReviewUseCase{
		reviewRepo:     reviewRepo,
		formAnswerRepo: formAnswerRepo,
	}
}
func (uc *ReviewUseCase) FindAll() ([]domain.Review, error) {
	return uc.reviewRepo.FindAll()
}

func (uc *ReviewUseCase) FindByID(id uint) (*domain.Review, error) {
	return uc.reviewRepo.FindByID(id)
}

func (uc *ReviewUseCase) Create(input *review.ReviewInput, userID uint) error {
	var answers []domain.FormAnswer
	for _, ans := range input.Answers {
		answers = append(answers, domain.FormAnswer{
			FormQuestionID: ans.QuestionID,
			Answer:         ans.Answer,
			Rating:         ans.Rating,
			Option:         ans.Option,
		})
	}

	domainReview := domain.Review{
		UserID:             userID,
		FormID:             input.FormId,
		CourseAssignmentID: input.CourseAssignmentId,
		FormAnswers:        answers,
	}

	return uc.reviewRepo.Create(&domainReview)
}

func (uc *ReviewUseCase) Update(q *domain.Review) error {
	return uc.reviewRepo.Update(q)
}

func (uc *ReviewUseCase) Delete(id uint) error {
	return uc.reviewRepo.Delete(id)
}
