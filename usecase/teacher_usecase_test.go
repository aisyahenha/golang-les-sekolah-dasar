package usecase

import (
	"errors"
	"testing"
	"time"

	repomock "github.com/aisyahenha/golang-les-sekolah-dasar/_mock_/repo_mock"
	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type TeacherUseCaseTestSuite struct {
	suite.Suite
	repoMock *repomock.TeacherRepoMock
	usecase  TeacherUseCase
}

func (s *TeacherUseCaseTestSuite) SetupTest() {
	s.repoMock = new(repomock.TeacherRepoMock)
	s.usecase = NewTeacherUseCase(s.repoMock)
}

func TestTeacherUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(TeacherUseCaseTestSuite))
}

// --------------------- TEST CREATE
func (suite *TeacherUseCaseTestSuite) TestRegisterNewTeacher_Success() {
	payload := &model.TeacherModel{
		BaseModel: model.BaseModel{
			ID:        "1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{}},
		BasePersonalModel: model.BasePersonalModel{
			Name:    "Guru1",
			Address: "Rumah Guru1",
			Gender:  true},
			Specialist: "Biologi",
	}
	suite.repoMock.On("Create", payload).Return(nil)
	err := suite.usecase.RegisterNewTeacher(payload)
	assert.Nil(suite.T(), err)
}

func (suite *TeacherUseCaseTestSuite) TestRegisterNewTeacher_Fail() {
	payload := &model.TeacherModel{}
	suite.repoMock.On("Create", payload).Return(errors.New("error"))
	// payload.IsValidField()
	err := suite.usecase.RegisterNewTeacher(payload)
	assert.Error(suite.T(), err)
}

// --------------------- TEST SELECT ALL LIST
func (suite *TeacherUseCaseTestSuite) TestFindAllTeacher_Success() {
	expected := []model.TeacherModel{
		{
			BaseModel: model.BaseModel{
				ID:        "1",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{}},
			BasePersonalModel: model.BasePersonalModel{
				Name:    "Guru1",
				Address: "Rumah Guru1",
				Gender:  true},
				Specialist: "Biologi",
		},
	}
	suite.repoMock.On("List").Return(expected, nil)
	actual, err := suite.usecase.FindAllTeacher()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *TeacherUseCaseTestSuite) TestFindAllTeacher_Fail() {
	suite.repoMock.On("List").Return(nil, errors.New("error"))
	actual, err := suite.usecase.FindAllTeacher()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

// --------------------- TEST SELECT BY ID
func (suite *TeacherUseCaseTestSuite) TestFindById_Success() {
	expected := model.TeacherModel{
		BaseModel: model.BaseModel{
			ID:        "1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{}},
		BasePersonalModel: model.BasePersonalModel{
			Name:    "Guru1",
			Address: "Rumah Guru1",
			Gender:  true},
			Specialist: "Biologi",
	}
	suite.repoMock.On("Get", expected.ID).Return(expected, nil)
	actual, err := suite.usecase.FindByTeacherId(expected.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *TeacherUseCaseTestSuite) TestFindById_Fail() {
	ID := "1"
	suite.repoMock.On("Get", ID).Return(nil, errors.New("error"))
	actual, err := suite.usecase.FindByTeacherId(ID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.TeacherModel{}, actual)
}

//--------------------- TEST UPDATE
//--------------------- TEST DELETE
