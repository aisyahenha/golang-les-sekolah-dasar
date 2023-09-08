package repository

import (
	"fmt"

	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"gorm.io/gorm"
)

type StudentRepository interface {
	Create(payload *model.StudentModel) error
	List() ([]model.StudentModel, error)
	Get(id string) (model.StudentModel, error)
	Update(payload *model.StudentModel) error
	Delete(id string) error
}

type studentRepository struct {
	db *gorm.DB
}

// Create implements StudentRepository.
func (u *studentRepository) Create(payload *model.StudentModel) error {
	result := u.db.Create(payload).Error
	fmt.Print("sampeee createeee: ", result)
	return result
}

// Delete implements studentRepository.
func (u *studentRepository) Delete(id string) error {
	result := u.db.Where("id = ? ", id).Delete(&model.StudentModel{}).Error
	return result
}

// Get implements studentRepository.
func (u *studentRepository) Get(id string) (model.StudentModel, error) {

	var Student model.StudentModel

	result := u.db.Where("id = ?", id).First(&Student).Error
	if result != nil {

		return model.StudentModel{}, result
	}
	return Student, nil
}

// List implements studentRepository.
func (u *studentRepository) List() ([]model.StudentModel, error) {
	var Students []model.StudentModel
	result := u.db.Find(&Students).Error
	if result != nil {
		return nil, result
	}
	return Students, nil
}

// Update implements studentRepository.
func (u *studentRepository) Update(payload *model.StudentModel) error {
	result := u.db.Model(&model.StudentModel{}).Where("id = ?", &payload.ID).Updates(&payload).Error
	return result
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db: db}
}
