package code

import (
	"fmt"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"strings"
)

func excluded(entryName string, all bool) bool {
	return !all && strings.HasPrefix(entryName, ".")
}

func realname(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return filepath.Base(abs), nil
}

func getEntry(path string, all bool) (os.FileInfo, error) {
	entry, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	fileName, err := realname(path)
	if err != nil {
		return nil, err
	}

	if excluded(fileName, all) {
		return nil, fmt.Errorf("hidden file ignored")
	}

	return entry, nil
}

func dirSize(_path string, all, recurcive bool) (int64, error) {
	var size int64
	err := filepath.Walk(_path, func(path string, info fs.FileInfo, err error) error {
	
		_excluded := excluded(info.Name(), all)
		if err != nil || path == _path {
			return err
		}
		if info.IsDir() {
			if _excluded || !recurcive {
				return filepath.SkipDir
			}
			return nil
		}
		if _excluded {
			return nil
		}

		size += info.Size()
		return nil
	})
	if err != nil {
		return 0, err
	}
	return size, nil
}

func rawPathSize(path string, recurcive, all bool) (int64, error) {
	entry, err := getEntry(path, all)
	if err != nil {
		return 0, err
	}

	if entry.IsDir() {
		return dirSize(path, all, recurcive)
	}
	return entry.Size(), nil
}

var units = map[int]string{
	0: "B",
	1: "KB",
	2: "MB",
	3: "GB",
	4: "TB",
	5: "PB",
	6: "EB",
}

func formatSize(size int64, humanize bool) string {
	resultSize := float64(size)
	unit := units[0]

	if humanize && size > 0 {
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

// GetPathSize выводит размер указанного файла или директории в формате "<размер entity><Единицы измерения>"
// Рассчитанные значения размера могут меняться в зависимости от использования флагов all и recurcive
// Флаг recurcive указывает на то будут ли учитываться вложенные директории
// Флаг all  указывает на то будут ли учитываться скрытые файлы и директории (название начинается с '.')
// Флаг human отвечает за то будет вывод приведён к человекопонятным единицам.
// Eсли human=false то вывод в байтах, иначе автоматически определяется исходя из размера файла. Пример 1234567B => 1.2MB
func GetPathSize(path string, recurcive, human, all bool) (string, error) {

	size, err := rawPathSize(path, recurcive, all)
	if err != nil {
		return "", err
	}

	return formatSize(size, human), nil
}
