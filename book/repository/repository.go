package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Ph4ra0hXX/go-book-api/book/model"
	_ "github.com/lib/pq"
)

var (
	db     *sql.DB
	config Config
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
}

func init() {
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatal(err)
	}

	connStr := getDBConnectionString()
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func getDBConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.Name)
}

func GetBooks() []model.Book {
	rows, err := db.Query("SELECT id, image, author, title FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.ID, &book.Image, &book.Author, &book.Title); err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	return books
}

func GetBookByID(id int) *model.Book {
	row := db.QueryRow("SELECT id, image, author, title FROM books WHERE id = $1", id)
	var book model.Book
	if err := row.Scan(&book.ID, &book.Image, &book.Author, &book.Title); err != nil {
		return nil
	}
	return &book
}

func CreateBook(newBook model.Book) {
	_, err := db.Exec("INSERT INTO books (id, image, author, title) VALUES ($1, $2, $3, $4)", newBook.ID, newBook.Image, newBook.Author, newBook.Title)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateBook(updatedBook model.Book) bool {
	result, err := db.Exec("UPDATE books SET image = $2, author = $3, author = $4 WHERE id = $1", updatedBook.ID, updatedBook.Image, updatedBook.Author, updatedBook.Title)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected > 0
}

func DeleteBook(id int) bool {
	result, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected > 0
}
