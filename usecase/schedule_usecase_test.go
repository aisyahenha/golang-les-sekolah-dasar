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

type ScheduleUseCaseTestSuite struct {
	suite.Suite
	repoMock *repomock.ScheduleRepoMock
	usecase  ScheduleUseCase
}

func (s *ScheduleUseCaseTestSuite) SetupTest() {
	s.repoMock = new(repomock.ScheduleRepoMock)
	s.usecase = NewScheduleUseCase(s.repoMock)
}

func TestScheduleUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleUseCaseTestSuite))
}

// --------------------- TEST CREATE
func (suite *ScheduleUseCaseTestSuite) TestRegisterNewSchedule_Success() {
	payload := &model.ScheduleModel{
		BaseModel: model.BaseModel{
			ID:        "1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{}},
		CourseID:  "1",
		TeacherID: "1",
		StudentID: "1",
	}
	suite.repoMock.On("Create", payload).Return(nil)
	err := suite.usecase.RegisterNewSchedule(payload)
	assert.Nil(suite.T(), err)
}

func (suite *ScheduleUseCaseTestSuite) TestRegisterNewSchedule_Fail() {
	payload := &model.ScheduleModel{}
	suite.repoMock.On("Create", payload).Return(errors.New("error"))
	// payload.IsValidField()
	err := suite.usecase.RegisterNewSchedule(payload)
	assert.Error(suite.T(), err)
}

// --------------------- TEST SELECT ALL LIST
func (suite *ScheduleUseCaseTestSuite) TestFindAllSchedule_Success() {
	expected := []model.ScheduleModel{
		{
			BaseModel: model.BaseModel{
				ID:        "1",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{}},
			CourseID:  "1",
			TeacherID: "1",
			StudentID: "1"},
	}
	suite.repoMock.On("List").Return(expected, nil)
	actual, err := suite.usecase.FindAllSchedule()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *ScheduleUseCaseTestSuite) TestFindAllSchedule_Fail() {
	suite.repoMock.On("List").Return(nil, errors.New("error"))
	actual, err := suite.usecase.FindAllSchedule()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

// --------------------- TEST SELECT BY ID
func (suite *ScheduleUseCaseTestSuite) TestFindById_Success() {
	expected := model.ScheduleModel{
		BaseModel: model.BaseModel{
			ID:        "1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{}},
			CourseID:  "1",
			TeacherID: "1",
			StudentID: "1",
		
	}
	suite.repoMock.On("Get", expected.ID).Return(expected, nil)
	actual, err := suite.usecase.FindByScheduleId(expected.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *ScheduleUseCaseTestSuite) TestFindById_Fail() {
	ID := "1"
	suite.repoMock.On("Get", ID).Return(nil, errors.New("error"))
	actual, err := suite.usecase.FindByScheduleId(ID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.ScheduleModel{}, actual)
}

//--------------------- TEST UPDATE
//--------------------- TEST DELETE
