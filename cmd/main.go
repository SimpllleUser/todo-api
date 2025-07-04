package main

import (
	"example/todo-api/config"
	"example/todo-api/internal/database"
	handler "example/todo-api/internal/handlers"
	"example/todo-api/internal/routes"
	service "example/todo-api/internal/services"

	"github.com/gin-gonic/gin"
)

//	@title						TODO API
//	@version					0.0.1
//	@description				This is a simple API for managing TODOs.
//	@host						localhost:8080
//	@BasePath					/api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your token **with** Bearer prefix, e.g. `Bearer <token>`

func main() {

	config.LoadEnv()

	const PATH_TO_DB = "internal/database/app-db.db"

	database.InitDB(PATH_TO_DB)

	defer database.CloseDB()

	todoService := service.NewTodoService(database.DB)
	userService := service.NewUserService(database.DB)
	authService := service.NewAuthService(userService)

	todoController := handler.NewTodoController(todoService)
	userController := handler.NewUserController(userService, authService)
	authController := handler.NewAuthController(authService)

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routes.SetupRoutes(r,
		todoController,
		userController,
		authController,
		userService,
	)

	r.Run("localhost:8080")
}
