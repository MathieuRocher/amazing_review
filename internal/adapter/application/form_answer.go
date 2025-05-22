package application

import (
	domain "github.com/MathieuRocher/amazing_domain"
)

type FormAnswerUseCaseInterface interface {
	FindAll() ([]domain.FormAnswer, error)
	FindAllWithPagination(p int, l int) ([]domain.FormAnswer, error)
	FindByID(id uint) (*domain.FormAnswer, error)
	Create(q *domain.FormAnswer) error
	Update(q *domain.FormAnswer) error
	Delete(id uint) error
}

type FormAnswerRepositoryInterface interface {
	FindAll() ([]domain.FormAnswer, error)
	FindAllWithPagination(p int, l int) ([]domain.FormAnswer, error)
	FindByID(id uint) (*domain.FormAnswer, error)
	Create(formAnswer *domain.FormAnswer) error
	Update(formAnswer *domain.FormAnswer) error
	Delete(id uint) error
}

type FormAnswerUseCase struct {
	repo FormAnswerRepositoryInterface
}

func NewFormAnswerUseCase(r FormAnswerRepositoryInterface) FormAnswerUseCaseInterface {
	return &FormAnswerUseCase{repo: r}
}

func (uc *FormAnswerUseCase) FindAll() ([]domain.FormAnswer, error) {
	return uc.repo.FindAll()
}

func (uc *FormAnswerUseCase) FindAllWithPagination(page int, limit int) ([]domain.FormAnswer, error) {
	return uc.repo.FindAllWithPagination(page, limit)
}

func (uc *FormAnswerUseCase) FindByID(id uint) (*domain.FormAnswer, error) {
	return uc.repo.FindByID(id)
}

func (uc *FormAnswerUseCase) Create(q *domain.FormAnswer) error {
	return uc.repo.Create(q)
}

func (uc *FormAnswerUseCase) Update(q *domain.FormAnswer) error {
	return uc.repo.Update(q)
}

func (uc *FormAnswerUseCase) Delete(id uint) error {
	return uc.repo.Delete(id)
}
