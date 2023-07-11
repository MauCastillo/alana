package utils

import (
	"os"
	"testing"

	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/stretchr/testify/require"
)

func TestRunCollector(t *testing.T) {
	c := require.New(t)

	defer cleanup()

	limitKline := 60
	waitingPeriod := 0
	periodSell := 5
	cycles := 1
	coin := symbols.BtcBusd

	u, err := RunCollector(coin, limitKline, waitingPeriod, cycles, periodSell)
	c.Equal(u.Earn, float64(0))
	c.NoError(err)
	c.Equal(limitKline, 60)
}

func cleanup() {
	_ = os.Remove("data-warehouse.sqlite3")
}

