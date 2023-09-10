package usecase
/*
import (
	"fmt"
	"testing"
	"time"

	// repomock "github.com/aisyahenha/golang-les-sekolah-dasar/_mock_/repo_mock"
	usecasemock "github.com/aisyahenha/golang-les-sekolah-dasar/_mock_/usecase_mock"
	"github.com/aisyahenha/golang-les-sekolah-dasar/config"
	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/aisyahenha/golang-les-sekolah-dasar/utils/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthUseCaseTestSuite struct {
	suite.Suite
	ucMock *usecasemock.UserUseCaseMock
	usecase    AuthUseCase
	jwtService service.JwtService	
	// cfg *config.JwtConfig
}

func (s *AuthUseCaseTestSuite) SetupTest() {
	tmpCfg, err := config.NewConfig()
	fmt.Println("ini errornya: ",err)
	// s.cfg=&tmpCfg.JwtConfig
	s.ucMock = new(usecasemock.UserUseCaseMock)
	s.jwtService = service.NewJwtService(tmpCfg.JwtConfig)
	s.usecase = NewAuthUseCase(s.ucMock, s.jwtService)

}

func TestAuthUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(AuthUseCaseTestSuite))
}

func (suite *AuthUseCaseTestSuite) TestLogin_Success() {
	username, password := "username", "password"
	user := model.User{
		BaseModel: model.BaseModel{
			ID:        "1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username: "username",
		Password: "password",
		Role:     "admin",
	}
	fmt.Println("satu")
	token, err1 := suite.jwtService.CreateAccessToken(user)
	
	fmt.Println("duaaaaaa:",err1)
	suite.ucMock.On("LoginUser", username, password).Return(token, nil)
	fmt.Println("token 1: ",token)
	actual, err := suite.usecase.Login(username, password)
	fmt.Println("actual 1: ", actual)
	fmt.Println("eror token : ", err)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), token, actual)
	fmt.Print("actualllll.......: ", actual)
	fmt.Print("tokennnnnn.......: ", token)

}
*/