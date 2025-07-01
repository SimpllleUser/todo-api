package main

import (
	"example/todo-api/config"
	"example/todo-api/internal/database"
	handler "example/todo-api/internal/handlers"
	model "example/todo-api/internal/models"
	"example/todo-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	const PATH_TO_DB = "../internal/database/app-db.db"

	database.InitDB("internal/database/app-db.db")

	defer database.CloseDB()

	config.LoadEnv()

	todoService := model.NewTodoService(database.DB)
	userService := model.NewUserService(database.DB)

	todoController := handler.NewTodoController(todoService)
	userController := handler.NewUserController(userService)
	authController := handler.NewAuthController(userService)

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
