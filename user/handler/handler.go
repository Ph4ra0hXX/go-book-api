package handler

import (
	"net/http"

	"github.com/Ph4ra0hXX/go-book-api/service"
	"github.com/Ph4ra0hXX/go-book-api/user/model"
	"github.com/Ph4ra0hXX/go-book-api/user/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RegisterUserHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criptografar a senha"})
		return
	}
	user.Password = hashedPassword

	err = repository.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criar o usuário"})
		return
	}

	token, err := service.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao gerar token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func LoginUserHandler(c *gin.Context) {
	var loginRequest model.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	storedUser, err := repository.GetUserByUsername(loginRequest.Username)
	if err != nil || storedUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha incorretos"})
		return
	}

	if !CheckPasswordHash(loginRequest.Password, storedUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha incorretos"})
		return
	}

	token, err := service.GenerateToken(storedUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
