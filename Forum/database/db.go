package database

import (
	"database/sql"
	"fmt"
	"log"
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
		log.Printf("Error while trying create directory.\nDescription: %s", err.Error())
		return nil, err
	}
	dbPathField := filepath.Join(configs.DB_PATH, configs.DB_NAME)
	// connection field from configs
	if configs.DB_USERNAME != "" && configs.DB_PASSWORD != "" {
		authField := fmt.Sprintf("?_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=%s", configs.DB_USERNAME, configs.DB_PASSWORD, configs.DB_AUTHCRYPT)
		dbPathField += authField
	}
	db, err := sql.Open(configs.DBDriverName, dbPathField)
	if err != nil {
		log.Printf("Error while trying connect to database.\nDescription: %s", err.Error())
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Printf("Error while trying ping to database.\nDescription: %s", err.Error())
		db.Close()
		return nil, err
	}
	// db.SetMaxIdleConns(100)
	return db, err
}
