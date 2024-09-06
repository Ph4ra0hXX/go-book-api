package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	handlerBook "github.com/Ph4ra0hXX/go-book-api/book/handler"
	"github.com/Ph4ra0hXX/go-book-api/middleware"
	handlerPage "github.com/Ph4ra0hXX/go-book-api/page/handler"
	handlerTranslation "github.com/Ph4ra0hXX/go-book-api/translation/handler"
	handlerUser "github.com/Ph4ra0hXX/go-book-api/user/handler"

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
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Authorization", "Content-Length"},
		AllowCredentials: true,
	}))

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middleware.AuthMiddleware())

	protectedRoutes.GET("/books", handlerBook.GetBooks)
	protectedRoutes.GET("/books/:id", handlerBook.GetBookByID)
	protectedRoutes.POST("/books", handlerBook.CreateBook)
	protectedRoutes.PUT("/books", handlerBook.UpdateBook)
	protectedRoutes.DELETE("/books/:id", handlerBook.DeleteBook)

	protectedRoutes.GET("/pages/:book_id/pages", handlerPage.GetPagesHandler)
	protectedRoutes.GET("/pages/:book_id/pages/:page_number", handlerPage.GetPageByIDHandler)
	protectedRoutes.POST("/pages", handlerPage.CreatePageHandler)
	protectedRoutes.PUT("/pages/:book_id/pages/:page_number", handlerPage.UpdatePageHandler)
	protectedRoutes.DELETE("/pages/:book_id/pages/:page_number", handlerPage.DeletePageHandler)

	router.GET("/translate/:word", handlerTranslation.GetTranslationHandler)

	router.POST("/register", handlerUser.RegisterUserHandler)
	router.POST("/login", handlerUser.LoginUserHandler)

	address := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	router.Run(address)
}
