package code

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExisitingFileSize(t *testing.T) {
	path := "testdata/hello.txt"
	res, err := GetPathSize(path, true, true, true)
	require.Equal(t, err, nil)
	require.Equal(t, res, fmt.Sprintf("5B\t%s", path))
}

func TestNonExistentFileSize(t *testing.T) {
	path := "testdata/unknown.txt"
	res, err := GetPathSize(path, true, true, true)
	require.NotEqual(t, err, nil)
	require.Equal(t, res, "")
}

func TestDirFileSize(t *testing.T) {
	path := "testdata"
	res, err := GetPathSize(path, false, false, true)
	require.Equal(t, err, nil)
	require.Equal(t, res, fmt.Sprintf("100683B\t%s", path))
}

func TestSizeHumanize(t *testing.T) {
	var base float64 = 1024
	require.Equal(t, FormatSize(10, true), "10B")
	require.Equal(t, FormatSize(int64(base), true), "1KB")
	require.Equal(t, FormatSize(int64(math.Pow(base, 2)), true), "1MB")
	require.Equal(t, FormatSize(1234567, true), "1.2MB")
	require.Equal(t, FormatSize(int64(math.Pow(base, 3)), true), "1GB")
	require.Equal(t, FormatSize(int64(math.Pow(base, 4)), true), "1TB")
	require.Equal(t, FormatSize(int64(math.Pow(base, 5)), true), "1PB")
	require.Equal(t, FormatSize(int64(math.Pow(base, 6)), true), "1EB")
}
