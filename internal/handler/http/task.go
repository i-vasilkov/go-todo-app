package http

import (
	"github.com/gin-gonic/gin"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
	"net/http"
)

func (h *Handler) InitTaskRoutes(router *gin.RouterGroup) {
	task := router.Group("/task", h.AuthMiddleware)
	{
		task.GET("/", h.taskGetAll)
		task.POST("/", h.taskCreate)
		task.GET("/:id", h.taskGetOne)
		task.PUT("/:id", h.taskUpdate)
		task.DELETE("/:id", h.taskDelete)
	}
}

// @Summary Getting one task
// @Description Get one task by id
// @Security ApiAuth
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse{data=domain.Task}
// @Failure 400,422,500 {object} ErrorResponse
// @Router /task/{id} [get]
func (h *Handler) taskGetOne(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := GetUserIdFromCtx(ctx)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusUnauthorized, err)
		return
	}

	task, err := h.services.Task.Get(ctx, id, userId)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusBadRequest, err)
		return
	}

	NewSuccessResponse(ctx, task)
}

// @Summary Getting tasks
// @Description Get user tasks
// @Security ApiAuth
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse{data=[]domain.Task}
// @Failure 400,422,500 {object} ErrorResponse
// @Router /task [get]
func (h *Handler) taskGetAll(ctx *gin.Context) {
	userId, err := GetUserIdFromCtx(ctx)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusUnauthorized, err)
		return
	}

	tasks, err := h.services.Task.GetAll(ctx, userId)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusBadRequest, err)
		return
	}

	NewSuccessResponse(ctx, tasks)
}

// @Summary Creating task
// @Description Create task by input data
// @Security ApiAuth
// @Tags Task
// @Accept json
// @Produce json
// @Param input body domain.CreateTaskInput true "input data"
// @Success 200 {object} SuccessResponse{data=domain.Task}
// @Failure 400,422,500 {object} ErrorResponse
// @Router /task [post]
func (h *Handler) taskCreate(ctx *gin.Context) {
	var in domain.CreateTaskInput
	if err := ctx.ShouldBind(&in); err != nil {
		NewValidatorErrorResponse(ctx, err)
		return
	}

	userId, err := GetUserIdFromCtx(ctx)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusUnauthorized, err)
		return
	}

	task, err := h.services.Task.Create(ctx, userId, in)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusBadRequest, err)
		return
	}

	NewSuccessResponse(ctx, task)
}

// @Summary Updating task
// @Description Update task by input data
// @Security ApiAuth
// @Tags Task
// @Accept json
// @Produce json
// @Param input body domain.UpdateTaskInput true "input data"
// @Success 200 {object} SuccessResponse{data=domain.Task}
// @Failure 400,422,500 {object} ErrorResponse
// @Router /task/{id} [put]
func (h *Handler) taskUpdate(ctx *gin.Context) {
	var in domain.UpdateTaskInput
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

	task, err := h.services.Task.Update(ctx, id, userId, in)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusBadRequest, err)
		return
	}

	NewSuccessResponse(ctx, task)
}

// @Summary Deleting task
// @Description Delete task by id
// @Security ApiAuth
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse{data=object}
// @Failure 400,422,500 {object} ErrorResponse
// @Router /task/{id} [delete]
func (h *Handler) taskDelete(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := GetUserIdFromCtx(ctx)
	if err != nil {
		NewErrorResponseFromError(ctx, http.StatusUnauthorized, err)
		return
	}

	if err := h.services.Task.Delete(ctx, id, userId); err != nil {
		NewErrorResponseFromError(ctx, http.StatusBadRequest, err)
		return
	}

	NewSuccessResponse(ctx, nil)
}
