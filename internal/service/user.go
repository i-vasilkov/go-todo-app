package service

import (
	"context"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
)

type UserRepositoryI interface {
	Create(context.Context, domain.CreateUserInput) (domain.User, error)
	GetByCredentials(context.Context, domain.LoginUserInput) (domain.User, error)
	Get(context.Context, string) (domain.User, error)
}
