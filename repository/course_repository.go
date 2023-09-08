package repository

import (
	"fmt"

	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Create(payload *model.CourseModel) error
	List() ([]model.CourseModel, error)
	Get(id string) (model.CourseModel, error)
	Update(payload *model.CourseModel) error
	Delete(id string) error
}

type courseRepository struct {
	db *gorm.DB
}

// Create implements CourseRepository.
func (u *courseRepository) Create(payload *model.CourseModel) error {
	result := u.db.Create(payload).Error
	fmt.Print("sampeee createeee: ", result)
	return result
}

// Delete implements CourseRepository.
func (u *courseRepository) Delete(id string) error {
	result := u.db.Where("id = ? ", id).Delete(&model.CourseModel{}).Error
	return result
}

// Get implements CourseRepository.
func (u *courseRepository) Get(id string) (model.CourseModel, error) {

	var Course model.CourseModel

	result := u.db.Where("id = ?", id).First(&Course).Error
	if result != nil {

		return model.CourseModel{}, result
	}
	return Course, nil
}

// List implements CourseRepository.
func (u *courseRepository) List() ([]model.CourseModel, error) {
	var Courses []model.CourseModel
	result := u.db.Find(&Courses).Error
	if result != nil {
		return nil, result
	}
	return Courses, nil
}

// Update implements CourseRepository.
func (u *courseRepository) Update(payload *model.CourseModel) error {
	result := u.db.Model(&model.CourseModel{}).Where("id = ?", &payload.ID).Updates(&payload).Error
	return result
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}
