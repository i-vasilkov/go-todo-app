package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func NewValidatorErrorResponse(ctx *gin.Context, err error) {
	var messages []string
	for _, fieldErr := range err.(validator.ValidationErrors) {
		messages = append(messages, fieldErr.Error())
	}
	NewErrorResponse(ctx, http.StatusBadRequest, messages)
}

func NewSuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, SuccessResponse{true, data})
}
