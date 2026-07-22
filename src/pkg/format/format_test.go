package format

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

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
