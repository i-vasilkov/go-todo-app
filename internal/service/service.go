package service

import (
	"github.com/i-vasilkov/go-todo-app/pkg/auth/jwt"
	"github.com/i-vasilkov/go-todo-app/pkg/hash"
)

type Services struct {
	Task *TaskService
	Auth *AuthService
}

type Repositories struct {
	Task TaskRepositoryI
	User UserRepositoryI
}

type Dependencies struct {
	Hasher     hash.Hasher
	JwtManager *jwt.Manager
}

func NewServices(rep *Repositories, deps *Dependencies) *Services {
	return &Services{
		Task: NewTaskService(rep.Task),
		Auth: NewAuthService(rep.User, deps.Hasher, deps.JwtManager),
	}
}
