package usecasemock

import (
	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/stretchr/testify/mock"
)

type ScheduleUseCaseMock struct {
	mock.Mock
}
// DeleteSchedule implements ScheduleUseCase.
func (u *ScheduleUseCaseMock) DeleteSchedule(id string) error {
	return u.Called(id).Error(0)
}

// FindAllSchedule implements ScheduleUseCaseMock.
func (u *ScheduleUseCaseMock) FindAllSchedule() ([]model.ScheduleModel, error) {
	args := u.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.ScheduleModel), nil
}

// FindByScheduleId implements ScheduleUseCaseMock.
func (u *ScheduleUseCaseMock) FindByScheduleId(id string) (model.ScheduleModel, error) {
	args := u.Called(id)
	if args.Get(1) != nil {
		return model.ScheduleModel{}, args.Error(1)
	}
	return args.Get(0).(model.ScheduleModel), nil
}
// RegisterNewSchedule implements ScheduleUseCaseMock.
func (u *ScheduleUseCaseMock) RegisterNewSchedule(payload *model.ScheduleModel) error {

	return u.Called(payload).Error(0)
}

// UpdateSchedule implements ScheduleUseCaseMock.
func (u *ScheduleUseCaseMock) UpdateSchedule(payload *model.ScheduleModel) error {
	return u.Called(payload).Error(0)
}
