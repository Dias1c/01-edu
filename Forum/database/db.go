package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"forum/configs"

	_ "github.com/mattn/go-sqlite3"
)

// initialize connection with sqlite db
// returns (connection, error)
func InitDB() (*sql.DB, error) {
	err := os.MkdirAll(configs.DB_PATH, os.ModeDir)
	if err != nil {
		return nil, fmt.Errorf("InitDB: %w", err)
	}
	dbPathField := filepath.Join(configs.DB_PATH, configs.DB_NAME)
	// connection field from configs
	if configs.DB_USERNAME != "" && configs.DB_PASSWORD != "" {
		authField := fmt.Sprintf("?_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=%s", configs.DB_USERNAME, configs.DB_PASSWORD, configs.DB_AUTHCRYPT)
		dbPathField += authField
	}
	db, err := sql.Open(configs.DBDriverName, dbPathField)
	if err != nil {
		return nil, fmt.Errorf("InitDB: %w", err)
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("InitDB: %w", err)
	}
	err = checkDB(db)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("InitDB: %w", err)
	}
	db.SetMaxIdleConns(100)
	return db, err
}

// check the scheme
func checkDB(db *sql.DB) error {
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			nickname TEXT,
			email TEXT,
			createdAt TEXT,
			password TEXT,
			firstName TEXT,
			lastName TEXT
		);`,
	)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS sessions (
			id INTEGER PRIMARY KEY,
			uuid TEXT,
			user_id INTEGER,
			created_at TEXT
		);`,
	)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS tags (
			id INTEGER PRIMARY KEY,
			name TEXT
		);`,
	)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY,
			content TEXT,
			user_id INTEGER,
			created_at TEXT
		);`,
	)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS comments (
			id INTEGER PRIMARY KEY,
			content TEXT,
			user_id INTEGER,
			created_at TEXT,
			post_id INTEGER
		);`,
	)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS questions (
			id INTEGER PRIMARY KEY,
			post_id INTEGER,
			title TEXT,
			views INTEGER
		);`,
	)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS likes (
			id INTEGER PRIMARY KEY,
			user_id TEXT,
			post_id TEXT,
			liked INTEGER
		);`,
	)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS tags_questions (
			id INTEGER PRIMARY KEY,
			tag_ig INTEGER,
			question_id INTEGER
		);`,
	)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}
	return nil
}
