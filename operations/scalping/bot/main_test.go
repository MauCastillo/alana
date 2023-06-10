package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	c := require.New(t)

	main()
	c.Equal(limitKline, 4)
}
