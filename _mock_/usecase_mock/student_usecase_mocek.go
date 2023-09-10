package usecasemock

import (
	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/stretchr/testify/mock"
)

type StudentUseCaseMock struct {
	mock.Mock
}
// DeleteStudent implements StudentUseCase.
func (u *StudentUseCaseMock) DeleteStudent(id string) error {
	return u.Called(id).Error(0)
}

// FindAllStudent implements StudentUseCaseMock.
func (u *StudentUseCaseMock) FindAllStudent() ([]model.StudentModel, error) {
	args := u.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.StudentModel), nil
}

// FindByStudentId implements StudentUseCaseMock.
func (u *StudentUseCaseMock) FindByStudentId(id string) (model.StudentModel, error) {
	args := u.Called(id)
	if args.Get(1) != nil {
		return model.StudentModel{}, args.Error(1)
	}
	return args.Get(0).(model.StudentModel), nil
}
// RegisterNewStudent implements StudentUseCaseMock.
func (u *StudentUseCaseMock) RegisterNewStudent(payload *model.StudentModel) error {

	return u.Called(payload).Error(0)
}

// UpdateStudent implements StudentUseCaseMock.
func (u *StudentUseCaseMock) UpdateStudent(payload *model.StudentModel) error {
	return u.Called(payload).Error(0)
}
