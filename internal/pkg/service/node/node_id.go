package node

import (
	"fmt"
	"io"
	"os"

	"github.com/gavril-s/attic-node/internal/config"
)

func GetNodeID() (string, error) {
	file, err := os.Open(config.NodeIDFilePath.String())
	if err != nil {
		return "", fmt.Errorf("error openning node id file: %w", err)
	}
	defer file.Close()
	id, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error reading from node id file: %w", err)
	}
	return string(id), nil
}

func WriteNodeID(id string) error {
	file, err := os.OpenFile(config.NodeIDFilePath.String(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening node id file: %w", err)
	}
	defer file.Close()
	_, err = file.WriteString(id)
	if err != nil {
		return fmt.Errorf("error writing id to node id file: %w", err)
	}
	return nil
}
