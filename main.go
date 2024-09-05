package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Ph4ra0hXX/go-book-api/book/handler"
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
	// Carregar configurações do arquivo TOML
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatalf("Erro ao ler o arquivo de configuração: %v", err)
	}
}

func main() {

	router := gin.Default()

	// Rotas para o CRUD de livros
	router.GET("/books", handler.GetBooks)
	router.GET("/books/:id", handler.GetBookByID)
	router.POST("/books", handler.CreateBook)
	router.PUT("/books", handler.UpdateBook)
	router.DELETE("/books/:id", handler.DeleteBook)

	// Iniciar o servidor
	address := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	router.Run(address)
}
