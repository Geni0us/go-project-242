package code

import (
	"fmt"
	"math"
	"os"
	syspath "path"
	"strconv"
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

func GetPathSize(path string, _, human, _ bool) (string, error) {

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
				eInf, err = os.Lstat(syspath.Join(path, eInf.Name()))
				size += eInf.Size()
			}
		}
	} else {
		size = inf.Size()
	}

	return fmt.Sprintf("%s\t%s", FormatSize(size, human), path), nil
}
