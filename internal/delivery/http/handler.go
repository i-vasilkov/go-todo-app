package http

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/i-vasilkov/go-todo-app/internal/delivery/http/v1"
	"github.com/i-vasilkov/go-todo-app/internal/service"
	"net/http"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() http.Handler {
	router := gin.Default()

	router.GET("/", h.AppMsg)

	handlerV1 := v1.NewHandler(h.services)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}

	return router
}

func (h *Handler) AppMsg(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Application is working")
}
