package controller

import (
	"example/todo-api/model"
	db "example/todo-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService *model.TodoService
}

func NewTodoController(tService *model.TodoService) *TodoController {
	return &TodoController{
		todoService: tService,
	}
}

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
	var todo db.TodoModel

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
