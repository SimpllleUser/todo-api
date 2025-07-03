package handler

import (
	model "example/todo-api/internal/models"
	service "example/todo-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(aService *service.AuthService) *AuthController {
	return &AuthController{
		authService: aService,
	}
}

// Login godoc
//
//	@Summary		Login user
//	@Description	Login user
//	@Tags			Auth
//	@Param			user	body	model.AuthInput	true	"User data"
//	@Produce		json
//	@accept			json
//	@Success		200	{object}	model.AuthResponse
//	@Failure		400	{object}	model.HTTPError	"Invalid request"
//	@Failure		401	{object}	model.HTTPError	"Error authenticating user"
//	@Failure		500	{object}	model.HTTPError	"Could not generate token"
//	@Router			/auth/login [post]
func (ac *AuthController) Login(c *gin.Context) {
	var authInput model.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userFound, err := ac.authService.Authenticate(authInput.Login, authInput.Password)
	if err != nil || userFound.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Error authenticating user",
			"message": err.Error(),
		})
		return
	}

	token, err := ac.authService.GenerateToken(userFound.ID)
	if err != nil || userFound == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": token,
	})

}
