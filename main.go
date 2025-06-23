package main

import (
	"example/todo-api/controller"
	"example/todo-api/model"
	"example/todo-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	model.ConnectDatabase()

	defer model.CloseDB()

	todoService := model.NewTodoService(model.DB)
	todoController := controller.NewTodoController(todoService)

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routes.SetupRoutes(r, todoController)

	r.Run("localhost:8080")
}
