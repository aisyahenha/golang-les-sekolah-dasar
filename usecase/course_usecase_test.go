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

type CourseUseCaseTestSuite struct {
	suite.Suite
	repoMock *repomock.CourseRepoMock
	usecase  CourseUseCase
}

func (s *CourseUseCaseTestSuite) SetupTest() {
	s.repoMock = new(repomock.CourseRepoMock)
	s.usecase = NewCourseUseCase(s.repoMock)
}

func TestCourseUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CourseUseCaseTestSuite))
}

// --------------------- TEST CREATE
func (suite *CourseUseCaseTestSuite) TestRegisterNewCourse_Success() {
	payload := &model.CourseModel{
		BaseModel: model.BaseModel{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}},
		Subject:   "IPA",
		Class:     "A1",
		Day:       "Monday",
		Start:     "09:00",
		End:       "10:00",
	}
	suite.repoMock.On("Create", payload).Return(nil)
	err := suite.usecase.RegisterNewCourse(payload)
	assert.Nil(suite.T(), err)
}

func (suite *CourseUseCaseTestSuite) TestRegisterNewCourse_Fail() {
	payload := &model.CourseModel{}
	suite.repoMock.On("Create", payload).Return(errors.New("error"))
	// payload.IsValidField()
	err := suite.usecase.RegisterNewCourse(payload)
	assert.Error(suite.T(), err)
}

// --------------------- TEST SELECT ALL LIST
func (suite *CourseUseCaseTestSuite) TestFindAllCourse_Success() {
	expected := []model.CourseModel{
		{
			BaseModel: model.BaseModel{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}},
			Subject:   "IPA",
			Class:     "A1",
			Day:       "Monday",
			Start:     "09:00",
			End:       "10:00",
		},
	}
	suite.repoMock.On("List").Return(expected, nil)
	actual, err := suite.usecase.FindAllCourse()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *CourseUseCaseTestSuite) TestFindAllCourse_Fail() {
	suite.repoMock.On("List").Return(nil, errors.New("error"))
	actual, err := suite.usecase.FindAllCourse()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

// --------------------- TEST SELECT BY ID
func (suite *CourseUseCaseTestSuite) TestFindById_Success() {
	expected := model.CourseModel{

		BaseModel: model.BaseModel{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}},
		Subject:   "IPA",
		Class:     "A1",
		Day:       "Monday",
		Start:     "09:00",
		End:       "10:00",
	}
	suite.repoMock.On("Get", expected.ID).Return(expected, nil)
	actual, err := suite.usecase.FindByCourseId(expected.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *CourseUseCaseTestSuite) TestFindById_Fail() {
	ID := "1"
	suite.repoMock.On("Get", ID).Return(nil, errors.New("error"))
	actual, err := suite.usecase.FindByCourseId(ID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.CourseModel{}, actual)
}

//--------------------- TEST UPDATE
//--------------------- TEST DELETE
