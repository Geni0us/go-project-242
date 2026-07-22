package counting

import (
	internalPath "code/src/internal/path"
	"os"
	fsPath "path"
)

func RawPathSize(path string, recurcive, all bool) (int64, error) {
	entry, err := internalPath.ValidatePath(path, all)
	if err != nil {
		return 0, err
	}

	if entry.IsDir() {
		return dirSize(path, all, recurcive), nil
	}
	return entry.Size(), nil
}

func dirSize(path string, all, recurcive bool) int64 {
	files, err := os.ReadDir(path)
	var size int64
	if err != nil {
		return size
	}
	for _, entry := range files {
		if internalPath.Excluded(entry.Name(), all) {
			continue
		}

		if entry.IsDir() {
			if recurcive {
				size += dirSize(fsPath.Join(path, entry.Name()), all, recurcive)
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
