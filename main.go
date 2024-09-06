package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	handlerBook "github.com/Ph4ra0hXX/go-book-api/book/handler"
	handlerPage "github.com/Ph4ra0hXX/go-book-api/page/handler"
	"github.com/Ph4ra0hXX/go-book-api/translation/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Database struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		User     string `toml:"user"`
		Password string `toml:"password"`
		Name     string `toml:"name"`
	} `toml:"database"`
	Server struct {
		Host string `toml:"host"`
		Port int    `toml:"port"`
	} `toml:"server"`
}

var config Config

func init() {
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatalf("Erro ao ler o arquivo de configuração: %v", err)
	}
}

func main() {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/books", handlerBook.GetBooks)
	router.GET("/books/:id", handlerBook.GetBookByID)
	router.POST("/books", handlerBook.CreateBook)
	router.PUT("/books", handlerBook.UpdateBook)
	router.DELETE("/books/:id", handlerBook.DeleteBook)

	router.GET("/pages/:book_id/pages", handlerPage.GetPagesHandler)
	router.GET("/pages/:book_id/pages/:page_number", handlerPage.GetPageByIDHandler)
	router.POST("/pages", handlerPage.CreatePageHandler)
	router.PUT("/pages/:book_id/pages/:page_number", handlerPage.UpdatePageHandler)
	router.DELETE("/pages/:book_id/pages/:page_number", handlerPage.DeletePageHandler)

	router.GET("/translate/:word", handler.GetTranslationHandler)

	address := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	router.Run(address)
}
