package http

import (
	"github.com/gin-gonic/gin"
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

	return router
}

func (h *Handler) AppMsg(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Application is working")
}
