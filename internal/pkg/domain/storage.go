package domain

import "github.com/gavril-s/attic-node/internal/config"

type Storage struct {
	Path     Path
	Capacity Capacity
}

func NewStorage(cfg config.StorageConfig) (Storage, error) {
	capacity, err := ParseCapacity(cfg.Capacity)
	if err != nil {
		return Storage{}, err
	}
	return Storage{
		Path:     Path(cfg.Path),
		Capacity: capacity,
	}, nil
}
