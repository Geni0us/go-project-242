package code

import (
	"math"
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

func TestDirFileSize(t *testing.T) {
	path := "testdata"
	res, err := GetPathSize(path, false, false, true)
	require.Equal(t, nil, err)
	require.Equal(t, "100011B", res)
}

func TestSizeHumanize(t *testing.T) {
	var base float64 = 1024
	require.Equal(t, "10B", FormatSize(10, true))
	require.Equal(t, "1.0KB", FormatSize(int64(base), true))
	require.Equal(t, "1.0MB", FormatSize(int64(base*base), true))
	require.Equal(t, "1.2MB", FormatSize(1234567, true))
	require.Equal(t, "1.0GB", FormatSize(int64(base*base*base), true))
	require.Equal(t, "1.0TB", FormatSize(int64(math.Pow(base, 4)), true))
	require.Equal(t, "1.0PB", FormatSize(int64(math.Pow(base, 5)), true))
	require.Equal(t, "1.0EB", FormatSize(int64(math.Pow(base, 6)), true))
}

func TestHiddenCounting(t *testing.T) {
	r1, _ := RawPathSize(".", false, false)
	r2, _ := RawPathSize(".", false, false)
	require.Equal(t, true, r1 >= r2)
}

func TestRecursiveWalk(t *testing.T) {
	r1, _ := RawPathSize("testdata", true, false)
	r2, _ := RawPathSize("testdata", false, false)
	require.Equal(t, true, r1 >= r2)
}
