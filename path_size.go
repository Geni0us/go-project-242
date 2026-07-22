package code

import (
	"fmt"
	"math"
	"os"
	syspath "path"
	"path/filepath"
	"strings"
)

func FormatSize(size int64, humanize bool) string {
	resultSize := float64(size)
	units := map[int]string{
		0: "B",
		1: "KB",
		2: "MB",
		3: "GB",
		4: "TB",
		5: "PB",
		6: "EB",
	}
	unit := units[0]

	if humanize {
		var devider float64 = 1024

		pow := math.Floor(math.Log(resultSize) / math.Log(devider))
		_unit, ok := units[int(pow)]
		if !ok {
			unit = "MoreThanYouCanImagineBytes"
		} else {
			unit = _unit
		}
		resultSize = math.Round((resultSize/math.Pow(devider, pow))*10) / 10
	}
	if unit == units[0] {
		return fmt.Sprintf("%d%s", int64(resultSize), unit)
	}
	return fmt.Sprintf("%.1f%s", resultSize, unit)
}

func isHidden(filename string) bool {
	return strings.HasPrefix(filename, ".")
}

func scanDir(path string, all, recurcive bool) int64 {
	files, err := os.ReadDir(path)
	var size int64
	if err != nil {
		return size
	}
	for _, entry := range files {
		if !all && isHidden(entry.Name()) {
			continue
		}

		if entry.IsDir() {
			if recurcive {
				size += scanDir(syspath.Join(path, entry.Name()), all, recurcive)
			}
		} else {
			f, err := entry.Info()
			if err == nil {
				size += f.Size()
			}
		}
	}
	return size
}

func realname(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return filepath.Base(abs), nil
}

func ValidatePath(path string, all bool) (os.FileInfo, error) {
	entry, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	fileName, err := realname(path)
	if err != nil {
		return nil, err
	}

	if !all && isHidden(fileName) {
		return nil, fmt.Errorf("hidden file ignored")
	}

	return entry, nil
}

func RawPathSize(path string, recurcive, all bool) (int64, error) {
	entry, err := ValidatePath(path, all)
	if err != nil {
		return 0, err
	}

	if entry.IsDir() {
		return scanDir(path, all, recurcive), nil
	}
	return entry.Size(), nil
}

func GetPathSize(path string, recurcive, human, all bool) (string, error) {
	size, err := RawPathSize(path, recurcive, all)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s\t%s", FormatSize(size, human), path), nil
}
