package usecase

import (
	// "errors"

	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/aisyahenha/golang-les-sekolah-dasar/repository"
)

type TeacherUseCase interface {
	RegisterNewTeacher(payload *model.TeacherModel) error
	FindAllTeacher() ([]model.TeacherModel, error)
	FindByTeacherId(id string) (model.TeacherModel, error)
	UpdateTeacher(payload *model.TeacherModel) error
	DeleteTeacher(id string) error
}

type teacherUseCase struct {
	repo repository.TeacherRepository
}

// DeleteTeacher implements TeacherUseCase.
func (u *teacherUseCase) DeleteTeacher(id string) error {
	return u.repo.Delete(id)
}

// FindAllTeacher implements TeacherUseCase.
func (u *teacherUseCase) FindAllTeacher() ([]model.TeacherModel, error) {
	return u.repo.List()
}

// FindByTeacherId implements TeacherUseCase.
func (u *teacherUseCase) FindByTeacherId(id string) (model.TeacherModel, error) {
	return u.repo.Get(id)
}

// RegisterNewTeacher implements TeacherUseCase.
func (u *teacherUseCase) RegisterNewTeacher(payload *model.TeacherModel) error {

	return u.repo.Create(payload)
}

// UpdateTeacher implements TeacherUseCase.
func (u *teacherUseCase) UpdateTeacher(payload *model.TeacherModel) error {
	return u.repo.Update(payload)
}

func NewTeacherUseCase(repo repository.TeacherRepository) TeacherUseCase {
	return &teacherUseCase{repo: repo}
}
