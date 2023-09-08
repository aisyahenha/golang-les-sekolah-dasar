package usecase

import (
	// "errors"

	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/aisyahenha/golang-les-sekolah-dasar/repository"
	// "github.com/aisyahenha/golang-les-sekolah-dasar/repository"
)

type CourseUseCase interface {
	RegisterNewCourse(payload *model.CourseModel) error
	FindAllCourse() ([]model.CourseModel, error)
	FindByCourseId(id string) (model.CourseModel, error)
	UpdateCourse(payload *model.CourseModel) error
	DeleteCourse(id string) error
}

type courseUseCase struct {
	repo repository.CourseRepository
}

// DeleteCourse implements CourseUseCase.
func (u *courseUseCase) DeleteCourse(id string) error {
	return u.repo.Delete(id)
}

// FindAllCourse implements CourseUseCase.
func (u *courseUseCase) FindAllCourse() ([]model.CourseModel, error) {
	return u.repo.List()
}

// FindByCourseId implements CourseUseCase.
func (u *courseUseCase) FindByCourseId(id string) (model.CourseModel, error) {
	return u.repo.Get(id)
}

// RegisterNewCourse implements CourseUseCase.
func (u *courseUseCase) RegisterNewCourse(payload *model.CourseModel) error {

	return u.repo.Create(payload)
}

// UpdateCourse implements CourseUseCase.
func (u *courseUseCase) UpdateCourse(payload *model.CourseModel) error {
	return u.repo.Update(payload)
}

func NewCourseUseCase(repo repository.CourseRepository) CourseUseCase {
	return &courseUseCase{repo: repo}
}
