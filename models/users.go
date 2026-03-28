package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var DB *sql.DB

func dbConfig() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables.")
	}

	return os.Getenv("DB_PATH")
}

func Init() {
	DB = setupDB()
}

func setupDB() *sql.DB {
	db, err := sql.Open("sqlite", dbConfig())
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	log.Println("Database connected successfully!")
	return db
}

func GetUsers() []User {
	var users []User

	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		log.Println("Error querying users: ", err)
	}

	defer rows.Close()

	for rows.Next() {

		// Scan row data
		var user User

		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Println("Error scanning user: ", err)
			continue
		}
		users = append(users, user)
	}
	return users
}

func GetUser(id int) (*User, error) {
	row, err := DB.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		log.Println("Error querying user: ", err)
		return nil, err
	}

	defer row.Close()

	if row.Next() {
		var user User

		err := row.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Println("Error scanning user: ", err)
			return nil, err
		}

		return &user, nil
	}
	return nil, fmt.Errorf("User not found")
}
