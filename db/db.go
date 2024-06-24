package db

import (
	"database/sql"
	"fmt"
	"lowkyvideo/config"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func getAppDataDir() (string, error) {
	var appDataDir string
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// switch runtime.GOOS {
	// case "windows":
	// 	appDataDir = filepath.Join(os.Getenv("APPDATA"), config.AppName)
	// case "darwin":
	// 	appDataDir = filepath.Join(homeDir, "Library", "Application Support", config.AppName)
	// case "linux":
	// 	appDataDir = filepath.Join(homeDir, ".local", "share", config.AppName)
	// default:
	// 	appDataDir = filepath.Join(homeDir, config.AppName)
	// }
	appDataDir = filepath.Join(homeDir, "Users", "ethanolsen", "Desktop", config.AppName)
	fmt.Printf("App data directory: %s\n", appDataDir)
	err = os.MkdirAll(appDataDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	return appDataDir, nil
}

func InitDB() error {
	var err error

	appDataDir, err := getAppDataDir()

	if err != nil {
		fmt.Printf("Error getting app data directory: %v\n", err)
		return err
	}

	dbPath := filepath.Join(appDataDir, "app.db")

	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		return err
	}

	createTable := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
    );`

	_, err = DB.Exec(createTable)
	if err != nil {
		fmt.Printf("Error creating table: %v\n", err)
		return err
	}

	return nil
}
