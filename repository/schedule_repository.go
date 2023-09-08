package repository

import (
	"fmt"

	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	Create(payload *model.ScheduleModel) error
	List() ([]model.ScheduleModel, error)
	Get(id string) (model.ScheduleModel, error)
	Update(payload *model.ScheduleModel) error
	Delete(id string) error
}

type scheduleRepository struct {
	db *gorm.DB
}

// Create implements ScheduleRepository.
func (u *scheduleRepository) Create(payload *model.ScheduleModel) error {
	result := u.db.Create(payload).Error
	fmt.Print("sampeee createeee: ", result)
	return result
}

// Delete implements ScheduleRepository.
func (u *scheduleRepository) Delete(id string) error {
	result := u.db.Where("id = ? ", id).Delete(&model.ScheduleModel{}).Error
	return result
}

// Get implements ScheduleRepository.
func (u *scheduleRepository) Get(id string) (model.ScheduleModel, error) {

	var Schedule model.ScheduleModel

	result := u.db.Where("id = ?", id).First(&Schedule).Error
	if result != nil {

		return model.ScheduleModel{}, result
	}
	return Schedule, nil
}

// List implements ScheduleRepository.
func (u *scheduleRepository) List() ([]model.ScheduleModel, error) {
	var Schedules []model.ScheduleModel
	result := u.db.Find(&Schedules).Error
	if result != nil {
		return nil, result
	}
	return Schedules, nil
}

// Update implements ScheduleRepository.
func (u *scheduleRepository) Update(payload *model.ScheduleModel) error {
	result := u.db.Model(&model.ScheduleModel{}).Where("id = ?", &payload.ID).Updates(&payload).Error
	return result
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db: db}
}
