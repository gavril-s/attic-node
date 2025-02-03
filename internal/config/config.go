package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	ConfigPath    = "./config.json"
	FileChunkSize = 1024
)

type StorageConfig struct {
	Path     string `json:"path"`
	Capacity string `json:"capacity"`
}

type Config struct {
	MasterURL string          `json:"master_url"`
	Host      string          `json:"host"`
	Port      uint16          `json:"port"`
	Storages  []StorageConfig `json:"storages"`
}

func NewConfig() (*Config, error) {
	configFile, err := os.Open(ConfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config: %w", err)
	}
	defer configFile.Close()

	configContent, err := io.ReadAll(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var config Config
	err = json.Unmarshal(configContent, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &config, nil
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
