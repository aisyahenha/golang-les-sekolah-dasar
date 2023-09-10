package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ScheduleRepositoryTestSuite struct {
	suite.Suite
	mockDB *gorm.DB
	mock   sqlmock.Sqlmock
	repo   ScheduleRepository
}

// --------------------- TEST CONFIG DB
func (suite *ScheduleRepositoryTestSuite) SetupTest() {
	conn, mock, err := sqlmock.New()
	assert.Nil(suite.T(), err)
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}))
	assert.Nil(suite.T(), err)
	suite.mockDB = db
	suite.mock = mock
	suite.repo = NewScheduleRepository(suite.mockDB)
}
func TestScheduleRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleRepositoryTestSuite))
}

// --------------------- TEST CREATE
func (suite *ScheduleRepositoryTestSuite) TestCreateSchedule_Success() {
	payload := &model.ScheduleModel{
		BaseModel: model.BaseModel{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}},
		CourseID:  "1",
		TeacherID: "1",
		StudentID: "1",
	}
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery("INSERT INTO").
		WithArgs(
			payload.CreatedAt,
			payload.UpdatedAt,
			payload.DeletedAt,
			payload.CourseID,
			payload.TeacherID,
			payload.StudentID,
			payload.ID,
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(payload.ID))
	suite.mock.ExpectCommit()
	actualErr := suite.repo.Create(payload)
	assert.NoError(suite.T(), actualErr)
	assert.Nil(suite.T(), actualErr)
}

func (suite *ScheduleRepositoryTestSuite) TestCreate_Error() {
	payload := &model.ScheduleModel{
		BaseModel: model.BaseModel{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}},
		CourseID:  "1",
		TeacherID: "1",
		StudentID: "1",
	}
	suite.mock.ExpectExec("INSERT INTO").WillReturnError(errors.New("error"))
	actualErr := suite.repo.Create(payload)
	assert.Error(suite.T(), actualErr)
}

// --------------------- TEST SELECT ALL LIST
func (suite *ScheduleRepositoryTestSuite) TestList_Success() {
	expected := []model.ScheduleModel{
		{
			BaseModel: model.BaseModel{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}},
			CourseID:  "1",
			TeacherID: "1",
			StudentID: "1",
		},
	}

	rowDummies := make([]model.ScheduleModel, len(expected))
	rows := sqlmock.NewRows([]string{"id", "course_id", "teacher_id", "student_id", "created_at", "updated_at"})
	for i, schedule := range expected {
		rowDummies[i] = schedule
		rows.AddRow(schedule.ID, schedule.CourseID, schedule.TeacherID, schedule.StudentID, schedule.CreatedAt, schedule.UpdatedAt)
	}
	expectedQuery := `SELECT (.+) FROM "schedule_models"`
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(rows)
	repo := NewScheduleRepository(suite.mockDB)
	actual, err := repo.List()
	assert.Equal(suite.T(), expected, actual)
	assert.NoError(suite.T(), err)
}

func (suite *ScheduleRepositoryTestSuite) TestList_Error() {
	expectedQuery := `SELECT (.+) FROM "schedule_models"`
	suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New("error"))
	repo := NewScheduleRepository(suite.mockDB)
	actual, err := repo.List()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

// --------------------- TEST SELECT BY ID
func (suite *ScheduleRepositoryTestSuite) TestGet_Success() {

	expected := []model.ScheduleModel{
		{
			BaseModel: model.BaseModel{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}},
			CourseID:  "1",
			TeacherID: "1",
			StudentID: "1",
		},
	}

	rowDummies := make([]model.ScheduleModel, len(expected))
	rows := sqlmock.NewRows([]string{"id", "course_id", "teacher_id", "student_id", "created_at", "updated_at"})
	for i, schedule := range expected {
		rowDummies[i] = schedule
		rows.AddRow(schedule.ID, schedule.CourseID, schedule.TeacherID, schedule.StudentID, schedule.CreatedAt, schedule.UpdatedAt)
	}
	expectedQuery := (`SELECT (.+) FROM "schedule_models"`)
	suite.mock.ExpectQuery(expectedQuery).WithArgs(expected[0].ID).WillReturnRows(rows)
	repo := NewScheduleRepository(suite.mockDB)
	actual, err := repo.Get(expected[0].ID)
	assert.Equal(suite.T(), expected[0], actual)
	assert.NoError(suite.T(), err)
}

func (suite *ScheduleRepositoryTestSuite) TestGet_Error() {
	ID := "1"
	expectedQuery := (`SELECT (.+) FROM "schedule_models"`)
	suite.mock.ExpectQuery(expectedQuery).WithArgs(ID).WillReturnError(errors.New("error"))
	repo := NewScheduleRepository(suite.mockDB)
	actual, err := repo.Get(ID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.ScheduleModel{}, actual)
}

//--------------------- TEST UPDATE
//--------------------- TEST DELETE
