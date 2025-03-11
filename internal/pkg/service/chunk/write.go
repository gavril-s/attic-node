package chunk

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gavril-s/attic-master/api"
	"github.com/gavril-s/attic-node/internal/config"
)

func chunkFilename(fileID string, chunkIndex uint64) string {
	return fmt.Sprintf("%s_%d", fileID, chunkIndex)
}

func WriteChunk(chunk api.Chunk, storagePath config.Path) error {
	file, err := os.Create(filepath.Join(
		storagePath.String(),
		chunkFilename(chunk.FileID, chunk.ChunkIndex),
	))
	if err != nil {
		return fmt.Errorf("error creating file for chunk: %w", err)
	}
	defer file.Close()
	_, err = file.Write(chunk.Payload)
	if err != nil {
		return fmt.Errorf("error writing chunk to file: %w", err)
	}
	return nil
}

func MergeChunks(fileID string, chunksNumber uint64, storagePath config.Path) error {
	outputFile, err := os.Create(filepath.Join(storagePath.String(), fileID))
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer outputFile.Close()

	for chunkIndex := range chunksNumber {
		chunkFilePath := filepath.Join(
			storagePath.String(),
			chunkFilename(fileID, chunkIndex),
		)
		chunkFile, err := os.Open(chunkFilePath)
		if err != nil {
			return fmt.Errorf("error opening chunk file: %w", err)
		}

		chunk, err := io.ReadAll(chunkFile)
		if err != nil {
			return fmt.Errorf("error reading chunk from file: %w", err)
		}
		_, err = outputFile.Write(chunk)
		if err != nil {
			return fmt.Errorf("error writing chunk to output file: %w", err)
		}

		err = chunkFile.Close()
		if err != nil {
			return fmt.Errorf("error closing chunk file: %w", err)
		}
		err = os.Remove(chunkFilePath)
		if err != nil {
			return fmt.Errorf("error removing chunk file: %w", err)
		}
	}
	return nil
}
