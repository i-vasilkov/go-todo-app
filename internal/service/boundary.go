package service

import (
	"context"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
)

//go:generate mockgen -source=boundary.go -destination=mocks/mock.go

// -------------- Service boundary ------------------

type AuthServiceI interface {
	SignUp(ctx context.Context, in domain.CreateUserInput) (string, error)
	SignIn(ctx context.Context, in domain.LoginUserInput) (string, error)
	CheckToken(ctx context.Context, token string) (string, error)
}

type TaskServiceI interface {
	Get(ctx context.Context, id, userId string) (domain.Task, error)
	GetAll(ctx context.Context, userId string) ([]domain.Task, error)
	Create(ctx context.Context, userId string, in domain.CreateTaskInput) (domain.Task, error)
	Update(ctx context.Context, id, userId string, in domain.UpdateTaskInput) (domain.Task, error)
	Delete(ctx context.Context, id, userId string) error
}

// -------------- Repository boundary ------------------

type UserRepositoryI interface {
	Create(ctx context.Context, in domain.CreateUserInput) (domain.User, error)
	GetByCredentials(ctx context.Context, in domain.LoginUserInput) (domain.User, error)
	Get(ctx context.Context, id string) (domain.User, error)
}

type TaskRepositoryI interface {
	Get(ctx context.Context, id, userId string) (domain.Task, error)
	GetAll(ctx context.Context, userId string) ([]domain.Task, error)
	Create(ctx context.Context, userId string, in domain.CreateTaskInput) (domain.Task, error)
	Update(ctx context.Context, id, userId string, in domain.UpdateTaskInput) (domain.Task, error)
	Delete(ctx context.Context, id, userId string) error
}
