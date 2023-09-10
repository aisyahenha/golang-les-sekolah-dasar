package repository

import (
	"fmt"

	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"gorm.io/gorm"
)

type TeacherRepository interface {
	Create(payload *model.TeacherModel) error
	List() ([]model.TeacherModel, error)
	Get(id string) (model.TeacherModel, error)
	Update(payload *model.TeacherModel) error
	Delete(id string) error
}

type teacherRepository struct {
	db *gorm.DB
}

// Create implements teacherRepository.
func (u *teacherRepository) Create(payload *model.TeacherModel) error {
	result := u.db.Create(payload).Error
	fmt.Print("sampeee createeee: ", result)
	return result
}

// Delete implements teacherRepository.
func (u *teacherRepository) Delete(id string) error {
	result := u.db.Where("id = ? ", id).Delete(&model.TeacherModel{}).Error
	return result
}

// Get implements teacherRepository.
func (u *teacherRepository) Get(id string) (model.TeacherModel, error) {

	var teacher model.TeacherModel

	result := u.db.Where("id = ?", id).First(&teacher).Error
	if result != nil {

		return model.TeacherModel{}, result
	}
	return teacher, nil
}

// List implements teacherRepository.
func (u *teacherRepository) List() ([]model.TeacherModel, error) {
	var teachers []model.TeacherModel
	result := u.db.Find(&teachers).Error
	if result != nil {
		return nil, result
	}
	return teachers, nil
}

// Update implements teacherRepository.
func (u *teacherRepository) Update(payload *model.TeacherModel) error {
	result := u.db.Model(&model.TeacherModel{}).Where("id = ?", &payload.ID).Updates(&payload).Error
	return result
}

func NewTeacherRepository(db *gorm.DB) TeacherRepository {
	return &teacherRepository{db: db}
}
