package controller

import (
	"example/todo-api/model"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	userService *model.UserService
}

func NewAuthController(uService *model.UserService) *AuthController {
	return &AuthController{
		userService: uService,
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	var authInput model.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound, err = ac.userService.FindByLogin(authInput.Login)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		log.Println("Error comparing password:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Time to expire the token
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": token,
	})

}
