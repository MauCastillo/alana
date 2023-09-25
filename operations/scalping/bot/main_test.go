package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDatabase(t *testing.T) {
	c := require.New(t)

	main()
	c.True(cycles > 0)
}
