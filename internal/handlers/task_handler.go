package handler

import (
	model "example/todo-api/internal/models"
	service "example/todo-api/internal/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	userScope *service.UserScopeService
}

func NewTaskController(uService *service.UserScopeService) *TaskController {
	return &TaskController{
		userScope: uService,
	}
}

func (tc *TaskController) getUserId(c *gin.Context) uint {
	return c.GetUint("userId")
}

func (tc *TaskController) ForUser(c *gin.Context) *service.TaskService {
	userId := tc.getUserId(c)
	return tc.userScope.ForUser(userId).Task()
}

// GetTasks godoc
//
//	@Summary		Get all tasks
//	@Description	get tasks
//	@Tags			Tasks
//	@Produce		json
//	@Success		200	{array}		model.TaskModel
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/tasks [get]
//
// @Security BearerAuth
func (tc *TaskController) GetTasks(c *gin.Context) {

	tasks, err := tc.ForUser(c).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})
}

// CreateTasks godoc
//
//	@Summary		Create task
//	@Description	Create task
//	@Tags			Tasks
//	@Param			task	body	model.TaskCreateRequest	true	"Task data"
//	@Produce		json
//	@accept			json
//	@Success		200	{object}	model.TaskModel
//	@Failure		400	{object}	model.HTTPError	"Invalid request"
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/tasks [post]
//
// @Security BearerAuth
func (tc *TaskController) CreateTasks(c *gin.Context) {
	var task model.TaskCreateRequest

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdTask, err := tc.ForUser(c).Create(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("Created Task:", createdTask)

	c.JSON(http.StatusCreated, gin.H{
		"data": createdTask,
	})
}

// GetTaskById godoc
//
//	@Summary		Get task by id
//	@Description	get task by id
//	@Param			id	path	int	true	"Task ID"
//	@Tags			Tasks
//	@Produce		json
//	@Success		302	{object}	model.TaskModel
//	@Failure		400	{object}	model.HTTPError	"Invalid request"
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/tasks/:id [get]
//
// @Security BearerAuth
func (tc *TaskController) GetTaskById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return

	}

	task, err := tc.ForUser(c).GetById(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"data": task,
	})
}

// GetTaskByTitle godoc
//
//	@Summary		Get task by title
//	@Description	get task by title
//	@Param			title	path	string	true	"Task title"
//	@Tags			Tasks
//	@Produce		json
//	@Success		200	{object}	model.TaskModel
//	@Failure		400	{object}	model.HTTPError	"Invalid request"
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/tasks/title/:title [get]
//
// @Security BearerAuth
func (tc *TaskController) GetTaskByTitle(c *gin.Context) {
	println(c.Params)
	title := c.Param("title")

	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Title"})
		return

	}

	task, err := tc.ForUser(c).GetByTitle(title)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"data": task,
	})
}

// UpdateTask godoc
//
//	@Summary		Update task
//	@Description	update task
//	@Tags			Tasks
//	@Param			task	body	model.TaskModel	true	"Task data"
//	@Produce		json
//	@accept			json
//	@Success		200	{object}	model.TaskModel
//	@Failure		400	{object}	model.HTTPError	"Invalid request"
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/tasks [put]
//
// @Security BearerAuth
func (tc *TaskController) UpdateTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	task, err := tc.ForUser(c).GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := tc.ForUser(c).Update(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})
}

// DeleteTask godoc
//
//	@Summary		Delete task by id
//	@Description	delete task by id
//	@Param			id	path	int	true	"Task ID"
//	@Tags			Tasks
//	@Produce		json
//	@Success		200	{object}	model.BooleanResponse
//	@Failure		400	{object}	model.HTTPError	"Invalid request"
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/tasks/:id [delete]
//
// @Security BearerAuth
func (tc *TaskController) DeleteTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := tc.ForUser(c).Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": true,
	})

}

func (tc *TaskController) GetBoardTasks(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Board ID"})
		return
	}

	tasks, err := tc.ForUser(c).GetBoardTasks(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})
}
