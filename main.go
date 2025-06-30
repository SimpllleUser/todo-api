package main

import (
	"example/todo-api/config"
	"example/todo-api/controller"
	"example/todo-api/model"
	"example/todo-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	model.ConnectDatabase()

	defer model.CloseDB()

	config.LoadEnv()

	todoService := model.NewTodoService(model.DB)
	userService := model.NewUserService(model.DB)

	todoController := controller.NewTodoController(todoService)
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(userService)

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
