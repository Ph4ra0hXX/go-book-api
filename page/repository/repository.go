package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Ph4ra0hXX/go-book-api/page/model"
	_ "github.com/lib/pq" // PostgreSQL driver
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

func GetPages(bookID int) []model.Page {
	rows, err := db.Query("SELECT book_id, page_number, text FROM pages WHERE book_id = $1", bookID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pages []model.Page
	for rows.Next() {
		var page model.Page
		if err := rows.Scan(&page.BookID, &page.PageNumber, &page.Text); err != nil {
			log.Fatal(err)
		}
		pages = append(pages, page)
	}
	return pages
}

func GetPageByID(bookID, pageNumber int) *model.Page {
	row := db.QueryRow("SELECT book_id, page_number, text FROM pages WHERE book_id = $1 AND page_number = $2", bookID, pageNumber)
	var page model.Page
	if err := row.Scan(&page.BookID, &page.PageNumber, &page.Text); err != nil {
		return nil
	}
	return &page
}

func CreatePage(newPage model.Page) {
	_, err := db.Exec("INSERT INTO pages (book_id, page_number, text) VALUES ($1, $2, $3)", newPage.BookID, newPage.PageNumber, newPage.Text)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdatePage(updatedPage model.Page) bool {
	result, err := db.Exec("UPDATE pages SET text = $3 WHERE book_id = $1 AND page_number = $2", updatedPage.BookID, updatedPage.PageNumber, updatedPage.Text)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected > 0
}

func DeletePage(bookID, pageNumber int) bool {
	result, err := db.Exec("DELETE FROM pages WHERE book_id = $1 AND page_number = $2", bookID, pageNumber)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected > 0
}
