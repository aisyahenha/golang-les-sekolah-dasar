package manager

import "github.com/aisyahenha/golang-les-sekolah-dasar/usecase"

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
	CourseUseCase() usecase.CourseUseCase
	ScheduleUseCase() usecase.ScheduleUseCase
	StudentUseCase() usecase.StudentUseCase
	TeacherUseCase() usecase.TeacherUseCase
}

type useCaseManager struct {
	repo RepoManager
}

// CouseUseCase implements UseCaseManager.
func (u *useCaseManager) CourseUseCase() usecase.CourseUseCase {
	return usecase.NewCourseUseCase(u.repo.CourseRepo())
}

// ScheduleUsecase implements UseCaseManager.
func (u *useCaseManager) ScheduleUseCase() usecase.ScheduleUseCase {
	return usecase.NewScheduleUseCase(u.repo.ScheduleRepo())
}

// StudentUseCase implements UseCaseManager.
func (u *useCaseManager) StudentUseCase() usecase.StudentUseCase {
	return usecase.NewStudentUseCase(u.repo.StudentRepo())
}

// TeacherUsecase implements UseCaseManager.
func (u *useCaseManager) TeacherUseCase() usecase.TeacherUseCase {
	return usecase.NewTeacherUseCase(u.repo.TeacherRepo())
}

// UserUseCase implements UseCaseManager.
func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.UserRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
