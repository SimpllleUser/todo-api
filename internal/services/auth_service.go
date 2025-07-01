package service

import (
	model "example/todo-api/internal/models"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userService *UserService
}

const EXPIRE = time.Hour * 24

func NewAuthService(us *UserService) *AuthService {
	return &AuthService{
		userService: us,
	}
}

func (as *AuthService) Authenticate(login string, password string) (*model.UserModel, error) {
	var userFound, err = as.userService.FindByLogin(login)

	if err != nil {
		log.Println("Error find user by login", err)
		return nil, err
	}

	if err := compareHashAndPassword([]byte(userFound.Password), []byte(password)); err != nil {
		log.Println("Error comparing password:", err)
		return nil, err
	}

	return userFound, nil
}

func (as *AuthService) GenerateToken(userId uint) (string, error) {

	generateToken := getClaim(jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(EXPIRE).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		log.Println("Could not generate token", err)
		return "", err
	}

	return token, nil
}

func compareHashAndPassword(userPassword []byte, inputPassword []byte) error {
	if len(userPassword) == 0 || len(inputPassword) == 0 {
		return bcrypt.ErrMismatchedHashAndPassword
	}
	return bcrypt.CompareHashAndPassword(userPassword, inputPassword)
}

func getClaim(claims jwt.Claims) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}
