package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Success  bool     `json:"success"`
	Messages []string `json:"messages"`
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func NewErrorResponse(ctx *gin.Context, code int, messages []string) {
	log.Println(messages)
	ctx.AbortWithStatusJSON(code, ErrorResponse{false, messages})
}

func NewErrorResponseFromError(ctx *gin.Context, code int, err error) {
	NewErrorResponse(ctx, code, []string{err.Error()})
}

func NewSuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, SuccessResponse{true, data})
}
