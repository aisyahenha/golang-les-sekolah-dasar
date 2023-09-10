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

type UserRepositoryTestSuite struct {
	suite.Suite
	mockDB *gorm.DB
	mock   sqlmock.Sqlmock
	repo   UserRepository
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	conn, mock, err := sqlmock.New()
	assert.Nil(suite.T(), err)
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}))
	assert.Nil(suite.T(), err)
	suite.mockDB = db
	suite.mock = mock
	suite.repo = NewUserRepository(suite.mockDB)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

// --------------------- TEST CREATE
func (suite *UserRepositoryTestSuite) TestCreate_Success() {
	payload := &model.User{
		BaseModel:  model.BaseModel{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}},
		Username:   "admin",
		Password:   "password",
		Role:       "admin",
		ResetToken: "",
	}
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery("INSERT INTO").
		WithArgs(
			payload.CreatedAt,
			payload.UpdatedAt,
			payload.DeletedAt,
			payload.Username,
			payload.Password,
			payload.Role,
			payload.ResetToken,
			payload.ID,
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(payload.ID))
	suite.mock.ExpectCommit()
	actualErr := suite.repo.Create(payload)
	assert.NoError(suite.T(), actualErr)
	assert.Nil(suite.T(), actualErr)
}

func (suite *UserRepositoryTestSuite) TestCreate_Error() {
	payload := &model.User{
		BaseModel:  model.BaseModel{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}},
		Username:   "admin",
		Password:   "password",
		Role:       "admin",
		ResetToken: "",
	}
	suite.mock.ExpectExec("INSERT INTO").WillReturnError(errors.New("error"))
	actualErr := suite.repo.Create(payload)
	assert.Error(suite.T(), actualErr)
}

// --------------------- TEST SELECT ALL LIST
func (suite *UserRepositoryTestSuite) TestList_Success() {
	expected := []model.UserRespon{
		{
			ID:         "1",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			Username:   "admin",
			Role:       "admin",
			ResetToken: "",
		},
	}

	rowDummies := make([]model.UserRespon, len(expected))
	rows := sqlmock.NewRows([]string{"id", "username", "role", "reser_token", "created_at", "updated_at"})
	for i, user := range expected {
		rowDummies[i] = user
		rows.AddRow(user.ID, user.Username, user.Role, user.ResetToken, user.CreatedAt, user.UpdatedAt)
	}
	expectedQuery := `SELECT (.+) FROM "users"`
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(rows)
	repo := NewUserRepository(suite.mockDB)
	actual, err := repo.List()
	assert.Equal(suite.T(), expected, actual)
	assert.NoError(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestList_Error() {
	expectedQuery := `SELECT (.+) FROM "users"`
	suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New("error"))
	repo := NewUserRepository(suite.mockDB)
	actual, err := repo.List()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

//--------------------- TEST SELECT BY ID

func (suite *UserRepositoryTestSuite) TestGet_Success() {

	expected := []model.UserRespon{
		{
			ID:         "1",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			Username:   "admin",
			Role:       "admin",
			ResetToken: "",
		},
	}

	rowDummies := make([]model.UserRespon, len(expected))
	rows := sqlmock.NewRows([]string{"id", "username", "role", "reser_token", "created_at", "updated_at"})
	for i, user := range expected {
		rowDummies[i] = user
		rows.AddRow(user.ID, user.Username, user.Role, user.ResetToken, user.CreatedAt, user.UpdatedAt)
	}
	expectedQuery := (`SELECT (.+) FROM "users"`)
	suite.mock.ExpectQuery(expectedQuery).WithArgs(expected[0].ID).WillReturnRows(rows)
	repo := NewUserRepository(suite.mockDB)
	actual, err := repo.Get(expected[0].ID)
	assert.Equal(suite.T(), expected[0], actual)
	assert.NoError(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestGet_Error() {
	ID := "1"
	expectedQuery := (`SELECT (.+) FROM "users"`)
	suite.mock.ExpectQuery(expectedQuery).WithArgs(ID).WillReturnError(errors.New("error"))
	repo := NewUserRepository(suite.mockDB)
	actual, err := repo.Get(ID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.UserRespon{}, actual)
}

// --------------------- TEST GET BY USERNAME
func (suite *UserRepositoryTestSuite) TestUsername_Success() {

	expected := []model.User{{
		BaseModel:  model.BaseModel{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}},
		Username:   "admin",
		Password:   "password",
		Role:       "admin",
		ResetToken: ""},
	}

	rowDummies := make([]model.User, len(expected))
	rows := sqlmock.NewRows([]string{"id", "username","password", "role", "reser_token", "created_at", "updated_at"})
	for i, user := range expected {
		rowDummies[i] = user
		rows.AddRow(user.ID, user.Username,user.Password, user.Role, user.ResetToken, user.CreatedAt, user.UpdatedAt)
	}
	expectedQuery := (`SELECT (.+) FROM "users" WHERE username = ? `)
	suite.mock.ExpectQuery(expectedQuery).WithArgs(expected[0].Username).WillReturnRows(rows)
	repo := NewUserRepository(suite.mockDB)
	actual, err := repo.GetByUsername(expected[0].Username)
	assert.Equal(suite.T(), expected[0], actual)
	assert.NoError(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestUsername_Error() {
	username := "admin"
	expectedQuery := (`SELECT (.+) FROM "users" WHERE username = ?`)
	suite.mock.ExpectQuery(expectedQuery).WithArgs(username).WillReturnError(errors.New("error"))
	repo := NewUserRepository(suite.mockDB)
	actual, err := repo.GetByUsername(username)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.User{}, actual)
}

//GetByUsername(username string) (model.User, error)

//--------------------- TEST UPDATE
//--------------------- TEST DELETE
