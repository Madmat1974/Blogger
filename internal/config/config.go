package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg Config) SetUser(user_name string) error {
	cfg.CurrentUserName = user_name
	err := write(cfg)
	if err != nil {
		return err
	}
	return nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	err = os.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		return fmt.Errorf("error modifying file: %v", err)
	}

	return nil
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	fBytes, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("error opening file: %v", err)
	}

	var cfg Config
	err = json.Unmarshal(fBytes, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	const configFileName = "blogger/gatorconfig.json"
	homepath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error with home directory")
	}

	fpath := filepath.Join(homepath, configFileName)
	return fpath, nil

}
