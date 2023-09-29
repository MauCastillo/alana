package analizistrend

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBalanceRealtime(t *testing.T) {
	c := require.New(t)

	analizis := NewAnalizisTrend()

	balance, err := analizis.GetBalanceTrendsRealTime(context.Background(), "EN", "US", "b")
	c.NoError(err)
	c.Equal(0, balance)

}

func TestGetBalanceDaily(t *testing.T) {
	c := require.New(t)

	analizis := NewAnalizisTrend()

	balance, err := analizis.GetBalanceDaily(context.Background(), "EN", "US")
	c.NoError(err)
	c.Equal(0, balance)
}
