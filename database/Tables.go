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

	usersT := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL,
        email TEXT,
        password TEXT
    );`
	postsT := `
    CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL,
		post TEXT NOT NULL,
		category TEXT
    );`

	_, err = db.Exec(usersT)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}
	_, err = db.Exec(postsT)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}
	return db, nil
}
