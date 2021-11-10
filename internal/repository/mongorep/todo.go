package mongorep

import (
	"context"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type ToDoRepository struct {
	db *mongo.Database
}

func NewMongoToDoRepository(db *mongo.Database) *ToDoRepository {
	return &ToDoRepository{
		db: db,
	}
}

func (t *ToDoRepository) Get(ctx context.Context, id string) (domain.Todo, error) {
	return domain.Todo{}, nil
}

func (t *ToDoRepository) GetAll(ctx context.Context) ([]domain.Todo, error) {
	return nil, nil
}

func (t *ToDoRepository) Create(ctx context.Context, in domain.CreateTodoInput) (domain.Todo, error) {
	return domain.Todo{}, nil
}

func (t *ToDoRepository) Update(ctx context.Context, id string, in domain.UpdateTodoInput) (domain.Todo, error) {
	return domain.Todo{}, nil
}

func (t *ToDoRepository) Delete(ctx context.Context, id string) error {
	return nil
}
