package code

import (
	"fmt"
	"os"
	"strconv"
)

func GetPathSize(path string, _, _, _ bool) (string, error) {
	inf, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%sB\t%s", strconv.FormatInt(inf.Size(), 10), path), nil
}
