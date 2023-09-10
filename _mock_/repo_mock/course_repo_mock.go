package repomock

import (
	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/stretchr/testify/mock"

)

type CourseRepoMock struct {
	mock.Mock
}

// Create implements CourseRepoMock.
func (u *CourseRepoMock) Create(payload *model.CourseModel) error {
	return u.Called(payload).Error(0)
}

// Delete implements CourseRepoMock.
func (u *CourseRepoMock) Delete(id string) error {
	return u.Called(id).Error(0)
}

// Get implements CourseRepoMock.
func (u *CourseRepoMock) Get(id string) (model.CourseModel, error) {

	args := u.Called(id)
	if args.Get(1) != nil {
		return model.CourseModel{}, args.Error(1)
	}
	// var userR model.CourseModel = mappingutil.MappingUser(args.Get(0).(model.User))
	return args.Get(0).(model.CourseModel), nil
}

// List implements CourseRepoMock.
func (u *CourseRepoMock) List() ([]model.CourseModel, error) {
	args := u.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.CourseModel) , nil
}

// Update implements CourseRepoMock.
func (u *CourseRepoMock) Update(payload *model.CourseModel) error {
	return u.Called(payload).Error(0)
}

