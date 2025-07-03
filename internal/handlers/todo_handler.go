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
//	@Success		200	{array}		model.TodoModel
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/todos [get]
//
// @Security BearerAuth
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

// CreateTodos godoc
//
//	@Summary		Create todo
//	@Description	Create todo
//	@Tags			Todos
//	@Param			todo	body	model.TodoCreateRequest	true	"Todo data"
//	@Produce		json
//	@accept			json
//	@Success		200	{object}	model.TodoModel
//	@Failure		400	{object}	model.HTTPError	"Invalid request"
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/todos [post]
//
// @Security BearerAuth
func (tc *TodoController) CreateTodos(c *gin.Context) {
	var todo model.TodoCreateRequest

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
//
//	@Summary		Get todo by id
//	@Description	get todo by id
//	@Param			id	path	int	true	"Todo ID"
//	@Tags			Todos
//	@Produce		json
//	@Success		302	{object}	model.TodoModel
//	@Failure		400	{object}	model.HTTPError	"Invalid request"
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/todos/:id [get]
//
// @Security BearerAuth
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

// GetTodoByTitle godoc
//
//	@Summary		Get todo by title
//	@Description	get todo by title
//	@Param			title	path	string	true	"Todo title"
//	@Tags			Todos
//	@Produce		json
//	@Success		200	{object}	model.TodoModel
//	@Failure		400	{object}	model.HTTPError	"Invalid request"
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/todos/title/:title [get]
//
// @Security BearerAuth
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

// UpdateTodo godoc
//
//	@Summary		Update todo
//	@Description	update todo
//	@Tags			Todos
//	@Param			todo	body	model.TodoModel	true	"Todo data"
//	@Produce		json
//	@accept			json
//	@Success		200	{object}	model.TodoModel
//	@Failure		400	{object}	model.HTTPError	"Invalid request"
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/todos [put]
//
// @Security BearerAuth
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

// DeleteTodo godoc
//
//	@Summary		Delete todo by id
//	@Description	delete todo by id
//	@Param			id	path	int	true	"Todo ID"
//	@Tags			Todos
//	@Produce		json
//	@Success		200	{object}	model.BooleanResponse
//	@Failure		400	{object}	model.HTTPError	"Invalid request"
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/todos/:id [delete]
//
// @Security BearerAuth
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
