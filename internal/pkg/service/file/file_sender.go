package file

import (
	"fmt"
	"path/filepath"

	"github.com/gavril-s/attic-node/internal/config"
	io "github.com/gavril-s/attic-node/internal/pkg/service/chunk"
)

type FileSender struct {
	Cfg          *config.Config
	FileID       string
	ChunkSize    uint64
	StorageIndex uint64
	ChunkChan    chan io.ChunkWithError
}

func NewFileSender(
	cfg *config.Config, fileID string, chunkSize uint64, storageIndex uint64,
) (*FileSender, error) {
	sender := &FileSender{
		Cfg:          cfg,
		FileID:       fileID,
		ChunkSize:    chunkSize,
		StorageIndex: storageIndex,
	}
	if storageIndex >= uint64(len(cfg.Storages)) {
		return nil, fmt.Errorf("storageIndex is out of range for the given config")
	}
	return sender, nil
}

func (f *FileSender) Init() (chunksNumber uint64, err error) {
	f.ChunkChan, err = io.ReadFileInChunks(f.FileID, f.Cfg.Storages[f.StorageIndex].Path, config.Capacity(f.ChunkSize))
	if err != nil {
		return 0, err
	}

	filepath := config.Path(filepath.Join(f.Cfg.Storages[f.StorageIndex].Path.String(), f.FileID))
	size, err := filepath.Size()
	if err != nil {
		return 0, fmt.Errorf("error getting file size: %w", err)
	}
	if size.Bytes()%f.ChunkSize == 0 {
		chunksNumber = size.Bytes() / f.ChunkSize
	} else {
		chunksNumber = size.Bytes()/f.ChunkSize + 1
	}
	return chunksNumber, nil
}

func (f *FileSender) GetNextChunk() ([]byte, error) {
	chunk, ok := <-f.ChunkChan
	if !ok {
		return nil, fmt.Errorf("error receiving chunk from channel (possibly the channel is closed)")
	}
	if chunk.Error != nil {
		return nil, fmt.Errorf("error receiving chunk: %w", chunk.Error)
	}
	return chunk.Chunk, nil
}
