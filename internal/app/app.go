package app

import (
	delivery "github.com/i-vasilkov/go-todo-app/internal/delivery/http"
	"github.com/i-vasilkov/go-todo-app/internal/server"
	"log"
)

func Run() {
	handler := delivery.NewHandler()

	srv := server.NewServer(handler)
	if err := srv.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
