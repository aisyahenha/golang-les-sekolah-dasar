package repository

import (
	"errors"
	// "fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CourseRepositoryTestSuite struct {
	suite.Suite
	mockDB *gorm.DB
	mock   sqlmock.Sqlmock
	repo   CourseRepository
}

// --------------------- TEST CONFIG DB
func (suite *CourseRepositoryTestSuite) SetupTest() {
	conn, mock, err := sqlmock.New()
	assert.Nil(suite.T(), err)
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}))
	assert.Nil(suite.T(), err)
	suite.mockDB = db
	suite.mock = mock
	suite.repo = NewCourseRepository(suite.mockDB)
}
func TestCourseRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CourseRepositoryTestSuite))
}

// --------------------- TEST CREATE
func (suite *CourseRepositoryTestSuite) TestCreateCourse_Success() {
	payload := &model.CourseModel{
		BaseModel: model.BaseModel{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}},
		Subject:   "IPA",
		Class:     "A1",
		Day:       "Monday",
		Start:     "0900",
		End:       "1000",
	}
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery("INSERT INTO").
		WithArgs(
			payload.CreatedAt,
			payload.UpdatedAt,
			payload.DeletedAt,
			payload.Subject,
			payload.Class,
			payload.Day,
			payload.Start,
			payload.End,
			payload.ID,
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(payload.ID))
	suite.mock.ExpectCommit()
	actualErr := suite.repo.Create(payload)
	assert.NoError(suite.T(), actualErr)
	assert.Nil(suite.T(), actualErr)
}

func (suite *CourseRepositoryTestSuite) TestCreate_Error() {
	payload := &model.CourseModel{
		BaseModel: model.BaseModel{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}},
		Subject:   "IPA",
		Class:     "A1",
		Day:       "Monday",
		Start:     "09:00",
		End:       "10:00",
	}
	suite.mock.ExpectExec("INSERT INTO").WillReturnError(errors.New("error"))
	actualErr := suite.repo.Create(payload)
	assert.Error(suite.T(), actualErr)
}

// --------------------- TEST SELECT ALL LIST
func (suite *CourseRepositoryTestSuite) TestList_Success() {
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

	rowDummies := make([]model.CourseModel, len(expected))
	rows := sqlmock.NewRows([]string{"id", "subject", "class", "day", "start", "end", "created_at", "updated_at"})
	for i, course := range expected {
		rowDummies[i] = course
		rows.AddRow(course.ID, course.Subject, course.Class, course.Day, course.Start, course.End, course.CreatedAt, course.UpdatedAt)
	}
	expectedQuery := `SELECT (.+) FROM "course_models"`
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(rows)
	repo := NewCourseRepository(suite.mockDB)
	actual, err := repo.List()
	assert.Equal(suite.T(), expected, actual)
	assert.NoError(suite.T(), err)
}

func (suite *CourseRepositoryTestSuite) TestList_Error() {
	expectedQuery := `SELECT (.+) FROM "course_models"`
	suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New("error"))
	repo := NewCourseRepository(suite.mockDB)
	actual, err := repo.List()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

// --------------------- TEST SELECT BY ID
func (suite *CourseRepositoryTestSuite) TestGet_Success() {
	
	expected := []model.CourseModel{
		{
			BaseModel: model.BaseModel{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}},
			Subject:   "IPA",
			Class:     "A1",
			Day:       "Monday",
			Start:     "09:00",
			End:       "10:00",
		}, 
		{
				BaseModel: model.BaseModel{ID: "2", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}},
				Subject:   "IPS",
				Class:     "A1",
				Day:       "Monday",
				Start:     "11:00",
				End:       "12:00",
			},
		}

		rowDummies := make([]model.CourseModel, len(expected))
		rows := sqlmock.NewRows([]string{"id", "subject", "class", "day", "start", "end", "created_at", "updated_at"})
		// rows.AddRow(expected[0])
		for i, course := range expected {
				rowDummies[i] = course
				rows.AddRow(course.ID, course.Subject, course.Class, course.Day, course.Start, course.End, course.CreatedAt, course.UpdatedAt)
			}
			expectedQuery := (`SELECT (.+) FROM "course_models"`)
			suite.mock.ExpectQuery(expectedQuery).WithArgs(expected[0].ID).WillReturnRows(rows)
			repo := NewCourseRepository(suite.mockDB)
			actual, err := repo.Get(expected[0].ID)
			assert.Equal(suite.T(), expected[0], actual)
			assert.NoError(suite.T(), err)
}


func (suite *CourseRepositoryTestSuite) TestGet_Error() {
		ID := "1"
		expectedQuery := (`SELECT (.+) FROM "course_models"`)
		suite.mock.ExpectQuery(expectedQuery).WithArgs(ID).WillReturnError(errors.New("error"))
		repo := NewCourseRepository(suite.mockDB)
		actual, err := repo.Get(ID)
		assert.Error(suite.T(), err)
		assert.Equal(suite.T(),model.CourseModel{}, actual)
}

//--------------------- TEST UPDATE
//--------------------- TEST DELETE