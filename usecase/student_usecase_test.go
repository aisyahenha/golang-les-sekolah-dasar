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

type StudentUseCaseTestSuite struct {
	suite.Suite
	repoMock *repomock.StudentRepoMock
	usecase  StudentUseCase
}

func (s *StudentUseCaseTestSuite) SetupTest() {
	s.repoMock = new(repomock.StudentRepoMock)
	s.usecase = NewStudentUseCase(s.repoMock)
}

func TestStudentUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(StudentUseCaseTestSuite))
}

// --------------------- TEST CREATE
func (suite *StudentUseCaseTestSuite) TestRegisterNewStudent_Success() {
	payload := &model.StudentModel{
		BaseModel: model.BaseModel{
			ID:        "1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{}},
		BasePersonalModel: model.BasePersonalModel{
			Name:    "Siswa1",
			Address: "Rumah Siswa1",
			Gender:  true},
	}
	suite.repoMock.On("Create", payload).Return(nil)
	err := suite.usecase.RegisterNewStudent(payload)
	assert.Nil(suite.T(), err)
}

func (suite *StudentUseCaseTestSuite) TestRegisterNewStudent_Fail() {
	payload := &model.StudentModel{}
	suite.repoMock.On("Create", payload).Return(errors.New("error"))
	// payload.IsValidField()
	err := suite.usecase.RegisterNewStudent(payload)
	assert.Error(suite.T(), err)
}

// --------------------- TEST SELECT ALL LIST
func (suite *StudentUseCaseTestSuite) TestFindAllStudent_Success() {
	expected := []model.StudentModel{
		{
			BaseModel: model.BaseModel{
				ID:        "1",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{}},
			BasePersonalModel: model.BasePersonalModel{
				Name:    "Siswa1",
				Address: "Rumah Siswa1",
				Gender:  true},
		},
	}
	suite.repoMock.On("List").Return(expected, nil)
	actual, err := suite.usecase.FindAllStudent()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *StudentUseCaseTestSuite) TestFindAllStudent_Fail() {
	suite.repoMock.On("List").Return(nil, errors.New("error"))
	actual, err := suite.usecase.FindAllStudent()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

// --------------------- TEST SELECT BY ID
func (suite *StudentUseCaseTestSuite) TestFindById_Success() {
	expected := model.StudentModel{
		BaseModel: model.BaseModel{
			ID:        "1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{}},
		BasePersonalModel: model.BasePersonalModel{
			Name:    "Siswa1",
			Address: "Rumah Siswa1",
			Gender:  true},
	}
	suite.repoMock.On("Get", expected.ID).Return(expected, nil)
	actual, err := suite.usecase.FindByStudentId(expected.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *StudentUseCaseTestSuite) TestFindById_Fail() {
	ID := "1"
	suite.repoMock.On("Get", ID).Return(nil, errors.New("error"))
	actual, err := suite.usecase.FindByStudentId(ID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.StudentModel{}, actual)
}

//--------------------- TEST UPDATE
//--------------------- TEST DELETE
