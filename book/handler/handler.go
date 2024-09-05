package handler

import (
	"net/http"
	"strconv"

	"github.com/Ph4ra0hXX/go-book-api/book/model"
	"github.com/Ph4ra0hXX/go-book-api/book/repository"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	books := repository.GetBooks()
	c.IndentedJSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}
	book := repository.GetBookByID(id)
	if book == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Livro não encontrado"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var newBook model.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Dados inválidos"})
		return
	}
	repository.CreateBook(newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
	var updatedBook model.Book
	if err := c.BindJSON(&updatedBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Dados inválidos"})
		return
	}

	if repository.UpdateBook(updatedBook) {
		c.IndentedJSON(http.StatusOK, updatedBook)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Livro não encontrado"})
	}
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	if repository.DeleteBook(id) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Livro excluído"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Livro não encontrado"})
	}
}
