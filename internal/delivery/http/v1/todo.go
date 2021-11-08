package v1

import "github.com/gin-gonic/gin"

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
	c.Writer.Write([]byte("todoGetOne"))
}

func (h *Handler) todoGetAll(c *gin.Context) {
	c.Writer.Write([]byte("todoGetAll"))
}

func (h *Handler) todoCreate(c *gin.Context) {
	c.Writer.Write([]byte("todoCreate"))
}

func (h *Handler) todoUpdate(c *gin.Context) {
	c.Writer.Write([]byte("todoUpdate"))
}

func (h *Handler) todoDelete(c *gin.Context) {
	c.Writer.Write([]byte("todoDelete"))
}
