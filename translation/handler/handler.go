package handler

import (
	"net/http"

	"github.com/Ph4ra0hXX/go-book-api/translation/repository"
	"github.com/gin-gonic/gin"
)

func GetTranslationHandler(c *gin.Context) {
	word := c.Param("word")

	translation, err := repository.GetTranslation(word)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao buscar a tradução"})
		return
	}
	if translation == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tradução não encontrada"})
		return
	}

	c.JSON(http.StatusOK, translation)
}
