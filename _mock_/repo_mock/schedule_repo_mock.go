package repomock

import (
	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/stretchr/testify/mock"
	
)

type ScheduleRepoMock struct {
	mock.Mock
}

// Create implements ScheduleRepoMock.
func (u *ScheduleRepoMock) Create(payload *model.ScheduleModel) error {
	return u.Called(payload).Error(0)
}

// Delete implements ScheduleRepoMock.
func (u *ScheduleRepoMock) Delete(id string) error {
	return u.Called(id).Error(0)
}

// Get implements ScheduleRepoMock.
func (u *ScheduleRepoMock) Get(id string) (model.ScheduleModel, error) {

	args := u.Called(id)
	if args.Get(1) != nil {
		return model.ScheduleModel{}, args.Error(1)
	}
	// var userR model.ScheduleModel = mappingutil.MappingUser(args.Get(0).(model.User))
	return args.Get(0).(model.ScheduleModel), nil
}

// List implements ScheduleRepoMock.
func (u *ScheduleRepoMock) List() ([]model.ScheduleModel, error) {
	args := u.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.ScheduleModel) , nil
}

// Update implements ScheduleRepoMock.
func (u *ScheduleRepoMock) Update(payload *model.ScheduleModel) error {
	return u.Called(payload).Error(0)
}

