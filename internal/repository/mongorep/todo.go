package mongorep

import (
	"context"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ToDoRepository struct {
	db *mongo.Database
}

func NewMongoToDoRepository(db *mongo.Database) *ToDoRepository {
	return &ToDoRepository{
		db: db,
	}
}

func (rep *ToDoRepository) Get(ctx context.Context, id string) (domain.Todo, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Todo{}, err
	}

	var todo domain.Todo
	err = rep.db.Collection("todos").
		FindOne(ctx, bson.M{"_id": objId}).
		Decode(&todo)

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (rep *ToDoRepository) GetAll(ctx context.Context) ([]domain.Todo, error) {
	cursor, err := rep.db.Collection("todos").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var todos []domain.Todo
	if err := cursor.All(ctx, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (rep *ToDoRepository) Create(ctx context.Context, in domain.CreateTodoInput) (domain.Todo, error) {
	result, err := rep.db.Collection("todos").
		InsertOne(ctx, bson.M{
			"name":       in.Name,
			"created_at": time.Now().Format(time.RFC3339),
			"updated_at": time.Now().Format(time.RFC3339),
		})

	if err != nil {
		return domain.Todo{}, err
	}

	objId := result.InsertedID.(primitive.ObjectID)

	var todo domain.Todo
	err = rep.db.Collection("todos").
		FindOne(ctx, bson.M{"_id": objId}).
		Decode(&todo)

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (rep *ToDoRepository) Update(ctx context.Context, id string, in domain.UpdateTodoInput) (domain.Todo, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Todo{}, err
	}

	update := bson.M{"$set": bson.M{"name": in.Name, "updated_at": time.Now().Format(time.RFC3339)}}
	_, err = rep.db.Collection("todos").UpdateOne(ctx, bson.M{"_id": objId}, update)

	var todo domain.Todo
	err = rep.db.Collection("todos").
		FindOne(ctx, bson.M{"_id": objId}).
		Decode(&todo)

	if err != nil {
		return todo, err
	}

	return todo, nil

	return todo, nil
}

func (rep *ToDoRepository) Delete(ctx context.Context, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = rep.db.Collection("todos").DeleteOne(ctx, bson.M{"_id": objId})
	return err
}
