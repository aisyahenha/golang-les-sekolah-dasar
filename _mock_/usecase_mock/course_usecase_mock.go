package usecasemock

import (
	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/stretchr/testify/mock"
)

type CourseUseCaseMock struct {
	mock.Mock
}
// DeleteCourse implements CourseUseCase.
func (u *CourseUseCaseMock) DeleteCourse(id string) error {
	return u.Called(id).Error(0)
}

// FindAllCourse implements CourseUseCaseMock.
func (u *CourseUseCaseMock) FindAllCourse() ([]model.CourseModel, error) {
	args := u.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.CourseModel), nil
}

// FindByCourseId implements CourseUseCaseMock.
func (u *CourseUseCaseMock) FindByCourseId(id string) (model.CourseModel, error) {
	args := u.Called(id)
	if args.Get(1) != nil {
		return model.CourseModel{}, args.Error(1)
	}
	return args.Get(0).(model.CourseModel), nil
}
// RegisterNewCourse implements CourseUseCaseMock.
func (u *CourseUseCaseMock) RegisterNewCourse(payload *model.CourseModel) error{

	return u.Called(payload).Error(0)
}

// UpdateCourse implements CourseUseCaseMock.
func (u *CourseUseCaseMock) UpdateCourse(payload *model.CourseModel) error {
	return u.Called(payload).Error(0)
}
