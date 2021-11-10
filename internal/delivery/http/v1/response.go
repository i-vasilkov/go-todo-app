package v1

import (
	"github.com/gin-gonic/gin"
	"log"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(ctx *gin.Context, code int, msg string) {
	log.Println(msg)
	ctx.AbortWithStatusJSON(code, ErrorResponse{msg})
}