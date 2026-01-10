package routes

import (
	_ "example/todo-api/docs"
	handler "example/todo-api/internal/handlers"
	"example/todo-api/internal/middlewares"
	service "example/todo-api/internal/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			My API
//	@version		1.0
//	@description	This is a sample REST API server.
//	@host			localhost:8080
//	@BasePath		/api/v1

func SetupRoutes(r *gin.Engine,
	todoHandler *handler.TodoController,
	userHandler *handler.UserController,
	authHandler *handler.AuthController,
	boardHandler *handler.BoardController,
	userService *service.UserService,
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

		boards := api.Group("/boards")
		boards.Use(middlewares.CheckAuth(userService))
		{
			boards.GET("", boardHandler.GetBoards)
			boards.GET("/:id", boardHandler.GetBoardById)
			boards.GET("/title/:title", boardHandler.GetBoardByTitle)
			boards.POST("", boardHandler.CreateBoards)
			boards.PATCH("/:id", boardHandler.UpdateBoard)
			boards.DELETE("/:id", boardHandler.DeleteBoard)
		}

		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/registration", userHandler.CreateUser)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
