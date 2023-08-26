package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDatabase(t *testing.T) {
	c := require.New(t)

	defer os.Remove("data-warehouse.sqlite3")
	main()
	c.True(cycles > 0)
}
