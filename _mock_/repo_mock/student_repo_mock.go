package repomock

import (
	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/stretchr/testify/mock"
	
)

type StudentRepoMock struct {
	mock.Mock
}

// Create implements StudentRepoMock.
func (u *StudentRepoMock) Create(payload *model.StudentModel) error {
	return u.Called(payload).Error(0)
}

// Delete implements StudentRepoMock.
func (u *StudentRepoMock) Delete(id string) error {
	return u.Called(id).Error(0)
}

// Get implements StudentRepoMock.
func (u *StudentRepoMock) Get(id string) (model.StudentModel, error) {

	args := u.Called(id)
	if args.Get(1) != nil {
		return model.StudentModel{}, args.Error(1)
	}
	// var userR model.StudentModel = mappingutil.MappingUser(args.Get(0).(model.User))
	return args.Get(0).(model.StudentModel), nil
}

// List implements StudentRepoMock.
func (u *StudentRepoMock) List() ([]model.StudentModel, error) {
	args := u.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.StudentModel) , nil
}

// Update implements StudentRepoMock.
func (u *StudentRepoMock) Update(payload *model.StudentModel) error {
	return u.Called(payload).Error(0)
}

