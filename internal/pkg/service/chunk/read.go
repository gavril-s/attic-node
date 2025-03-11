package chunk

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gavril-s/attic-node/internal/config"
)

type ChunkWithError struct {
	Chunk []byte
	Error error
}

func readChunk(reader *bufio.Reader, buf []byte) (message ChunkWithError, eof bool) {
	bytesRead, err := reader.Read(buf)
	if err != nil {
		if err != io.EOF {
			message.Error = err
		} else {
			eof = true
		}
		return message, eof
	}

	if bytesRead == len(buf) {
		message.Chunk = buf
	} else {
		message.Chunk = buf[:bytesRead]
	}
	return message, eof
}

func ReadFileInChunks(
	fileID string, storagePath config.Path, chunkSize config.Capacity,
) (chunkChan chan ChunkWithError, err error) {
	chunkChan = make(chan ChunkWithError)

	filepath := filepath.Join(storagePath.String(), fileID)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	go func() {
		reader := bufio.NewReader(file)
		buf := make([]byte, chunkSize.Bytes())
		for {
			message, eof := readChunk(reader, buf)
			chunkChan <- message
			if message.Error != nil || eof {
				close(chunkChan)
				break
			}
		}
	}()

	return chunkChan, nil
}
