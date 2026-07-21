package code

import (
	"fmt"
	"math"
	"os"
	syspath "path"
	"path/filepath"
	"strconv"
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
		unit = units[int(pow)]
		resultSize = math.Round((resultSize/math.Pow(devider, pow))*10) / 10
	}
	return fmt.Sprintf("%s%s", strconv.FormatFloat(resultSize, 'f', -1, 64), unit)
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
		if entry.IsDir() {
			if recurcive && (all || !isHidden(entry.Name())) {
				size += scanDir(syspath.Join(path, entry.Name()), all, recurcive)
			}
		} else {
			f, err := entry.Info()
			if err == nil && all || !isHidden(entry.Name()) {
				size += f.Size()
			}
		}
	}
	return size
}

func GetPathSize(path string, recurcive, human, all bool) (string, error) {

	entry, err := os.Stat(path)

	if err != nil {
		return "", err
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	if !all && isHidden(filepath.Base(abs)) {
		return "", fmt.Errorf("hidden file ignored")
	}

	var size int64
	if entry.IsDir() {
		size = scanDir(path, all, recurcive)
	} else {
		size = entry.Size()
	}

	return fmt.Sprintf("%s\t%s", FormatSize(size, human), path), nil
}
