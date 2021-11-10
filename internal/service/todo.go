package service

import (
	"context"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
)

type ToDoService struct {
	rep ToDoRepositoryI
}

type ToDoRepositoryI interface {
	Get(ctx context.Context, id string) (domain.Todo, error)
	GetAll(ctx context.Context) ([]domain.Todo, error)
	Create(ctx context.Context, in domain.CreateTodoInput) (domain.Todo, error)
	Update(ctx context.Context, id string, in domain.UpdateTodoInput) (domain.Todo, error)
	Delete(ctx context.Context, id string) error
}

func NewToDo(rep ToDoRepositoryI) *ToDoService {
	return &ToDoService{
		rep: rep,
	}
}

func (t *ToDoService) Get(ctx context.Context, id string) (domain.Todo, error) {
	return t.rep.Get(ctx, id)
}

func (t *ToDoService) GetAll(ctx context.Context) ([]domain.Todo, error) {
	return t.rep.GetAll(ctx)
}

func (t *ToDoService) Create(ctx context.Context, in domain.CreateTodoInput) (domain.Todo, error) {
	return t.rep.Create(ctx, in)
}

func (t *ToDoService) Update(ctx context.Context, id string, in domain.UpdateTodoInput) (domain.Todo, error) {
	return t.rep.Update(ctx, id, in)
}

func (t *ToDoService) Delete(ctx context.Context, id string) error {
	return t.rep.Delete(ctx, id)
}
