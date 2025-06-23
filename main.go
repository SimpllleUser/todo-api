package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsActive    bool   `json:"isACtive"`
}

var todos = []todo{
	{ID: "1", Title: "Todo - 1", Description: "Some example todo description - 1", IsActive: false},
	{ID: "2", Title: "Todo - 2", Description: "Some example todo description - 2", IsActive: true},
	{ID: "3", Title: "Todo - 3", Description: "Some example todo description - 3", IsActive: false},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func getMain(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "You on main route of api",
	})
}

func postTodo(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoByID(c *gin.Context) {
	id := c.Param("id")
	for _, t := range todos {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotExtended, gin.H{"message": "Todo not found :("})
}

func main() {
	router := gin.Default()
	router.GET("/", getMain)
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoByID)
	router.POST("/todos", postTodo)

	router.Run("localhost:8080")
}
