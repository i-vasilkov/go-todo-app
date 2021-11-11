package service

import (
	"github.com/i-vasilkov/go-todo-app/pkg/auth/jwt"
	"github.com/i-vasilkov/go-todo-app/pkg/hash"
)

type Services struct {
	ToDo *ToDoService
	Auth *AuthService
}

type Repositories struct {
	ToDo ToDoRepositoryI
	User UserRepositoryI
}

type Dependencies struct {
	Hasher     hash.Hasher
	JwtManager *jwt.Manager
}

func NewServices(rep Repositories, deps Dependencies) *Services {
	return &Services{
		ToDo: NewToDoService(rep.ToDo),
		Auth: NewAuthService(rep.User, deps.Hasher, deps.JwtManager),
	}
}
