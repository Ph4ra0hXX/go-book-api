package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Ph4ra0hXX/go-book-api/user/model"
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

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := db.QueryRow("SELECT id, username, email, password FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *model.User) error {
	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		user.Username, user.Email, user.Password)
	return err
}

func UpdateUser(username string, user *model.User) error {
	_, err := db.Exec("UPDATE users SET email = $1, password = $2 WHERE username = $3",
		user.Email, user.Password, username)
	return err
}

func DeleteUser(username string) error {
	_, err := db.Exec("DELETE FROM users WHERE username = $1", username)
	return err
}
