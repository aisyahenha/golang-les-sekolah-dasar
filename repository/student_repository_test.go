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

type StudentRepositoryTestSuite struct {
	suite.Suite
	mockDB *gorm.DB
	mock   sqlmock.Sqlmock
	repo   StudentRepository
}

// --------------------- TEST CONFIG DB
func (suite *StudentRepositoryTestSuite) SetupTest() {
	conn, mock, err := sqlmock.New()
	assert.Nil(suite.T(), err)
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}))
	assert.Nil(suite.T(), err)
	suite.mockDB = db
	suite.mock = mock
	suite.repo = NewStudentRepository(suite.mockDB)
}
func TestStudentRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(StudentRepositoryTestSuite))
}

// --------------------- TEST CREATE
func (suite *StudentRepositoryTestSuite) TestCreateStudent_Success() {
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
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery("INSERT INTO").
		WithArgs(
			payload.CreatedAt,
			payload.UpdatedAt,
			payload.DeletedAt,
			payload.Name,
			payload.Address,
			payload.Gender,
			payload.ID,
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(payload.ID))
	suite.mock.ExpectCommit()
	actualErr := suite.repo.Create(payload)
	assert.NoError(suite.T(), actualErr)
	assert.Nil(suite.T(), actualErr)
}

func (suite *StudentRepositoryTestSuite) TestCreate_Error() {
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
	suite.mock.ExpectExec("INSERT INTO").WillReturnError(errors.New("error"))
	actualErr := suite.repo.Create(payload)
	assert.Error(suite.T(), actualErr)
}

// --------------------- TEST SELECT ALL LIST
func (suite *StudentRepositoryTestSuite) TestList_Success() {
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

	rowDummies := make([]model.StudentModel, len(expected))
	rows := sqlmock.NewRows([]string{"id", "name", "address", "gender", "created_at", "updated_at"})
	for i, student := range expected {
		rowDummies[i] = student
		rows.AddRow(student.ID, student.Name, student.Address, student.Gender, student.CreatedAt, student.UpdatedAt)
	}
	expectedQuery := `SELECT (.+) FROM "student_models"`
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(rows)
	repo := NewStudentRepository(suite.mockDB)
	actual, err := repo.List()
	assert.Equal(suite.T(), expected, actual)
	assert.NoError(suite.T(), err)
}

func (suite *StudentRepositoryTestSuite) TestList_Error() {
	expectedQuery := `SELECT (.+) FROM "student_models"`
	suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New("error"))
	repo := NewStudentRepository(suite.mockDB)
	actual, err := repo.List()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

//--------------------- TEST SELECT BY ID

func (suite *StudentRepositoryTestSuite) TestGet_Success() {

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

	rowDummies := make([]model.StudentModel, len(expected))
	rows := sqlmock.NewRows([]string{"id", "name", "address", "gender", "created_at", "updated_at"})
	for i, student := range expected {
		rowDummies[i] = student
		rows.AddRow(student.ID, student.Name, student.Address, student.Gender, student.CreatedAt, student.UpdatedAt)
	}
	expectedQuery := (`SELECT (.+) FROM "student_models"`)
	suite.mock.ExpectQuery(expectedQuery).WithArgs(expected[0].ID).WillReturnRows(rows)
	repo := NewStudentRepository(suite.mockDB)
	actual, err := repo.Get(expected[0].ID)
	assert.Equal(suite.T(), expected[0], actual)
	assert.NoError(suite.T(), err)
}

func (suite *StudentRepositoryTestSuite) TestGet_Error() {
	ID := "1"
	expectedQuery := (`SELECT (.+) FROM "student_models"`)
	suite.mock.ExpectQuery(expectedQuery).WithArgs(ID).WillReturnError(errors.New("error"))
	repo := NewStudentRepository(suite.mockDB)
	actual, err := repo.Get(ID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.StudentModel{}, actual)
}

//--------------------- TEST UPDATE
//--------------------- TEST DELETE
