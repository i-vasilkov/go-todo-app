package app

import (
	"context"
	"github.com/i-vasilkov/go-todo-app/internal/config"
	delivery "github.com/i-vasilkov/go-todo-app/internal/handler/http"
	"github.com/i-vasilkov/go-todo-app/internal/repository"
	"github.com/i-vasilkov/go-todo-app/internal/server"
	"github.com/i-vasilkov/go-todo-app/internal/service"
	"github.com/i-vasilkov/go-todo-app/pkg/auth/jwt"
	"github.com/i-vasilkov/go-todo-app/pkg/database/mongodb"
	"github.com/i-vasilkov/go-todo-app/pkg/hash"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	repBuilder := repository.NewMongoRepositoriesBuilder(db)
	serviceBuilder := service.NewAppServiceBuilder(&deps, repBuilder.Build())
	handler := delivery.NewHandler(serviceBuilder.Build())

	srv := server.NewServer(handler.Init(), &cfg)
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	timeout := time.Second * 5
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Fatal(err.Error())
	}

	if err := mongoClient.Disconnect(context.Background()); err != nil {
		log.Fatal(err.Error())
	}
}
