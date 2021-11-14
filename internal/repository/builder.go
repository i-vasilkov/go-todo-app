package repository

import (
	"github.com/i-vasilkov/go-todo-app/internal/repository/mongorep"
	"github.com/i-vasilkov/go-todo-app/internal/repository/postgresrep"
	"github.com/i-vasilkov/go-todo-app/internal/service"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

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

type PostgresRepositoriesBuilder struct {
	db *sqlx.DB
}

func NewPostgresRepositoriesBuilder(db *sqlx.DB) *PostgresRepositoriesBuilder {
	return &PostgresRepositoriesBuilder{db: db}
}

func (rb *PostgresRepositoriesBuilder) Build() *service.Repositories {
	return &service.Repositories{
		Task: postgresrep.NewPostgresTaskRepository(rb.db),
		User: postgresrep.NewPostgresUserRepository(rb.db),
	}
}
