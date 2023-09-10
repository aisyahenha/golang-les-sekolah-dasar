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

type TeacherRepositoryTestSuite struct {
	suite.Suite
	mockDB *gorm.DB
	mock   sqlmock.Sqlmock
	repo   TeacherRepository
}

// --------------------- TEST CONFIG DB
func (suite *TeacherRepositoryTestSuite) SetupTest() {
	conn, mock, err := sqlmock.New()
	assert.Nil(suite.T(), err)
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}))
	assert.Nil(suite.T(), err)
	suite.mockDB = db
	suite.mock = mock
	suite.repo = NewTeacherRepository(suite.mockDB)
}
func TestTeacherRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TeacherRepositoryTestSuite))
}

// --------------------- TEST CREATE
func (suite *TeacherRepositoryTestSuite) TestCreateTeacher_Success() {
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
		Specialist: "Ilmu Pendidikan - Biologi",
	}
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery("INSERT INTO").
		WithArgs(
			payload.CreatedAt,
			payload.UpdatedAt,
			payload.DeletedAt,
			payload.Name,
			payload.Address,
			payload.Gender,
			payload.Specialist,
			payload.ID,
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(payload.ID))
	suite.mock.ExpectCommit()
	actualErr := suite.repo.Create(payload)
	assert.NoError(suite.T(), actualErr)
	assert.Nil(suite.T(), actualErr)
}

func (suite *TeacherRepositoryTestSuite) TestCreate_Error() {
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
		Specialist: "Ilmu Pendidikan - Biologi",
	}
	suite.mock.ExpectExec("INSERT INTO").WillReturnError(errors.New("error"))
	actualErr := suite.repo.Create(payload)
	assert.Error(suite.T(), actualErr)
}

//--------------------- TEST SELECT ALL LIST

func (suite *TeacherRepositoryTestSuite) TestList_Success() {
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
			Specialist: "Ilmu Pendidikan - Biologi",
		},
	}

	rowDummies := make([]model.TeacherModel, len(expected))
	rows := sqlmock.NewRows([]string{"id", "name", "address", "gender","specialist", "created_at", "updated_at"})
	for i, teacher := range expected {
		rowDummies[i] = teacher
		rows.AddRow(teacher.ID, teacher.Name, teacher.Address, teacher.Gender,teacher.Specialist, teacher.CreatedAt, teacher.UpdatedAt)
	}
	expectedQuery := `SELECT (.+) FROM "teacher_models"`
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(rows)
	repo := NewTeacherRepository(suite.mockDB)
	actual, err := repo.List()
	assert.Equal(suite.T(), expected, actual)
	assert.NoError(suite.T(), err)
}

func (suite *TeacherRepositoryTestSuite) TestList_Error() {
	expectedQuery := `SELECT (.+) FROM "teacher_models"`
	suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New("error"))
	repo := NewTeacherRepository(suite.mockDB)
	actual, err := repo.List()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

//--------------------- TEST SELECT BY ID

func (suite *TeacherRepositoryTestSuite) TestGet_Success() {

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
			Specialist: "Ilmu Pendidikan - Biologi",
		},
	}


	rowDummies := make([]model.TeacherModel, len(expected))
	rows := sqlmock.NewRows([]string{"id", "name", "address", "gender","specialist", "created_at", "updated_at"})
	for i, teacher := range expected {
		rowDummies[i] = teacher
		rows.AddRow(teacher.ID, teacher.Name, teacher.Address, teacher.Gender,teacher.Specialist, teacher.CreatedAt, teacher.UpdatedAt)
	}
	expectedQuery := (`SELECT (.+) FROM "teacher_models"`)
	suite.mock.ExpectQuery(expectedQuery).WithArgs(expected[0].ID).WillReturnRows(rows)
	repo := NewTeacherRepository(suite.mockDB)
	actual, err := repo.Get(expected[0].ID)
	assert.Equal(suite.T(), expected[0], actual)
	assert.NoError(suite.T(), err)
}

func (suite *TeacherRepositoryTestSuite) TestGet_Error() {
	ID := "1"
	expectedQuery := (`SELECT (.+) FROM "teacher_models"`)
	suite.mock.ExpectQuery(expectedQuery).WithArgs(ID).WillReturnError(errors.New("error"))
	repo := NewTeacherRepository(suite.mockDB)
	actual, err := repo.Get(ID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.TeacherModel{}, actual)
}


//--------------------- TEST UPDATE
//--------------------- TEST DELETE
