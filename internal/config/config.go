package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func getConfigFilePath() (string, error) {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error returning the current user's home directory: %w", err)
	}

	return homeDirectory, nil
}

func Read() (Config, error) {
	homeDirectory, err := getConfigFilePath()
	configStruct := Config{}
	if err != nil {
		return configStruct, fmt.Errorf("error returning the current user's home directory: %w", err)
	}

	jsonFile, err := os.Open(homeDirectory + "/" + configFilename)
	if err != nil {
		return configStruct, fmt.Errorf("error opening the .gatorconfig.json file: %w", err)
	}

	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		return configStruct, fmt.Errorf("error reading JSON file: %w", err)
	}

	err = json.Unmarshal(data, &configStruct)
	if err != nil {
		return configStruct, fmt.Errorf("error unmarshalling JSON file to struct:\n %w", err)
	}

	return configStruct, nil
}

func (c *Config) SetUser(currentUsername string) error {
	homeDirectory, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error returning the current user's home directory:\n %w", err)
	}

	c.CurrentUsername = currentUsername

	configJSON, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("error marshalling Config struct to JSON:\n %w", err)
	}

	err = os.WriteFile(homeDirectory+"/"+configFilename, configJSON, 0644)
	if err != nil {
		return fmt.Errorf("error writing config JSON to file:\n %w", err)
	}

	return nil
}
