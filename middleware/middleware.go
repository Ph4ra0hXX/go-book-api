package middleware

import (
	"net/http"
	"strings"

	"github.com/Ph4ra0hXX/go-book-api/service"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token não fornecido"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		_, err := service.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token inválido"})
			c.Abort()
			return
		}

		c.Next()
	}
}
