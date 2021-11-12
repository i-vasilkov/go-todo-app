package app

import (
	"github.com/i-vasilkov/go-todo-app/internal/builder"
	"github.com/i-vasilkov/go-todo-app/internal/config"
	delivery "github.com/i-vasilkov/go-todo-app/internal/delivery/http"
	"github.com/i-vasilkov/go-todo-app/internal/server"
	"github.com/i-vasilkov/go-todo-app/internal/service"
	"github.com/i-vasilkov/go-todo-app/pkg/auth/jwt"
	"github.com/i-vasilkov/go-todo-app/pkg/database/mongodb"
	"github.com/i-vasilkov/go-todo-app/pkg/hash"
	"log"
)

// @title Golang ToDoApp API
// @version 1.0
// @description API Server for ToDoApp

// @host localhost:8000
// @BasePath /api/

// @securityDefinitions.apikey ApiAuth
// @in header
// @name Authorization

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

	deps := service.Dependencies{
		Hasher:     hash.NewSHA1Hasher(cfg.Auth.PwdSalt),
		JwtManager: jwt.NewManager(cfg.Jwt.Ttl, cfg.Jwt.Signature),
	}

	repBuilder := builder.NewMongoRepositoriesBuilder(db)
	services := service.NewServices(repBuilder.Build(), &deps)
	handler := delivery.NewHandler(services)

	srv := server.NewServer(handler.Init(), &cfg)
	if err := srv.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
