package main

import (
	"context"
	"fmt"
	"lowkyvideo/db"
	"os/exec"
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

	db.InitDB()

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

func (a *App) GetUsers() ([]*db.User, error) {
	return db.GetUsers()
}

func (a *App) CreateUser(name string) error {
	return db.CreateUser(name)
}

func (a *App) GetUser(id int) (*db.User, error) {
	return db.GetUser(id)
}
