package database

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func CreateTable() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database/Database.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL,
        email TEXT,
        password TEXT
    );`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}
	return db, nil
}
