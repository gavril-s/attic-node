package file

import (
	"fmt"

	"github.com/gavril-s/attic-master/api"
	"github.com/gavril-s/attic-node/internal/config"
	io "github.com/gavril-s/attic-node/internal/pkg/service/chunk"
)

type FileReceiver struct {
	Cfg            *config.Config
	FileID         string
	Size           config.Capacity
	ChunksNumber   uint64
	ChunksReceived map[uint64]bool
	StorageIndexes []uint64
}

func NewFileReceiver(cfg *config.Config, fileHeader api.MasterFileHeader) (*FileReceiver, error) {
	receiver := &FileReceiver{
		Cfg:            cfg,
		FileID:         fileHeader.FileID,
		Size:           config.Capacity(fileHeader.Size),
		ChunksNumber:   fileHeader.ChunksNumber,
		ChunksReceived: make(map[uint64]bool),
		StorageIndexes: fileHeader.StorageIndexes,
	}
	for _, storageIndex := range receiver.StorageIndexes {
		if storageIndex >= uint64(len(receiver.Cfg.Storages)) {
			return nil, fmt.Errorf("storage index is out of range for the given config")
		}
		freeStorageCapacity, err := receiver.Cfg.Storages[storageIndex].FreeCapacity()
		if err != nil {
			return nil, fmt.Errorf("error getting free capacity for the storage: %w", err)
		}
		if freeStorageCapacity < receiver.Size {
			return nil, fmt.Errorf("file size is bigger then free capacity of the storage")
		}
	}
	return receiver, nil
}

func (f *FileReceiver) ReceiveNewChunk(chunk api.Chunk) error {
	if f.FileID != chunk.FileID {
		return fmt.Errorf("file id of chunk and FileReceiver don't match")
	}
	for _, storageIndex := range f.StorageIndexes {
		freeStorageCapacity, err := f.Cfg.Storages[storageIndex].FreeCapacity()
		if err != nil {
			return fmt.Errorf("error getting free capacity for the storage: %w", err)
		}
		if freeStorageCapacity.Bytes() < uint64(len(chunk.Payload)) {
			return fmt.Errorf("provided chunk size is bigger then free capacity of the storage")
		}
		err = io.WriteChunk(chunk, f.Cfg.Storages[storageIndex].Path)
		if err != nil {
			return fmt.Errorf("error writing chunk on disk: %w", err)
		}
	}
	f.ChunksReceived[chunk.ChunkIndex] = true
	return nil
}

func (f *FileReceiver) CanBeFinished() bool {
	for i := range f.ChunksNumber {
		if v, ok := f.ChunksReceived[i]; !v || !ok {
			return false
		}
	}
	return true
}

func (f *FileReceiver) Finish() error {
	for _, storageIndex := range f.StorageIndexes {
		err := io.MergeChunks(f.FileID, f.ChunksNumber, f.Cfg.Storages[storageIndex].Path)
		if err != nil {
			return fmt.Errorf("error merging chunks: %w", err)
		}
	}
	return nil
}
