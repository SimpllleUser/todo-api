package routes

import (
	"example/todo-api/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, todoController *controller.TodoController) {

	api := r.Group("/api/v1")
	{
		todos := api.Group("/todos")
		{
			todos.GET("", todoController.GetTodos)
			todos.GET("/:id", todoController.GetTodoById)
			todos.GET("/title/:title", todoController.GetTodoByTitle)
			todos.POST("", todoController.CreateTodos)
			todos.PATCH("/:id", todoController.UpdateTodo)
			todos.DELETE("", todoController.DeleteTodo)
		}
	}

}
