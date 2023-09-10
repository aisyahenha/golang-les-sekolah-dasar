package repomock

import (
	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/stretchr/testify/mock"
	
)

type TeacherRepoMock struct {
	mock.Mock
}

// Create implements TeacherRepoMock.
func (u *TeacherRepoMock) Create(payload *model.TeacherModel) error {
	return u.Called(payload).Error(0)
}

// Delete implements TeacherRepoMock.
func (u *TeacherRepoMock) Delete(id string) error {
	return u.Called(id).Error(0)
}

// Get implements TeacherRepoMock.
func (u *TeacherRepoMock) Get(id string) (model.TeacherModel, error) {

	args := u.Called(id)
	if args.Get(1) != nil {
		return model.TeacherModel{}, args.Error(1)
	}
	// var userR model.TeacherModel = mappingutil.MappingUser(args.Get(0).(model.User))
	return args.Get(0).(model.TeacherModel), nil
}

// List implements TeacherRepoMock.
func (u *TeacherRepoMock) List() ([]model.TeacherModel, error) {
	args := u.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.TeacherModel) , nil
}

// Update implements TeacherRepoMock.
func (u *TeacherRepoMock) Update(payload *model.TeacherModel) error {
	return u.Called(payload).Error(0)
}

