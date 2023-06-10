package technicalanalysis

import (
	"testing"

	"github.com/adshao/go-binance/v2"
	"github.com/stretchr/testify/require"
)

func TestCalculateStochasticOscillator(t *testing.T) {
	c := require.New(t)

	tempList := []*binance.Kline{{Close: "50.0"},{Close: "55.0"},{Close: "52.0"},{Close: "48"},{Close: "53"}}


	k := CalculateStochasticOscillator(tempList, 5) 
	c.Equal(float64(71.42857142857143), k)

	tempList = []*binance.Kline{{Close: "50.0"},{Close: "55.0"},{Close: "52.0"},{Close: "48"},{Close: "53"},{Close: "57"},{Close: "60"},{Close: "58"},{Close: "62"},{Close: "59"},{Close: "55"},{Close: "50"},{Close: "48"},{Close: "45"}}
	


	k = CalculateStochasticOscillator(tempList, 14) 
	c.Equal(float64(0), k)
}


func TestCalculateRSI(t *testing.T) {
	c := require.New(t)

	tempList := []*binance.Kline{{Close: "50.0"},{Close: "55.0"},{Close: "52.0"},{Close: "48"},{Close: "53"}}


	rsi := CalculateRSI(tempList) 
	c.Equal(float64(58.82352941176471), rsi)

	tempList = []*binance.Kline{{Close: "50.0"},{Close: "55.0"},{Close: "52.0"},{Close: "48"},{Close: "53"},{Close: "57"},{Close: "60"},{Close: "58"},{Close: "62"},{Close: "59"},{Close: "55"},{Close: "50"},{Close: "48"},{Close: "45"}}
	


	rsi = CalculateRSI(tempList) 
	c.Equal(float64(56.375838926174495), rsi)
}
