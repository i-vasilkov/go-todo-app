package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
	"net/http"
)

func (h *Handler) InitAuthRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.authSignIn)
		auth.POST("/sign-up", h.authSignUp)
	}
}

func (h *Handler) authSignIn(ctx *gin.Context) {
	var in domain.LoginUserInput
	if err := ctx.BindJSON(&in); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Auth.SignIn(ctx, in)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (h *Handler) authSignUp(ctx *gin.Context) {
	var in domain.CreateUserInput
	if err := ctx.BindJSON(&in); err != nil {
		NewErrorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}
	fmt.Println(in)

	token, err := h.services.Auth.SignUp(ctx, in)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
