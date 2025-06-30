package controller

import (
	"example/todo-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userService *model.UserService
}

func NewUserController(uService *model.UserService) *UserController {
	return &UserController{
		userService: uService,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var authInput model.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// var userFound, err = uc.userService.FindByLogin(authInput.Login)

	// println("User found:", userFound, "Error:", err)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := &model.User{
		Login:    authInput.Login,
		Password: string(passwordHash),
	}

	user, err = uc.userService.Create(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}
