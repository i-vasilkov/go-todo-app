package app

import (
	"github.com/i-vasilkov/go-todo-app/internal/builder"
	delivery "github.com/i-vasilkov/go-todo-app/internal/delivery/http"
	"github.com/i-vasilkov/go-todo-app/internal/server"
	"github.com/i-vasilkov/go-todo-app/internal/service"
	"github.com/i-vasilkov/go-todo-app/pkg/database/mongodb"
	"log"
)

// todo: move to .env
var (
	mongoUri  = "mongodb://mongo:27017"
	mongoUser = "root"
	mongoPass = "root"
	mongoDb   = "todo"
)

func Run() {
	mongoClient, err := mongodb.NewClient(mongoUri, mongoUser, mongoPass)
	if err != nil {
		log.Fatal(err.Error())
	}
	db := mongoClient.Database(mongoDb)

	repBuilder := builder.NewMongoRepositoriesBuilder(db)
	services := service.NewServices(repBuilder.Build())
	handler := delivery.NewHandler(services)

	srv := server.NewServer(handler.Init())
	if err := srv.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
