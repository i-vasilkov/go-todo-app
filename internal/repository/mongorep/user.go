package mongorep

import (
	"context"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepository struct {
	db *mongo.Database
}

func NewMongoUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (rep *UserRepository) Create(ctx context.Context, in domain.CreateUserInput) (domain.User, error) {
	var user domain.User

	result, err := rep.db.Collection("users").
		InsertOne(ctx, bson.M{
			"password":   in.Password,
			"login":      in.Login,
			"created_at": time.Now().Format(time.RFC3339),
		})
	if err != nil {
		return user, err
	}

	objId := result.InsertedID.(primitive.ObjectID)
	err = rep.db.Collection("users").FindOne(ctx, bson.M{"_id": objId}).Decode(&user)

	return user, err
}

func (rep *UserRepository) GetByCredentials(ctx context.Context, in domain.LoginUserInput) (domain.User, error) {
	var user domain.User

	err := rep.db.Collection("users").
		FindOne(ctx, bson.M{
			"login":    in.Login,
			"password": in.Password,
		}).
		Decode(&user)

	return user, err
}

func (rep *UserRepository) Get(ctx context.Context, id string) (domain.User, error) {
	var user domain.User

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = rep.db.Collection("users").FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	return user, err
}
