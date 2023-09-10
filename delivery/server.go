package delivery

import (
	"fmt"

	"github.com/aisyahenha/golang-les-sekolah-dasar/config"
	"github.com/aisyahenha/golang-les-sekolah-dasar/delivery/controller"
	"github.com/aisyahenha/golang-les-sekolah-dasar/delivery/middleware"
	"github.com/aisyahenha/golang-les-sekolah-dasar/manager"

	// "github.com/aisyahenha/golang-les-sekolah-dasar/repository"
	"github.com/aisyahenha/golang-les-sekolah-dasar/usecase"
	loggerutil "github.com/aisyahenha/golang-les-sekolah-dasar/utils/logger_util"
	"github.com/aisyahenha/golang-les-sekolah-dasar/utils/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	// uc            usecase.UserUseCase
	uc            manager.UseCaseManager
	authService   usecase.AuthUseCase
	engine        *gin.Engine
	host          string
	jwtService    service.JwtService
	loggerService loggerutil.LoggerUtil
}

func (s *Server) setupControllers() {
	//manggil logger
	s.engine.Use(middleware.NewLogMiddleware(s.loggerService).Logger())
	// semua controller di taruh disini
	// controller.NewUserController(s.uc, s.engine)
	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)
	controller.NewUserController(s.uc.UserUseCase(), s.engine, authMiddleware)
	controller.NewCourseController(s.uc.CourseUseCase(), s.engine, authMiddleware)
	controller.NewScheduleController(s.uc.ScheduleUseCase(), s.engine, authMiddleware)
	controller.NewStudentController(s.uc.StudentUseCase(), s.engine, authMiddleware)
	controller.NewTeacherController(s.uc.TeacherUseCase(), s.engine, authMiddleware)
	
	controller.NewAuthController(s.engine, s.authService)
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
	// ur := repository.NewUserRepository(infraManager.Conn())
	// uUc := usecase.NewUserUseCase(ur)
	ur := manager.NewRepoManager(infraManager)
	uUc := manager.NewUseCaseManager(ur)
	engine := gin.Default()
	loggerService := loggerutil.NewLoggerUtil(cfg.FileConfig)

	jwtService := service.NewJwtService(cfg.JwtConfig)
	authUseCase := usecase.NewAuthUseCase(uUc.UserUseCase(), jwtService)
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		uc:            uUc,
		engine:        engine,
		host:          host,
		loggerService: loggerService,
		jwtService:    jwtService,
		authService:   authUseCase,
	}
}
