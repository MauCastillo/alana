package analizistrend

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBalanceRealtime(t *testing.T) {
	c := require.New(t)

	analizis, err := NewAnalizisTrend(context.Background(), "EN", "US", "b")
	c.NoError(err)

	err = analizis.Refresh(context.Background())
	c.NoError(err)

	balance, err := analizis.GetBalanceTrendsRealTime(context.Background(), "EN", "US", "b")
	c.NoError(err)
	c.NotEqual(0, balance.Economic)

}

func TestGetBalanceDaily(t *testing.T) {
	c := require.New(t)

	analizis, err := NewAnalizisTrend(context.Background(), "EN", "US", "b")
	c.NoError(err)

	balance, err := analizis.GetBalanceDaily(context.Background(), "EN", "US")
	c.NoError(err)
	c.NotEqual(0, balance.Economic)
}
