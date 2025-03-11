package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	ConfigPath     Path = "./config.json"
	NodeIDFilePath Path = "./node_id.txt"
)

type Config struct {
	MasterURL         string
	Host              string
	Port              uint16
	IsPersistent      bool
	AcceptedChunkSize Capacity
	Storages          []Storage
}

func NewConfig(config JSONConfig) (*Config, error) {
	var err error

	storages := make([]Storage, len(config.Storages))
	for i, storage := range config.Storages {
		storages[i], err = NewStorage(storage)
		if err != nil {
			return nil, fmt.Errorf("NewStorage: %w", err)
		}
	}

	acceptedChunkSize, err := ParseCapacity(config.AcceptedChunkSize)
	if err != nil {
		return nil, fmt.Errorf("ParseCapacity: %w", err)
	}

	return &Config{
		MasterURL:         config.MasterURL,
		Host:              config.Host,
		Port:              config.Port,
		IsPersistent:      config.IsPersistent,
		AcceptedChunkSize: acceptedChunkSize,
		Storages:          storages,
	}, nil
}

func ReadConfig() (*Config, error) {
	configFile, err := os.Open(ConfigPath.String())
	if err != nil {
		return nil, fmt.Errorf("failed to open config: %w", err)
	}
	defer configFile.Close()

	configContent, err := io.ReadAll(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var config JSONConfig
	err = json.Unmarshal(configContent, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return NewConfig(config)
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
