package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init() http.Handler {
	router := gin.Default()

	router.GET("/", h.AppMsg)

	return router
}

func (h *Handler) AppMsg(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Application is working")
}
