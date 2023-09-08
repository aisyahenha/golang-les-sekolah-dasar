package delivery

import (
	"fmt"

	"github.com/aisyahenha/golang-les-sekolah-dasar/config"
	"github.com/aisyahenha/golang-les-sekolah-dasar/delivery/controller"
	"github.com/aisyahenha/golang-les-sekolah-dasar/manager"
	"github.com/aisyahenha/golang-les-sekolah-dasar/repository"
	"github.com/aisyahenha/golang-les-sekolah-dasar/usecase"
	"github.com/gin-gonic/gin"
)

type Server struct {
	uc     usecase.UserUseCase
	engine *gin.Engine
	host   string
}

func (s *Server) setupControllers() {
	// semua controller di taruh disini
	controller.NewUserController(s.uc, s.engine)
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running %s", err.Error()))
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		panic(err)
	}
	// repo
	ur := repository.NewUserRepository(infraManager.Conn())
	uUc := usecase.NewUserUseCase(ur)
	engine := gin.Default()
	// localhost:8888/api/v1/users
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		uc:     uUc,
		engine: engine,
		host:   host,
	}
}
