package uikline

import (
	"testing"

	"github.com/alana/binance-api/coins"
	"github.com/alana/binance-api/intervals"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	c := require.New(t)

	_, err := NewUIKline(coins.EthBusd, interval.ThirtyMinutes, 2000)
	c.EqualError(ErrLimitGreater, err.Error())

	_, err = NewUIKline(coins.EthBusd, interval.ThirtyMinutes, -2)
	c.EqualError(ErrLimitLess, err.Error())

	uiKline, err := NewUIKline(coins.EthBusd, interval.ThirtyMinutes, 2)
	c.NoError(err)

	c.Equal("", uiKline.Body)

}


