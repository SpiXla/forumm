package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

// Global DB connection
var Db *sql.DB

// PostData represents the structure of a post in the database
type PostData struct {
    ID        int
    Username  string
    Text      string
    Category  string
    CreatedAt string // Store the created_at timestamp as a string
}
// CreateTable creates the necessary tables and initializes the database
func CreateTable() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database/Database.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	Db = db // Assign the opened DB connection to global Db variable

	// Create 'users' table if not exists
	usersT := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL,
        email TEXT,
        password TEXT
    );`

	// Create 'posts' table if not exists
	postsT := `
    CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL,
        post TEXT NOT NULL,
        category TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

	_, err = db.Exec(usersT)
	if err != nil {
		return nil, fmt.Errorf("failed to create users table: %w", err)
	}

	_, err = db.Exec(postsT)
	if err != nil {
		return nil, fmt.Errorf("failed to create posts table: %w", err)
	}

	return db, nil
}

// GetAllPosts retrieves all posts from the database
func GetAllPosts() ([]PostData, error) {
    rows, err := Db.Query("SELECT username, post, category, created_at FROM posts")
    if err != nil {
        return nil, fmt.Errorf("failed to query posts table: %w", err)
    }
    defer rows.Close()

    var posts []PostData
    for rows.Next() {
        var post PostData
        if err := rows.Scan(&post.Username, &post.Text, &post.Category, &post.CreatedAt); err != nil {
            return nil, fmt.Errorf("failed to scan row: %w", err)
        }
        posts = append(posts, post)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error during row iteration: %w", err)
    }

    return posts, nil
}
