package usecase

import (
	// "errors"

	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/aisyahenha/golang-les-sekolah-dasar/repository"
)

type ScheduleUseCase interface {
	RegisterNewSchedule(payload *model.ScheduleModel) error
	FindAllSchedule() ([]model.ScheduleModel, error)
	FindByScheduleId(id string) (model.ScheduleModel, error)
	UpdateSchedule(payload *model.ScheduleModel) error
	DeleteSchedule(id string) error
}

type scheduleUseCase struct {
	repo repository.ScheduleRepository
}

// DeleteSchedule implements ScheduleUseCase.
func (u *scheduleUseCase) DeleteSchedule(id string) error {
	return u.repo.Delete(id)
}

// FindAllSchedule implements ScheduleUseCase.
func (u *scheduleUseCase) FindAllSchedule() ([]model.ScheduleModel, error) {
	return u.repo.List()
}

// FindByScheduleId implements ScheduleUseCase.
func (u *scheduleUseCase) FindByScheduleId(id string) (model.ScheduleModel, error) {
	return u.repo.Get(id)
}

// RegisterNewSchedule implements ScheduleUseCase.
func (u *scheduleUseCase) RegisterNewSchedule(payload *model.ScheduleModel) error {

	return u.repo.Create(payload)
}

// UpdateSchedule implements ScheduleUseCase.
func (u *scheduleUseCase) UpdateSchedule(payload *model.ScheduleModel) error {
	return u.repo.Update(payload)
}

func NewScheduleUseCase(repo repository.ScheduleRepository) ScheduleUseCase {
	return &scheduleUseCase{repo: repo}
}
