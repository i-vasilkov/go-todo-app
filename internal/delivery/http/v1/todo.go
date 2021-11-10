package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
	"net/http"
)

func (h *Handler) InitTodoRoutes(router *gin.RouterGroup) {
	todo := router.Group("/todo")
	{
		todo.GET("/", h.todoGetAll)
		todo.POST("/", h.todoCreate)
		todo.GET("/:id", h.todoGetOne)
		todo.PUT("/:id", h.todoUpdate)
		todo.DELETE("/:id", h.todoDelete)
	}
}

func (h *Handler) todoGetOne(c *gin.Context) {
	id := c.Param("id")

	todo, err := h.services.ToDo.Get(c, id)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *Handler) todoGetAll(c *gin.Context) {
	todos, err := h.services.ToDo.GetAll(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (h *Handler) todoCreate(c *gin.Context) {
	var in domain.CreateTodoInput
	if err := c.BindJSON(&in); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	todo, err := h.services.ToDo.Create(c, in)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *Handler) todoUpdate(c *gin.Context) {
	var in domain.UpdateTodoInput
	if err := c.BindJSON(&in); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")

	todo, err := h.services.ToDo.Update(c, id, in)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *Handler) todoDelete(c *gin.Context) {
	id := c.Param("id")
	if err := h.services.ToDo.Delete(c, id); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
