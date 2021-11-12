package builder

import (
	"github.com/i-vasilkov/go-todo-app/internal/repository/mongorep"
	"github.com/i-vasilkov/go-todo-app/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoriesBuilder interface {
	Build() service.Repositories
}

type MongoRepositoriesBuilder struct {
	db *mongo.Database
}

func NewMongoRepositoriesBuilder(db *mongo.Database) *MongoRepositoriesBuilder {
	return &MongoRepositoriesBuilder{db: db}
}

func (rb *MongoRepositoriesBuilder) Build() *service.Repositories {
	return &service.Repositories{
		Task: mongorep.NewMongoTaskRepository(rb.db),
		User: mongorep.NewMongoUserRepository(rb.db),
	}
}
