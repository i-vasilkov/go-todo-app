package http

import (
	"github.com/gin-gonic/gin"
	"github.com/i-vasilkov/go-todo-app/internal/service"
	"net/http"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/i-vasilkov/go-todo-app/docs"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			h.InitTaskRoutes(v1)
			h.InitAuthRoutes(v1)
		}
	}

	return router
}
