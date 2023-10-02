package utils

import (
	"context"
	"testing"

	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/shared/cnn"
	"github.com/MauCastillo/alana/shared/google/analizistrend"
	"github.com/stretchr/testify/require"
)

func TestRunCollector(t *testing.T) {
	c := require.New(t)

	limitKline := 60
	waitingPeriod := 0
	periodSell := 5
	cycles := 1
	coin := symbols.BtcBusd

	requestCNN, err := cnn.NewFearAndGreedCNN()
	c.NoError(err)

	analizis, err := analizistrend.NewAnalizisTrend(context.Background(), "EN", "US", "b")
	c.NoError(err)

	u, err := RunCollector(coin, limitKline, waitingPeriod, cycles, periodSell, analizis, requestCNN)
	c.Equal(u.Coin.Value, symbols.BtcBusd.Value)
	c.NoError(err)
}
