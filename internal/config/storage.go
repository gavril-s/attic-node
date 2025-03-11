package config

type Storage struct {
	Path     Path
	Capacity Capacity
}

func NewStorage(storage JSONStorage) (Storage, error) {
	capacity, err := ParseCapacity(storage.Capacity)
	if err != nil {
		return Storage{}, err
	}
	return Storage{
		Path:     Path(storage.Path),
		Capacity: capacity,
	}, nil
}

func (s Storage) FreeCapacity() (Capacity, error) {
	pathExists, err := s.Path.Exists()
	if err != nil {
		return 0, err
	}
	if !pathExists {
		return 0, nil
	}
	return s.Path.Size()
}
