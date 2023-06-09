package convertions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringToFloat64(t *testing.T) {
	c := require.New(t)

	f:= StringToFloat64("0.231")
	c.Equal(f, float64(0.231))

	f = StringToFloat64("cat")
	c.Equal(f, float64(0))
}

func TestStringToInt64(t *testing.T) {
	c := require.New(t)

	f:= StringToInt64("7458907890")
	c.Equal(f, int(7458907890))

	f= StringToInt64("cat")
	c.Equal(f, 0)
}

