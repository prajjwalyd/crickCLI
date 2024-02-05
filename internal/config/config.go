package config

import (
	_"fmt"
	"log"
	"os"
	"path/filepath"
)

var configFilePath = getConfigFilePath()

// ReadStoredMatchID reads match ID from the config file
func ReadStoredMatchID() (string, error) {
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// StoreMatchID stores the new match ID in the config file
func StoreMatchID(newMatchID string) error {
	return os.WriteFile(configFilePath, []byte(newMatchID), 0644)
}

func getConfigFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configDir := filepath.Join(home, ".crickcli")
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(configDir, "config.txt")
}
