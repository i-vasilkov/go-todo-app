package service

import (
	"context"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
)

type TaskService struct {
	rep TaskRepositoryI
}

func NewTaskService(rep TaskRepositoryI) *TaskService {
	return &TaskService{
		rep: rep,
	}
}

func (t *TaskService) Get(ctx context.Context, id, userId string) (domain.Task, error) {
	return t.rep.Get(ctx, id, userId)
}

func (t *TaskService) GetAll(ctx context.Context, userId string) ([]domain.Task, error) {
	return t.rep.GetAll(ctx, userId)
}

func (t *TaskService) Create(ctx context.Context, userId string, in domain.CreateTaskInput) (domain.Task, error) {
	return t.rep.Create(ctx, userId, in)
}

func (t *TaskService) Update(ctx context.Context, id, userId string, in domain.UpdateTaskInput) (domain.Task, error) {
	return t.rep.Update(ctx, id, userId, in)
}

func (t *TaskService) Delete(ctx context.Context, id, userId string) error {
	return t.rep.Delete(ctx, id, userId)
}
