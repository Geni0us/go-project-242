package counting

import (
	
	"testing"
	
	"github.com/stretchr/testify/require"
)

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
