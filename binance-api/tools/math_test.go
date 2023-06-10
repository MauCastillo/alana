package match

import (
	"testing"

	"github.com/adshao/go-binance/v2"
	"github.com/stretchr/testify/require"
)

func TestCalculateMedian(t *testing.T) {
	c := require.New(t)

	tempList := []*binance.Kline{{Close: "50.0"}, {Close: "55.0"}, {Close: "52.0"}, {Close: "48"}, {Close: "53"}}
	median := CalculateMedian(tempList)
	c.Equal(float64(52), median)

	tempList = []*binance.Kline{{Close: "50.0"}, {Close: "55.0"}, {Close: "52.0"}, {Close: "48"}}
	median = CalculateMedian(tempList)
	c.Equal(float64(51), median)
}

func TestCalculateAverage(t *testing.T) {
	c := require.New(t)

	tempList := []*binance.Kline{{Close: "50.0"}, {Close: "55.0"}, {Close: "52.0"}, {Close: "48"}, {Close: "53"}}
	avg := CalculateAverage(tempList)
	c.Equal(float64(51.6), avg)

	tempList = []*binance.Kline{{Close: "50.0"}, {Close: "55.0"}, {Close: "52.0"}, {Close: "48"}}
	avg = CalculateAverage(tempList)
	c.Equal(float64(51.25), avg)
}

func TestCalculateMedianLowPrice(t *testing.T) {
	c := require.New(t)

	tempList := []*binance.Kline{{Low: "50.0"}, {Low: "55.0"}, {Low: "52.0"}, {Low: "48"}, {Low: "53"}}
	median := CalculateMedianLowPrice(tempList)
	c.Equal(float64(52), median)

	tempList = []*binance.Kline{{Low: "50.0"}, {Low: "55.0"}, {Low: "52.0"}, {Low: "48"}}
	median = CalculateMedianLowPrice(tempList)
	c.Equal(float64(51), median)
}

func TestCalculateAverageLowPrice(t *testing.T) {
	c := require.New(t)

	tempList := []*binance.Kline{{Low: "50.0"}, {Low: "55.0"}, {Low: "52.0"}, {Low: "48"}, {Low: "53"}}
	avg := CalculateAverageLowPrice(tempList)
	c.Equal(float64(51.6), avg)

	tempList = []*binance.Kline{{Low: "50.0"}, {Low: "55.0"}, {Low: "52.0"}, {Low: "48"}}
	avg = CalculateAverageLowPrice(tempList)
	c.Equal(float64(51.25), avg)
}
