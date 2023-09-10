package manager

import "github.com/aisyahenha/golang-les-sekolah-dasar/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	CourseRepo() repository.CourseRepository
	ScheduleRepo() repository.ScheduleRepository
	StudentRepo() repository.StudentRepository
	TeacherRepo() repository.TeacherRepository
}

type repoManager struct {
	infra InfraManager
}

// CourseRepo implements RepoManager.
func (r *repoManager) CourseRepo() repository.CourseRepository {
	return repository.NewCourseRepository(r.infra.Conn())
}

// ScheduleRepo implements RepoManager.
func (r *repoManager) ScheduleRepo() repository.ScheduleRepository {
	return repository.NewScheduleRepository(r.infra.Conn())
}

// StudentRepo implements RepoManager.
func (r *repoManager) StudentRepo() repository.StudentRepository {
	return repository.NewStudentRepository(r.infra.Conn())
}

// TeacherRepo implements RepoManager.
func (r *repoManager) TeacherRepo() repository.TeacherRepository {
	return repository.NewTeacherRepository(r.infra.Conn())
}

// UserRepo implements RepoManager.
func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
