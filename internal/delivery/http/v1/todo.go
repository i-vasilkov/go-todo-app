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

// @Summary Getting one todo
// @Description Get one todo by id
// @Security ApiAuth
// @Tags Todo
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse{data=domain.Todo}
// @Failure 400,422,500 {object} ErrorResponse
// @Router /todo/{id} [get]
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

// @Summary Getting todos
// @Description Get user todos
// @Security ApiAuth
// @Tags Todo
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse{data=[]domain.Todo}
// @Failure 400,422,500 {object} ErrorResponse
// @Router /todo [get]
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

// @Summary Creating todo
// @Description Create todo by input data
// @Security ApiAuth
// @Tags Todo
// @Accept json
// @Produce json
// @Param input body domain.CreateTodoInput true "input data"
// @Success 200 {object} SuccessResponse{data=domain.Todo}
// @Failure 400,422,500 {object} ErrorResponse
// @Router /todo [post]
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

// @Summary Updating todo
// @Description Update todo by input data
// @Security ApiAuth
// @Tags Todo
// @Accept json
// @Produce json
// @Param input body domain.UpdateTodoInput true "input data"
// @Success 200 {object} SuccessResponse{data=domain.Todo}
// @Failure 400,422,500 {object} ErrorResponse
// @Router /todo/{id} [put]
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

// @Summary Deleting todo
// @Description Delete todo by id
// @Security ApiAuth
// @Tags Todo
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse{data=object}
// @Failure 400,422,500 {object} ErrorResponse
// @Router /todo/{id} [delete]
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
