package path

import "path/filepath"

func Realname(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return filepath.Base(abs), nil
}
