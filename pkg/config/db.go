package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// TODO make connStr able to be changed
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database connected")
	}

	TableInit()
}

func TableInit() {
	query := `
	CREATE TABLE IF NOT EXISTS todolist (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		content TEXT,
		status BOOLEAN NOT NULL DEFAULT FALSE
	);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table checked/created")
}
