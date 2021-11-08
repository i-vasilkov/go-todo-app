package service

import (
	"context"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
)

type ToDoService struct {
}

func NewToDo() *ToDoService {
	return &ToDoService{}
}

func (t *ToDoService) Get(ctx context.Context, id string) (domain.Todo, error) {
	return domain.Todo{}, nil
}

func (t *ToDoService) GetAll(ctx context.Context) ([]domain.Todo, error) {
	return nil, nil
}

func (t *ToDoService) Create(ctx context.Context, in domain.CreateTodoInput) (domain.Todo, error) {
	return domain.Todo{}, nil
}

func (t *ToDoService) Update(ctx context.Context, in domain.UpdateTodoInput) (domain.Todo, error) {
	return domain.Todo{}, nil
}

func (t *ToDoService) Delete(ctx context.Context, id string) error {
	return nil
}
