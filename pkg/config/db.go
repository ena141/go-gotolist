package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "user=price dbname=todoapp sslmode=disable password=123456"
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
		status BOOLEAN NOT FULL DEFAULT FALSE
	);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table checked/created")
}
