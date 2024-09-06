package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Ph4ra0hXX/go-book-api/translation/model"
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

func GetTranslation(word string) (*model.Translation, error) {
	var translation model.Translation
	err := db.QueryRow("SELECT word, translation FROM Translation WHERE word = $1", word).Scan(&translation.Word, &translation.Translation)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &translation, nil
}

func CreateTranslation(translation *model.Translation) error {
	_, err := db.Exec("INSERT INTO Translation (word, translation) VALUES ($1, $2)", translation.Word, translation.Translation)
	return err
}

func UpdateTranslation(word string, translation *model.Translation) error {
	_, err := db.Exec("UPDATE Translation SET translation = $1 WHERE word = $2", translation.Translation, word)
	return err
}

func DeleteTranslation(word string) error {
	_, err := db.Exec("DELETE FROM Translation WHERE word = $1", word)
	return err
}
