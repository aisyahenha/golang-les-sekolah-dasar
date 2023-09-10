package usecasemock

import (
	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/stretchr/testify/mock"
)

type TeacherUseCaseMock struct {
	mock.Mock
}
// DeleteTeacher implements TeacherUseCase.
func (u *TeacherUseCaseMock) DeleteTeacher(id string) error {
	return u.Called(id).Error(0)
}

// FindAllTeacher implements TeacherUseCaseMock.
func (u *TeacherUseCaseMock) FindAllTeacher() ([]model.TeacherModel, error) {
	args := u.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.TeacherModel), nil
}

// FindByTeacherId implements TeacherUseCaseMock.
func (u *TeacherUseCaseMock) FindByTeacherId(id string) (model.TeacherModel, error) {
	args := u.Called(id)
	if args.Get(1) != nil {
		return model.TeacherModel{}, args.Error(1)
	}
	return args.Get(0).(model.TeacherModel), nil
}
// RegisterNewTeacher implements TeacherUseCaseMock.
func (u *TeacherUseCaseMock) RegisterNewTeacher(payload *model.TeacherModel) error {
	return u.Called(payload).Error(0)
}


// UpdateTeacher implements TeacherUseCaseMock.
func (u *TeacherUseCaseMock) UpdateTeacher(payload *model.TeacherModel) error {
	return u.Called(payload).Error(0)
}
