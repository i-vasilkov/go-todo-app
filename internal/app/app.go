package app

import (
	delivery "github.com/i-vasilkov/go-todo-app/internal/delivery/http"
	"github.com/i-vasilkov/go-todo-app/internal/server"
	"github.com/i-vasilkov/go-todo-app/internal/service"
	"log"
)

func Run() {
	services := service.NewServices()
	handler := delivery.NewHandler(services)

	srv := server.NewServer(handler.Init())
	if err := srv.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
