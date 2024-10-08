package handler

import (
	"net/http"
	"strconv"

	"github.com/Ph4ra0hXX/go-book-api/page/model"
	"github.com/Ph4ra0hXX/go-book-api/page/repository"
	"github.com/gin-gonic/gin"
)

func GetPagesHandler(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do livro inválido"})
		return
	}

	pages := repository.GetPages(bookID)
	c.JSON(http.StatusOK, pages)
}

func GetPageByIDHandler(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do livro inválido"})
		return
	}
	pageNumber, err := strconv.Atoi(c.Param("page_number"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Número da página inválido"})
		return
	}

	page := repository.GetPageByID(bookID, pageNumber)
	if page == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Página não encontrada"})
		return
	}
	c.JSON(http.StatusOK, page)
}

func CreatePageHandler(c *gin.Context) {
	var newPage model.Page
	if err := c.ShouldBindJSON(&newPage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	repository.CreatePage(newPage)
	c.JSON(http.StatusCreated, gin.H{"message": "Página criada com sucesso"})
}

func UpdatePageHandler(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do livro inválido"})
		return
	}
	pageNumber, err := strconv.Atoi(c.Param("page_number"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Número da página inválido"})
		return
	}

	var updatedPage model.Page
	if err := c.ShouldBindJSON(&updatedPage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	updatedPage.BookID = bookID
	updatedPage.PageNumber = pageNumber

	if repository.UpdatePage(updatedPage) {
		c.JSON(http.StatusOK, gin.H{"message": "Página atualizada com sucesso"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Página não encontrada"})
	}
}

func DeletePageHandler(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do livro inválido"})
		return
	}
	pageNumber, err := strconv.Atoi(c.Param("page_number"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Número da página inválido"})
		return
	}

	if repository.DeletePage(bookID, pageNumber) {
		c.JSON(http.StatusOK, gin.H{"message": "Página removida com sucesso"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Página não encontrada"})
	}
}
