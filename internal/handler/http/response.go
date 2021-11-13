package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Success  bool     `json:"success" example:"false"`
	Messages []string `json:"messages"`
}

type SuccessResponse struct {
	Success bool        `json:"success" example:"true"`
	Data    interface{} `json:"data" extensions:"x-nullable"`
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

	validatorErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		NewErrorResponse(ctx, http.StatusBadRequest, []string{"invalid input body"})
		return
	}

	for _, fieldErr := range validatorErrors {
		messages = append(messages, fmt.Sprintf("invalid '%v' input", fieldErr.Field()))
	}
	NewErrorResponse(ctx, http.StatusBadRequest, messages)
}

func NewSuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, SuccessResponse{true, data})
}
