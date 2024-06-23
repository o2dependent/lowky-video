package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"lowkyvideo/config"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// App struct
type App struct {
	ctx context.Context
}

var crunchyCliInstalled = false

// check if cli tool exists
func checkTool(toolName string) (bool, error) {
	cmd := exec.Command("which", toolName)
	if _, err := cmd.Output(); err != nil {
		return false, err
	}
	return true, nil
}

// install cli tool
func installTool(ctx context.Context, toolName string) (bool, error) {
	cmd := exec.CommandContext(ctx, "brew", "install", toolName)

	if err := cmd.Start(); err != nil {
		return false, fmt.Errorf("error installing %s: %v", toolName, err)
	}

	done := make(chan error)

	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-ctx.Done():
		if err := cmd.Process.Kill(); err != nil {
			return false, fmt.Errorf("error killing process: %w", err)
		}
		return false, ctx.Err()
	case err := <-done:
		if err != nil {
			return false, fmt.Errorf("failed to install %s: %w", toolName, err)
		}
		return true, nil
	}
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	crunchyCliInstalled, err := checkTool("crunchy-cli")

	if !crunchyCliInstalled || err != nil {
		fmt.Printf("error checking for %s: %v\n", "crunchy-cli", err)
	}

	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetCLIInstalledStatus() bool {
	crunchyCliInstalled, err := checkTool("crunchy-cli")

	if !crunchyCliInstalled || err != nil {
		fmt.Printf("error checking for %s: %v\n", "crunchy-cli", err)
	}

	return crunchyCliInstalled
}

func (a *App) InstallCrunchyCLI() bool {
	// TODO
	return true
}

func (a *App) LoginToCrunchyCLI() bool {
	// TODO
	return true
}

var db *sql.DB

func getAppDataDir() (string, error) {
	var appDataDir string
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	switch runtime.GOOS {
	case "windows":
		appDataDir = filepath.Join(os.Getenv("APPDATA"), config.AppName)
	case "darwin":
		appDataDir = filepath.Join(homeDir, "Library", "Application Support", config.AppName)
	case "linux":
		appDataDir = filepath.Join(homeDir, ".local", "share", config.AppName)
	default:
		appDataDir = filepath.Join(homeDir, config.AppName)
	}

	err = os.MkdirAll(appDataDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	return appDataDir, nil
}

func initDB() {
	appDataDir, err := getAppDataDir()
	if err != nil {
		log.Fatal(err)
	}

	dbPath := filepath.Join(appDataDir, "app.db")
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	createTable := `CREATE TABLE IF NOT EXISTS profiles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) createProfile(name string) {
	insertProfile := `INSERT INTO profiles (name) VALUES (?);`
	_, err := db.Exec(insertProfile, name)
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) getProfiles() []string {
	var profiles []string
	rows, err := db.Query("SELECT name FROM profiles")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		profiles = append(profiles, name)
	}

	return profiles
}
