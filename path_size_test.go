package code

import (
	"fmt"
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
	res, err := GetPathSize(path, false, true, true)
	require.Equal(t, err, nil)
	require.Equal(t, res, fmt.Sprintf("11B\t%s", path))
}
