package app

import (
	"github.com/i-vasilkov/go-todo-app/internal/builder"
	"github.com/i-vasilkov/go-todo-app/internal/config"
	delivery "github.com/i-vasilkov/go-todo-app/internal/delivery/http"
	"github.com/i-vasilkov/go-todo-app/internal/server"
	"github.com/i-vasilkov/go-todo-app/internal/service"
	"github.com/i-vasilkov/go-todo-app/pkg/database/mongodb"
	"log"
)

func Run(cfgPath, envPath string) {
	cfg, err := config.Init(cfgPath, envPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	mongoClient, err := mongodb.NewClient(mongodb.Connection{
		Uri:  cfg.Mongo.GetURI(),
		User: cfg.Mongo.UserName,
		Pass: cfg.Mongo.Password,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	db := mongoClient.Database(cfg.Mongo.DbName)

	repBuilder := builder.NewMongoRepositoriesBuilder(db)
	services := service.NewServices(repBuilder.Build())
	handler := delivery.NewHandler(services)

	srv := server.NewServer(handler.Init(), cfg)
	if err := srv.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
