package http

import (
	"context"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
)

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
