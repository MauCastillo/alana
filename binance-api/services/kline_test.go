package services

import (
	"testing"

	"github.com/MauCastillo/alana/binance-api/intervals"
	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/stretchr/testify/require"
)

func TestGetNewKlineService(t *testing.T) {
	c := require.New(t)

	klineService, err := NewKlineService(*symbols.BtcBusd, *intervals.FifteenMinutes, 10)
	c.NoError(err)
	c.Len(klineService.Kline, 10)

	line := klineService.MaxValueClose()
	c.True(line.OpenTime > 11)

	price, err := klineService.ListPricesService(symbols.BtcBusd)
	c.NoError(err)
	c.Equal(symbols.BtcBusd.Value, price.Symbol)

	failSymbol := symbols.Symbols{Value: "cat"}
	_, err = NewKlineService(failSymbol, *intervals.FifteenMinutes, 10)
	c.Equal("<APIError> code=-1100, msg=Illegal characters found in parameter 'symbol'; legal range is '^[A-Z0-9-_.]{1,20}$'.", err.Error())
}

func TestGetNewKlineServiceError(t *testing.T) {
	c := require.New(t)

	failSymbol := symbols.Symbols{Value: "cat"}
	_, err := NewKlineService(failSymbol, *intervals.FifteenMinutes, 10)
	c.Equal("<APIError> code=-1100, msg=Illegal characters found in parameter 'symbol'; legal range is '^[A-Z0-9-_.]{1,20}$'.", err.Error())
}
