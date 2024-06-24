package db

import (
	"lowkyvideo/config"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	ID        int
	Directory string
}

func SeedConfig() error {
	config, err := GetConfig()

	if err == nil && config != nil {
		return nil
	}

	stmt, err := DB.Prepare("INSERT INTO config(id, directory) VALUES(?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(1, "")
	return err
}

func GetConfig() (*Config, error) {
	var config Config
	err := DB.QueryRow("SELECT id, directory FROM config WHERE id = ?", 1).Scan(&config.ID, &config.Directory)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func folderExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func UpdateConfigDirectory(directory string) error {
	// check if directory exists
	if !folderExists(directory) {
		return os.ErrNotExist
	}
	// check if config.AppName is in directory name
	splitDir := strings.Split(directory, "/")
	lastEl := splitDir[len(splitDir)-1]
	if lastEl != config.AppName {
		// if not, see if config.AppName folder exists in directory
		// if neither exists, create config.AppName folder in directory and add to directory path
		directory = filepath.Join(directory, config.AppName)
		if !folderExists(directory) {
			err := os.Mkdir(directory, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	stmt, err := DB.Prepare("UPDATE config SET directory = ? WHERE id = 1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(directory)
	return err
}
