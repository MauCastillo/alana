package cnn

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func TestGet(t *testing.T) {
	c := require.New(t)

	request, err := NewFearAndGreedCNN()
	c.NoError(err)

	req := request.Get()

	c.True(req.FearAndGreed.Score > 0)

}
