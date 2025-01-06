package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const fileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	err := write(*cfg)
	return err
}

func Read() (Config, error) {
	configFile, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	fileContent, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func write(cfg Config) error {
	configFile, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(homeDir, fileName)
	return fullPath, nil
}
