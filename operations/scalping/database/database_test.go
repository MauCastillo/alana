package database

import (
	"database/sql"
	"fmt"
	"os"
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

	defer os.Remove(file)

	coin := symbols.BtcBusd
	simulation, err := simultor.NewSimulator(coin, *intervals.Minute, 60)
	c.NoError(err)

	err = SavewareHouse(simulation, float64(123), tableName)
	c.NoError(err)

	db, err := sql.Open("sqlite3", file)
	c.NoError(err)

	querySelect := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := db.Query(querySelect)
	c.NoError(err)

	var uid int
	var operation string
	var goodPrice float64

	rows.Next()
	err = rows.Scan(&uid, &operation, &goodPrice)
	c.NoError(err)

	c.Equal(goodPrice, float64(123))

}
