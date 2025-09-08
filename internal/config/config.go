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

func (cfg *Config) SetUser(user_name string) error {
	cfg.CurrentUserName = user_name
	err := write(*cfg)
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

	b, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// No file yet: return default/empty config
			return Config{}, nil
		}
		return Config{}, fmt.Errorf("error opening file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return Config{}, fmt.Errorf("error unmarshaling JSON: %w", err)
	}
	return cfg, nil
}

func getConfigFilePath() (string, error) {
	const configFileName = ".gatorconfig.json"

	// Optional override for tricky environments
	if p := os.Getenv("GATOR_CONFIG"); p != "" {
		return p, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("get home dir: %w", err)
	}
	return filepath.Join(home, configFileName), nil
}
