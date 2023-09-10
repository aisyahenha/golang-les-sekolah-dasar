package usecase

import (
	"errors"
	"testing"
	"time"

	repomock "github.com/aisyahenha/golang-les-sekolah-dasar/_mock_/repo_mock"
	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	repoMock *repomock.UserRepoMock
	usecase  UserUseCase
}

func (s *UserUseCaseTestSuite) SetupTest() {
	s.repoMock = new(repomock.UserRepoMock)
	s.usecase = NewUserUseCase(s.repoMock)
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}

// --------------------- TEST CREATE
func (suite *UserUseCaseTestSuite) TestRegisterNewUser_Success() {
	payload := &model.User{
		BaseModel: model.BaseModel{
			ID:        "1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username: "admin",
		Password: "password",
		Role:     "admin",
	}
	suite.repoMock.On("Create", payload).Return(nil)
	err := suite.usecase.RegisterNewUser(payload)
	assert.Nil(suite.T(), err)
}

func (suite *UserUseCaseTestSuite) TestRegisterNewUser_Fail() {
	payload := &model.User{}
	suite.repoMock.On("Create", payload).Return(errors.New("error"))
	// payload.IsValidField()
	err := suite.usecase.RegisterNewUser(payload)
	assert.Error(suite.T(), err)
}

// --------------------- TEST SELECT ALL LIST
func (suite *UserUseCaseTestSuite) TestFindAllUser_Success() {
	expected := []model.UserRespon{
		{

			ID:        "1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Username:  "admin",
			Role:      "admin",
		},
	}
	suite.repoMock.On("List").Return(expected, nil)
	actual, err := suite.usecase.FindAllUser()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *UserUseCaseTestSuite) TestFindAllUser_Fail() {
	suite.repoMock.On("List").Return(nil, errors.New("error"))
	actual, err := suite.usecase.FindAllUser()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

// --------------------- TEST SELECT BY ID
func (suite *UserUseCaseTestSuite) TestFindById_Success() {
	expected := model.UserRespon{

		ID:        "1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  "admin",
		Role:      "admin",
	}
	suite.repoMock.On("Get", expected.ID).Return(expected, nil)
	actual, err := suite.usecase.FindByUserId(expected.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *UserUseCaseTestSuite) TestFindById_Fail() {
	ID := "1"
	suite.repoMock.On("Get", ID).Return(nil, errors.New("error"))
	actual, err := suite.usecase.FindByUserId(ID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.UserRespon{}, actual)
}

//--------------------- TEST Get By Username

func (suite *UserUseCaseTestSuite) TestFindUsername_Success() {
	expected := model.User{
		BaseModel: model.BaseModel{
			ID:        "1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username: "admin",
		Password: "password",
		Role:     "admin",
	}
	suite.repoMock.On("GetByUsername", expected.Username).Return(expected, nil)
	actual, err := suite.usecase.FindByUsername(expected.Username)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *UserUseCaseTestSuite) TestFindUsername_Fail() {
	Username := "1"
	suite.repoMock.On("GetByUsername", Username).Return(nil, errors.New("error"))
	actual, err := suite.usecase.FindByUsername(Username)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.User{}, actual)
}

// --------------------- TEST Get By Username Password
func (suite *UserUseCaseTestSuite) TestFindUsernamePass_Success() {
	expected := model.User{
		BaseModel: model.BaseModel{
			ID:        "1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username: "admin",
		Password: "password",
		Role:     "admin",
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(expected.Password), bcrypt.DefaultCost)
	expected.Password = string(bytes)

	password := "password"
	suite.repoMock.On("GetByUsernamePassword", expected.Username, password).Return(expected, nil)
	actual, err := suite.usecase.FindByUsernamePassword(expected.Username, password)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *UserUseCaseTestSuite) TestFindUsernamePass_Fail() {
	username := "admin"
	password:="password"
	suite.repoMock.On("GetByUsernamePassword", username,password).Return(nil, errors.New("error"))
	actual, err := suite.usecase.FindByUsernamePassword(username,password)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.User{}, actual)
}
//--------------------- TEST UPDATE
//--------------------- TEST DELETE
