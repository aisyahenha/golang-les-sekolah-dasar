package repository

import (
	"fmt"
	"unsafe"

	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	mappingutil "github.com/aisyahenha/golang-les-sekolah-dasar/utils/maping_util"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(payload *model.User) error
	List() ([]model.UserRespon, error)
	// List() ([]model.User, error)
	Get(id string) (model.UserRespon, error)
	Update(payload *model.User) error
	Delete(id string) error
}

type userRepository struct {
	db *gorm.DB
}

// Create implements UserRepository.
func (u *userRepository) Create(payload *model.User) error {
	result := u.db.Create(payload).Error
	fmt.Print("sampeee createeee: ", result)
	return result
}

// Delete implements UserRepository.
func (u *userRepository) Delete(id string) error {
	result := u.db.Where("id = ? ", id).Delete(&model.User{}).Error
	return result
}

// Get implements UserRepository.
func (u *userRepository) Get(id string) (model.UserRespon, error) {

	// var userRsp model.UserRespon
	var user model.User

	result := u.db.Where("id = ?", id).First(&user).Error
	if result != nil {

		return model.UserRespon{}, result
	}
	var userR model.UserRespon = mappingutil.MappingUser(user)
	return userR, nil
}

// List implements UserRepository.
func (u *userRepository) List() ([]model.UserRespon, error) {

	// func (u *userRepository) List() ([]model.User, error) {
	var users []model.User
	var userR []model.UserRespon
	userLength := unsafe.Sizeof(users)
	result := u.db.Find(&users).Error
	if result != nil {
		return nil, result
	}

	fmt.Println("panjang array nyaaaa: ", userLength)
	// for _, student := range allStudents {
	// 	fmt.Println(student.name, "age is", student.age)
	// }
	i := 0
	for _, user := range users {
		userR = append(userR, mappingutil.MappingUser(user))
		i++
	}
	// for i := 0; i < int(userLength); i++{
	// 	fmt.Print(users[i])
	// 	userR[i]= mappingutil.MappingUser(users[i])
	// }
	// fmt.Println("data ke 0 nya: ", users[0])
	return userR, nil
}

// Update implements UserRepository.
func (u *userRepository) Update(payload *model.User) error {
	result := u.db.Model(&model.User{}).Where("id = ?", &payload.ID).Updates(&payload).Error
	return result
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
