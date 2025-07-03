package handler

import (
	model "example/todo-api/internal/models"
	service "example/todo-api/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService *service.TodoService
}

func NewTodoController(tService *service.TodoService) *TodoController {
	return &TodoController{
		todoService: tService,
	}
}

// GetTodos godoc
//
//	@Summary		Get all todos
//	@Description	get todos
//	@Tags			Todos
//	@Produce		json
//	@Success		200	{array}	model.TodoModel
//	@Router			/todos [get]
func (tc *TodoController) GetTodos(c *gin.Context) {
	todos, err := tc.todoService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todos,
	})
}

func (tc *TodoController) CreateTodos(c *gin.Context) {
	var todo model.TodoModel

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := tc.todoService.Create(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": todo,
	})
}

// GetTodoById godoc
//	@Success	200	{object}	map[string]model.TodoModel
//	@Failure	400	{object}	model.HTTPError

func (tc *TodoController) GetTodoById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return

	}

	todo, err := tc.todoService.GetById(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"data": todo,
	})
}

func (tc *TodoController) GetTodoByTitle(c *gin.Context) {
	println(c.Params)
	title := c.Param("title")

	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Title"})
		return

	}

	todo, err := tc.todoService.GetByTitle(title)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"data": todo,
	})
}

func (tc *TodoController) UpdateTodo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	todo, err := tc.todoService.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := tc.todoService.Update(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}

func (tc *TodoController) DeleteTodo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := tc.todoService.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": true,
	})

}
