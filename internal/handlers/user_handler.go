package handler

import (
	model "example/todo-api/internal/models"
	service "example/todo-api/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userService *service.UserService
	authService *service.AuthService
}

func NewUserController(uService *service.UserService, aService *service.AuthService) *UserController {
	return &UserController{
		userService: uService,
		authService: aService,
	}
}

var errorBody = gin.H{"error": "Internal server error"}

func (uc *UserController) CreateUser(c *gin.Context) {

	var authInput model.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, errorBody)
		return
	}

	var userFound, err = uc.userService.FindByLogin(authInput.Login)
	if userFound.ID != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "User with this login already exists",
		})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Could not hash password:", err)
		c.JSON(http.StatusInternalServerError, errorBody)
		return
	}

	user := &model.UserModel{
		Login:    authInput.Login,
		Password: string(passwordHash),
	}

	err = uc.userService.Create(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorBody)
		return
	}

	token, err := uc.authService.GenerateToken(user.ID)
	if err != nil {
		log.Println("Could not generate token:", err)
		c.JSON(http.StatusInternalServerError, errorBody)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"user":  user,
			"token": token,
		},
	})
}
