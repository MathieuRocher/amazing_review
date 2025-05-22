package repository

import (
	"amazing_review/internal/infrastructure/database"

	domain "github.com/MathieuRocher/amazing_domain"
	"gorm.io/gorm"
)

type FormAnswer struct {
	ID             uint `gorm:"primaryKey"`
	ReviewID       uint
	FormQuestionID uint
	Option         uint
	Answer         string
	Rating         uint
}

type FormAnswerRepository struct {
	db *gorm.DB
}

func NewFormAnswerRepository() *FormAnswerRepository {
	return &FormAnswerRepository{database.DB}
}

func (r *FormAnswerRepository) FindAll() ([]domain.FormAnswer, error) {
	var repoFormAnswers []FormAnswer
	if err := r.db.Find(&repoFormAnswers).Error; err != nil {
		return nil, err
	}

	var domainFormAnswers []domain.FormAnswer
	for _, repoFormAnswer := range repoFormAnswers {
		domainFormAnswers = append(domainFormAnswers, *repoFormAnswer.ToDomain())
	}

	return domainFormAnswers, nil
}

func (r *FormAnswerRepository) FindAllWithPagination(page int, limit int) ([]domain.FormAnswer, error) {
	var repoFormAnswers []FormAnswer

	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	if err := r.db.
		Preload("FormAnswers").
		Limit(limit).
		Offset(offset).
		Find(&repoFormAnswers).Error; err != nil {
		return nil, err
	}

	var domainFormAnswers []domain.FormAnswer
	for _, repoFormAnswer := range repoFormAnswers {
		domainFormAnswers = append(domainFormAnswers, *repoFormAnswer.ToDomain())
	}

	return domainFormAnswers, nil
}

func (r *FormAnswerRepository) FindByID(id uint) (*domain.FormAnswer, error) {
	var obj FormAnswer
	if err := r.db.First(&obj, id).Error; err != nil {
		return nil, err
	}
	return obj.ToDomain(), nil
}

func (r *FormAnswerRepository) Create(obj *domain.FormAnswer) error {
	return r.db.Create(FormAnswerFromDomain(obj)).Error
}

func (r *FormAnswerRepository) Update(obj *domain.FormAnswer) error {
	return r.db.Save(FormAnswerFromDomain(obj)).Error
}

func (r *FormAnswerRepository) Delete(id uint) error {
	return r.db.Delete(&FormAnswer{}, id).Error
}

// ToDomain converts a repository FormAnswer to a domain FormAnswer
func (fa *FormAnswer) ToDomain() *domain.FormAnswer {
	return &domain.FormAnswer{
		ID:             fa.ID,
		ReviewID:       fa.ReviewID,
		FormQuestionID: fa.FormQuestionID,
		Option:         fa.Option,
		Answer:         fa.Answer,
		Rating:         fa.Rating,
	}
}

// FormAnswerFromDomain converts a domain FormAnswer to a repository FormAnswer
func FormAnswerFromDomain(fa *domain.FormAnswer) *FormAnswer {
	return &FormAnswer{
		ID:             fa.ID,
		ReviewID:       fa.ReviewID,
		FormQuestionID: fa.FormQuestionID,
		Option:         fa.Option,
		Answer:         fa.Answer,
		Rating:         fa.Rating,
	}
}

func FormAnswersFromDomain(assignments []domain.FormAnswer) []FormAnswer {
	var repoFormAnswers []FormAnswer
	for _, a := range assignments {
		repoFormAnswers = append(repoFormAnswers, *FormAnswerFromDomain(&a))
	}
	return repoFormAnswers
}

func ToDomainFormAnswers(repoAnswers []FormAnswer) []domain.FormAnswer {
	var domainAnswers []domain.FormAnswer
	for _, ans := range repoAnswers {
		domainAnswers = append(domainAnswers, *ans.ToDomain())
	}
	return domainAnswers
}
