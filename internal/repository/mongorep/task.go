package mongorep

import (
	"context"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type TaskRepository struct {
	db *mongo.Database
}

func NewMongoTaskRepository(db *mongo.Database) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (rep *TaskRepository) Get(ctx context.Context, id, userId string) (domain.Task, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{}, err
	}

	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return domain.Task{}, err
	}

	var task domain.Task
	err = rep.db.Collection(tasksCollection).
		FindOne(ctx, bson.M{"_id": objId, "user_id": userObjId}).
		Decode(&task)

	if err != nil {
		return task, err
	}

	return task, nil
}

func (rep *TaskRepository) GetAll(ctx context.Context, userId string) ([]domain.Task, error) {
	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	cursor, err := rep.db.Collection(tasksCollection).Find(ctx, bson.M{"user_id": userObjId})
	if err != nil {
		return nil, err
	}

	tasks := make([]domain.Task, 0)
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (rep *TaskRepository) Create(ctx context.Context, userId string, in domain.CreateTaskInput) (domain.Task, error) {
	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return domain.Task{}, err
	}

	result, err := rep.db.Collection(tasksCollection).
		InsertOne(ctx, bson.M{
			"name":       in.Name,
			"created_at": time.Now().Format(time.RFC3339),
			"updated_at": time.Now().Format(time.RFC3339),
			"user_id":    userObjId,
		})

	if err != nil {
		return domain.Task{}, err
	}

	objId := result.InsertedID.(primitive.ObjectID)

	var task domain.Task
	err = rep.db.Collection(tasksCollection).
		FindOne(ctx, bson.M{"_id": objId}).
		Decode(&task)

	if err != nil {
		return task, err
	}

	return task, nil
}

func (rep *TaskRepository) Update(ctx context.Context, id, userId string, in domain.UpdateTaskInput) (domain.Task, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{}, err
	}

	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return domain.Task{}, err
	}

	update := bson.M{"$set": bson.M{"name": in.Name, "updated_at": time.Now().Format(time.RFC3339)}}
	_, err = rep.db.Collection(tasksCollection).UpdateOne(ctx, bson.M{"_id": objId, "user_id": userObjId}, update)

	var tasks domain.Task
	err = rep.db.Collection(tasksCollection).
		FindOne(ctx, bson.M{"_id": objId}).
		Decode(&tasks)

	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (rep *TaskRepository) Delete(ctx context.Context, id, userId string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	_, err = rep.db.Collection(tasksCollection).DeleteOne(ctx, bson.M{"_id": objId, "user_id": userObjId})
	return err
}
