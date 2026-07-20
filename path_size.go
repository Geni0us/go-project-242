package code

import (
	"fmt"
	"os"
	"strconv"
)

func GetPathSize(path string, _, _, _ bool) (string, error) {

	inf, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	var size int64 = 0
	if inf.IsDir() {
		files, err := os.ReadDir(path)
		if err != nil {
			return "", err
		}
		for _, entry := range files {
			if entry.IsDir() {
				// TODO: recurcive walk
			} else {
				eInf, err := entry.Info()
				if err != nil {
					continue
				}
				size += eInf.Size()
			}
		}
	} else {
		size = inf.Size()
	}

	return fmt.Sprintf("%sB\t%s", strconv.FormatInt(size, 10), path), nil
}
