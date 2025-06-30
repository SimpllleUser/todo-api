package routes

import (
	"example/todo-api/controller"
	"example/todo-api/middlewares"
	"example/todo-api/model"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine,
	todoController *controller.TodoController,
	userController *controller.UserController,
	authController *controller.AuthController,
) {

	userService := model.NewUserService(model.DB)

	api := r.Group("/api/v1")
	{
		todos := api.Group("/todos")
		todos.Use(middlewares.CheckAuth(userService))
		{
			todos.GET("", todoController.GetTodos)
			todos.GET("/:id", todoController.GetTodoById)
			todos.GET("/title/:title", todoController.GetTodoByTitle)
			todos.POST("", todoController.CreateTodos)
			todos.PATCH("/:id", todoController.UpdateTodo)
			todos.DELETE("/:id", todoController.DeleteTodo)
		}

		auth := api.Group("/auth")
		{
			auth.POST("/login", authController.Login)

			/// TODO fix response on success
			auth.POST("/registration", userController.CreateUser)
		}
	}

}
