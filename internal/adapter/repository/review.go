package repository

import (
	"amazing_review/internal/infrastructure/database"

	domain "github.com/MathieuRocher/amazing_domain"
	"gorm.io/gorm"
)

type Review struct {
	ID                 uint `gorm:"primaryKey"`
	FormID             uint
	UserID             uint
	CourseAssignmentID uint

	FormAnswers []FormAnswer `gorm:"foreignKey:ReviewID"`
}

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository() *ReviewRepository {
	return &ReviewRepository{database.DB}
}

func (r *ReviewRepository) FindAll() ([]domain.Review, error) {
	var repoReviews []Review
	if err := r.db.Preload("FormAnswers").Find(&repoReviews).Error; err != nil {
		return nil, err
	}

	var domainReviews []domain.Review
	for _, repoReview := range repoReviews {
		domainReviews = append(domainReviews, *repoReview.ToDomain())
	}

	return domainReviews, nil
}

func (r *ReviewRepository) FindAllWithPagination(page int, limit int) ([]domain.Review, error) {
	var repoReviews []Review

	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	if err := r.db.
		Preload("FormAnswers").
		Limit(limit).
		Offset(offset).
		Find(&repoReviews).Error; err != nil {
		return nil, err
	}

	var domainReviews []domain.Review
	for _, repoReview := range repoReviews {
		domainReviews = append(domainReviews, *repoReview.ToDomain())
	}

	return domainReviews, nil
}

func (r *ReviewRepository) FindByID(id uint) (*domain.Review, error) {
	var repoReview Review
	if err := r.db.
		Preload("FormAnswers").
		First(&repoReview, id).Error; err != nil {
		return nil, err
	}
	return repoReview.ToDomain(), nil
}

func (r *ReviewRepository) Create(obj *domain.Review) error {
	return r.db.Create(ReviewFromDomain(obj)).Error
}

func (r *ReviewRepository) Update(obj *domain.Review) error {
	return r.db.Save(ReviewFromDomain(obj)).Error
}

func (r *ReviewRepository) Delete(id uint) error {
	return r.db.Delete(&Review{}, id).Error
}

// ToDomain converts a repository Review to a domain Review
func (r *Review) ToDomain() *domain.Review {
	return &domain.Review{
		ID:                 r.ID,
		FormID:             r.FormID,
		UserID:             r.UserID,
		CourseAssignmentID: r.CourseAssignmentID,
		FormAnswers:        ToDomainFormAnswers(r.FormAnswers),
	}
}

// ReviewFromDomain converts a domain Review to a repository Review
func ReviewFromDomain(r *domain.Review) *Review {
	return &Review{
		ID:                 r.ID,
		FormID:             r.FormID,
		UserID:             r.UserID,
		CourseAssignmentID: r.CourseAssignmentID,
		FormAnswers:        FormAnswersFromDomain(r.FormAnswers),
	}
}
