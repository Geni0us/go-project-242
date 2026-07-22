package path

import (
	"fmt"
	"os"
	"strings"
	extPath "code/src/pkg/path"
)

func Excluded(entryName string, all bool) bool {
	return !all && strings.HasPrefix(entryName, ".")
}

func GetEntry(path string, all bool) (os.FileInfo, error) {
	entry, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	fileName, err := extPath.Realname(path)
	if err != nil {
		return nil, err
	}

	if Excluded(fileName, all) {
		return nil, fmt.Errorf("hidden file ignored")
	}

	return entry, nil
}
