package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
	"net/http"
)

func (h *Handler) InitTodoRoutes(router *gin.RouterGroup) {
	todo := router.Group("/todo", h.AuthMiddleware)
	{
		todo.GET("/", h.todoGetAll)
		todo.POST("/", h.todoCreate)
		todo.GET("/:id", h.todoGetOne)
		todo.PUT("/:id", h.todoUpdate)
		todo.DELETE("/:id", h.todoDelete)
	}
}

func (h *Handler) todoGetOne(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := GetUserIdFromCtx(ctx)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusUnauthorized, err)
		return
	}

	todo, err := h.services.ToDo.Get(ctx, id, userId)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusBadRequest, err)
		return
	}

	NewSuccessResponse(ctx, todo)
}

func (h *Handler) todoGetAll(ctx *gin.Context) {
	userId, err := GetUserIdFromCtx(ctx)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusUnauthorized, err)
		return
	}

	todos, err := h.services.ToDo.GetAll(ctx, userId)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusBadRequest, err)
		return
	}

	NewSuccessResponse(ctx, todos)
}

func (h *Handler) todoCreate(ctx *gin.Context) {
	var in domain.CreateTodoInput
	if err := ctx.ShouldBind(&in); err != nil {
		NewValidatorErrorResponse(ctx, err)
		return
	}

	userId, err := GetUserIdFromCtx(ctx)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusUnauthorized, err)
		return
	}

	todo, err := h.services.ToDo.Create(ctx, userId, in)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusBadRequest, err)
		return
	}

	NewSuccessResponse(ctx, todo)
}

func (h *Handler) todoUpdate(ctx *gin.Context) {
	var in domain.UpdateTodoInput
	if err := ctx.ShouldBind(&in); err != nil {
		NewValidatorErrorResponse(ctx, err)
		return
	}

	id := ctx.Param("id")
	userId, err := GetUserIdFromCtx(ctx)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusUnauthorized, err)
		return
	}

	todo, err := h.services.ToDo.Update(ctx, id, userId, in)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusBadRequest, err)
		return
	}

	NewSuccessResponse(ctx, todo)
}

func (h *Handler) todoDelete(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := GetUserIdFromCtx(ctx)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusUnauthorized, err)
		return
	}

	if err := h.services.ToDo.Delete(ctx, id, userId); err != nil {
		NewErrorResponseFromError(ctx, http.StatusBadRequest, err)
		return
	}

	NewSuccessResponse(ctx, nil)
}
