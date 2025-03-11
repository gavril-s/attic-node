package config

import (
	"os"
	"path/filepath"
)

type Path string

func (p Path) Exists() (bool, error) {
	if _, err := os.Stat(string(p)); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (p Path) Size() (Capacity, error) {
	var size Capacity

	info, err := os.Stat(string(p))
	if err != nil {
		return 0, err
	}
	if !info.IsDir() {
		return Capacity(info.Size()), nil
	}

	calcSize := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += Capacity(info.Size())
		}
		return nil
	}
	filepath.Walk(string(p), calcSize)

	return size, nil
}

func (p Path) String() string {
	return string(p)
}
