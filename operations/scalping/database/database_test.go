package database

import (
	"testing"

	"github.com/MauCastillo/alana/binance-api/intervals"
	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/operations/scalping/simultor"
	"github.com/MauCastillo/alana/shared/cnn"
	"github.com/MauCastillo/alana/shared/google/analizistrend"
	"github.com/stretchr/testify/require"
)

func TestSavewareHouse(t *testing.T) {
	c := require.New(t)

	const file string = "data-warehouse.sqlite3"
	const tableName = "basic_training"

	requestCNN, err := cnn.NewFearAndGreedCNN()
	c.NoError(err)

	coin := symbols.BtcBusd
	simulation, err := simultor.NewSimulator(coin, *intervals.Minute, 60, requestCNN)
	c.NoError(err)

	trend := &analizistrend.AnalizisTrend{RealtimeArticleBalance: &analizistrend.Analizis{Economic: 8, Cryptocurrency: 5}}

	err = SavewareHouse(symbols.EthUsdt, simulation, trend, float64(123), float64(33), float64(33))
	c.NoError(err)
}
