package main

import (
	"context"
	"fmt"
	"lowkyvideo/db"
	"os/exec"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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
	db.SeedConfig()

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

func (a *App) OpenVideoFolderDialog() (string, error) {
	directory, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{Title: "Select existing Lowky Video folder or select parent folder where a Lowky Video folder will be created.", CanCreateDirectories: true})
	if err != nil {
		fmt.Println("Error selecting directory: ", err)
		return "", err
	}
	return directory, nil
}

func (a *App) SetConfigDirectory(directory string) error {
	return db.UpdateConfigDirectory(directory)
}

func (a *App) GetConfig() (*db.Config, error) {
	return db.GetConfig()
}

func (a *App) GetProfiles() ([]*db.Profile, error) {
	return db.GetProfiles()
}

func (a *App) CreateProfile(name string) error {
	return db.CreateProfile(name)
}

func (a *App) GetProfile(id int) (*db.Profile, error) {
	return db.GetProfile(id)
}
