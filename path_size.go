package code

import (
	"code/src/pkg/counting"
	"code/src/pkg/format"
)

func GetPathSize(path string, recurcive, human, all bool) (string, error) {

	size, err := counting.RawPathSize(path, recurcive, all)
	if err != nil {
		return "", err
	}

	return format.FormatSize(size, human), nil
}
