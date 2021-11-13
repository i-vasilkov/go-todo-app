package service

import (
	"github.com/i-vasilkov/go-todo-app/pkg/auth/jwt"
	"github.com/i-vasilkov/go-todo-app/pkg/hash"
)

type Repositories struct {
	Task TaskRepositoryI
	User UserRepositoryI
}

type Dependencies struct {
	Hasher     hash.Hasher
	JwtManager *jwt.Manager
}

type Services struct {
	Auth AuthServiceI
	Task TaskServiceI
}
