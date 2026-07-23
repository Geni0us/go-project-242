package code

import (
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExisitingFileSize(t *testing.T) {
	path := "testdata/hello.txt"
	res, err := GetPathSize(path, true, true, true)
	require.Equal(t, nil, err)
	require.Equal(t, "5B", res)
}

func TestNonExistentFileSize(t *testing.T) {
	path := "testdata/unknown.txt"
	res, err := GetPathSize(path, true, true, true)
	require.NotEqual(t, nil, err)
	require.Equal(t, "", res)
}

func TestEmptyFileSize(t *testing.T) {
	path := "testdata/empty.txt"
	res, err := GetPathSize(path, true, true, true)
	require.Equal(t, nil, err)
	require.Equal(t, "0B", res)
}

func TestDirFileSize(t *testing.T) {
	path := "testdata"
	res, err := GetPathSize(path, false, false, true)
	require.Equal(t, nil, err)
	require.Equal(t, "100011B", res)
}

func TestEmptyDirFileSize(t *testing.T) {
	path := "testdata/emptydir"
	mkdirErr := os.MkdirAll(path, os.ModePerm)
	require.Equal(t, nil, mkdirErr)
	res, err := GetPathSize(path, false, false, true)
	require.Equal(t, nil, err)
	require.Equal(t, "0B", res)
}

func TestHiddenCounting(t *testing.T) {
	r1, _ := rawPathSize(".", true, true)
	r2, _ := rawPathSize(".", true, false)

	require.Equal(t, true, r1 > r2)
}

func TestRecursiveWalk(t *testing.T) {
	r1, _ := rawPathSize("testdata", true, false)
	r2, _ := rawPathSize("testdata", false, false)
	require.Equal(t, true, r1 >= r2)
}

func TestSizeHumanize(t *testing.T) {
	var base float64 = 1024
	require.Equal(t, "10B", formatSize(10, true))
	require.Equal(t, "1.0KB", formatSize(int64(base), true))
	require.Equal(t, "1.0MB", formatSize(int64(base*base), true))
	require.Equal(t, "1.2MB", formatSize(1234567, true))
	require.Equal(t, "1.0GB", formatSize(int64(base*base*base), true))
	require.Equal(t, "1.0TB", formatSize(int64(math.Pow(base, 4)), true))
	require.Equal(t, "1.0PB", formatSize(int64(math.Pow(base, 5)), true))
	require.Equal(t, "1.0EB", formatSize(int64(math.Pow(base, 6)), true))
}
