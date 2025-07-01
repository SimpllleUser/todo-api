package routes

import (
	handler "example/todo-api/internal/handlers"
	"example/todo-api/internal/middlewares"
	model "example/todo-api/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine,
	todoHandler *handler.TodoController,
	userHandler *handler.UserController,
	authHandler *handler.AuthController,
	userService *model.UserService,
) {
	api := r.Group("/api/v1")
	{
		todos := api.Group("/todos")
		todos.Use(middlewares.CheckAuth(userService))
		{
			todos.GET("", todoHandler.GetTodos)
			todos.GET("/:id", todoHandler.GetTodoById)
			todos.GET("/title/:title", todoHandler.GetTodoByTitle)
			todos.POST("", todoHandler.CreateTodos)
			todos.PATCH("/:id", todoHandler.UpdateTodo)
			todos.DELETE("/:id", todoHandler.DeleteTodo)
		}

		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)

			/// TODO fix response on success
			auth.POST("/registration", userHandler.CreateUser)
		}
	}

}
