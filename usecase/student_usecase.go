package usecase

import (
	// "errors"

	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/aisyahenha/golang-les-sekolah-dasar/repository"
)

type StudentUseCase interface {
	RegisterNewStudent(payload *model.StudentModel) error
	FindAllStudent() ([]model.StudentModel, error)
	FindByStudentId(id string) (model.StudentModel, error)
	UpdateStudent(payload *model.StudentModel) error
	DeleteStudent(id string) error
}

type studentUseCase struct {
	repo repository.StudentRepository
}

// DeleteStudent implements StudentUseCase.
func (u *studentUseCase) DeleteStudent(id string) error {
	return u.repo.Delete(id)
}

// FindAllStudent implements StudentUseCase.
func (u *studentUseCase) FindAllStudent() ([]model.StudentModel, error) {
	return u.repo.List()
}

// FindByStudentId implements StudentUseCase.
func (u *studentUseCase) FindByStudentId(id string) (model.StudentModel, error) {
	return u.repo.Get(id)
}

// RegisterNewStudent implements StudentUseCase.
func (u *studentUseCase) RegisterNewStudent(payload *model.StudentModel) error {

	return u.repo.Create(payload)
}

// UpdateStudent implements StudentUseCase.
func (u *studentUseCase) UpdateStudent(payload *model.StudentModel) error {
	return u.repo.Update(payload)
}

func NewStudentUseCase(repo repository.StudentRepository) StudentUseCase {
	return &studentUseCase{repo: repo}
}
