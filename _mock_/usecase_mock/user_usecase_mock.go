package usecasemock

import (
	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/stretchr/testify/mock"
)

type UserUseCaseMock struct {
	mock.Mock
}

func (u *UserUseCaseMock) DeleteUser(id string) error {
	return u.Called(id).Error(0)
}

func (u *UserUseCaseMock) FindAllUser() ([]model.UserRespon, error) {
	args := u.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.UserRespon), nil
}

func (u *UserUseCaseMock) FindByUserId(id string) (model.UserRespon, error) {
	args := u.Called(id)
	if args.Get(1) != nil {
		return model.UserRespon{}, args.Error(1)
	}
	return args.Get(0).(model.UserRespon), nil
}

func (u *UserUseCaseMock) RegisterNewUser(payload *model.User) error {
	return u.Called(payload).Error(0)
}

func (u *UserUseCaseMock) FindByUsername(username string) (model.User, error) {
	args := u.Called(username)
	if args.Get(1) != nil {
		return model.User{}, args.Error(1)
	}
	return args.Get(0).(model.User), nil
}

func (u *UserUseCaseMock) FindByUsernamePassword(username string, password string) (model.User, error) {
	args := u.Called(username, password)
	if args.Get(1) != nil {
		return model.User{}, args.Error(1)
	}
	return args.Get(0).(model.User), nil
}

func (u *UserUseCaseMock) UpdateUser(payload *model.User) error {
	return u.Called(payload).Error(0)
}
func (u *UserUseCaseMock) LoginUser(username string, password string) (string, error) {
	args := u.Called(username, password)
	if args.Get(1) != nil {
		return "", args.Error(1)
	}
	return args.String(0), nil
}
