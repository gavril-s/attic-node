package domain

import "os"

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
