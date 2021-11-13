package http

import (
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

// @Summary Sign In
// @Description Login user by credentials
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body domain.LoginUserInput true "SignIn Input"
// @Success 200 {object} SuccessResponse{data=string}
// @Failure 400,422,500 {object} ErrorResponse
// @Router /auth/sign-in [post]
func (h *Handler) authSignIn(ctx *gin.Context) {
	var in domain.LoginUserInput
	if err := ctx.BindJSON(&in); err != nil {
		NewValidatorErrorResponse(ctx, err)
		//NewErrorResponseFromError(ctx, http.StatusBadRequest, errors.New("invalid input body"))
		return
	}

	token, err := h.services.Auth.SignIn(ctx.Request.Context(), in)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusInternalServerError, err)
		return
	}

	NewSuccessResponse(ctx, token)
}

// @Summary Sign Up
// @Description Registration user by credentials
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body domain.CreateUserInput true "SignUp Input"
// @Success 200 {object} SuccessResponse{data=string}
// @Failure 400,422,500 {object} ErrorResponse
// @Router /auth/sign-up [post]
func (h *Handler) authSignUp(ctx *gin.Context) {
	var in domain.CreateUserInput
	if err := ctx.BindJSON(&in); err != nil {
		NewValidatorErrorResponse(ctx, err)
		return
	}

	token, err := h.services.Auth.SignUp(ctx.Request.Context(), in)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusInternalServerError, err)
		return
	}

	NewSuccessResponse(ctx, token)
}
