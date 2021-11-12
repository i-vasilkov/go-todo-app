package service

import (
	"context"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
)

type ToDoService struct {
	rep ToDoRepositoryI
}

type ToDoRepositoryI interface {
	Get(ctx context.Context, id, userId string) (domain.Todo, error)
	GetAll(ctx context.Context, userId string) ([]domain.Todo, error)
	Create(ctx context.Context, userId string, in domain.CreateTodoInput) (domain.Todo, error)
	Update(ctx context.Context, id, userId string, in domain.UpdateTodoInput) (domain.Todo, error)
	Delete(ctx context.Context, id, userId string) error
}

func NewToDoService(rep ToDoRepositoryI) *ToDoService {
	return &ToDoService{
		rep: rep,
	}
}

func (t *ToDoService) Get(ctx context.Context, id, userId string) (domain.Todo, error) {
	return t.rep.Get(ctx, id, userId)
}

func (t *ToDoService) GetAll(ctx context.Context, userId string) ([]domain.Todo, error) {
	return t.rep.GetAll(ctx, userId)
}

func (t *ToDoService) Create(ctx context.Context, userId string, in domain.CreateTodoInput) (domain.Todo, error) {
	return t.rep.Create(ctx, userId, in)
}

func (t *ToDoService) Update(ctx context.Context, id, userId string, in domain.UpdateTodoInput) (domain.Todo, error) {
	return t.rep.Update(ctx, id, userId, in)
}

func (t *ToDoService) Delete(ctx context.Context, id, userId string) error {
	return t.rep.Delete(ctx, id, userId)
}
