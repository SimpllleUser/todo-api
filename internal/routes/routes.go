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
	taskHandler *handler.TaskController,
	userHandler *handler.UserController,
	authHandler *handler.AuthController,
	boardHandler *handler.BoardController,
	userService *service.UserService,
) {
	api := r.Group("/api/v1")
	{
		tasks := api.Group("/tasks")
		tasks.Use(middlewares.CheckAuth(userService))
		{
			tasks.GET("", taskHandler.GetTasks)
			tasks.GET("/:id", taskHandler.GetTaskById)
			tasks.GET("/title/:title", taskHandler.GetTaskByTitle)
			tasks.POST("", taskHandler.CreateTasks)
			tasks.PATCH("/:id", taskHandler.UpdateTask)
			tasks.DELETE("/:id", taskHandler.DeleteTask)
			tasks.GET("/boards/:id", taskHandler.GetBoardTasks)
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
			boards.PATCH("/:id/users", boardHandler.AddUsers)
		}

		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/registration", userHandler.CreateUser)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
