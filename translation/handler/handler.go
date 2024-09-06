package handler

import (
	"net/http"

	"github.com/Ph4ra0hXX/go-book-api/translation/model"
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

func CreateTranslationHandler(c *gin.Context) {
	var translation model.Translation
	if err := c.ShouldBindJSON(&translation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	err := repository.CreateTranslation(&translation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criar a tradução"})
		return
	}
	c.JSON(http.StatusCreated, translation)
}

func UpdateTranslationHandler(c *gin.Context) {
	word := c.Param("word")
	var updatedTranslation model.Translation
	if err := c.ShouldBindJSON(&updatedTranslation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	err := repository.UpdateTranslation(word, &updatedTranslation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao atualizar a tradução"})
		return
	}
	c.JSON(http.StatusOK, updatedTranslation)
}

func DeleteTranslationHandler(c *gin.Context) {
	word := c.Param("word")
	err := repository.DeleteTranslation(word)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao deletar a tradução"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
