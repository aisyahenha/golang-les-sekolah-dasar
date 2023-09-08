package usecase

import (
	"errors"

	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/aisyahenha/golang-les-sekolah-dasar/repository"
	// "github.com/google/uuid"
)

type UserUseCase interface {
	RegisterNewUser(payload *model.User) error
	FindAllUser() ([]model.User, error)
	FindByUserId(id string) (model.User, error)
	UpdateUser(payload *model.User) error
	DeleteUser(id string) error
}

type userUseCase struct {
	repo repository.UserRepository
}

// DeleteUser implements UserUseCase.
func (u *userUseCase) DeleteUser(id string) error {
	return u.repo.Delete(id)
}

// FindAllUser implements UserUseCase.
func (u *userUseCase) FindAllUser() ([]model.User, error) {
	return u.repo.List()
}

// FindByUserId implements UserUseCase.
func (u *userUseCase) FindByUserId(id string) (model.User, error) {
	return u.repo.Get(id)
}

// RegisterNewUser implements UserUseCase.
func (u *userUseCase) RegisterNewUser(payload *model.User) error {

	if payload.Username == "" || payload.Password == "" {
		return errors.New("username and password are required fields")
	}

	return u.repo.Create(payload)
}

// UpdateUser implements UserUseCase.
func (u *userUseCase) UpdateUser(payload *model.User) error {
	return u.repo.Update(payload)
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
