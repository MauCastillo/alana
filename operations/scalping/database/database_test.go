package database

import (
	"testing"

	"github.com/MauCastillo/alana/binance-api/intervals"
	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/operations/scalping/simultor"
	"github.com/stretchr/testify/require"
)

func TestSavewareHouse(t *testing.T) {
	c := require.New(t)

	const file string = "data-warehouse.sqlite3"
	const tableName = "basic_training"

	coin := symbols.BtcBusd
	simulation, err := simultor.NewSimulator(coin, *intervals.Minute, 60)
	c.NoError(err)

	err = SavewareHouse(symbols.AdaUsdt, simulation, float64(123), float64(33))
	c.NoError(err)
}
