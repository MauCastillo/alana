package match


import (
	"sort"

	"github.com/adshao/go-binance/v2"
	"github.com/MauCastillo/alana/shared/convertions"
)

// CalculateAverage return the Average
func CalculateAverage(klines []*binance.Kline) float64 {
	data := ExtractClose(klines)
	sum := 0.0
	for _, value := range data {
		sum += value
	}

	mean := sum / float64(len(data))
	return mean
}

// CalculateMedian return the Median
func CalculateMedian(klines []*binance.Kline) float64 {
	data := ExtractClose(klines)
	sort.Float64s(data)
	length := len(data)
	if length%2 == 0 {
		median := (data[length/2-1] + data[length/2]) / 2
		return median
	}
	median := data[length/2]
	return median
}

// CalculateAverage return the Average using low price
func CalculateAverageLowPrice(klines []*binance.Kline) float64 {
	data := ExtractLow(klines)
	sum := 0.0
	for _, value := range data {
		sum += value
	}

	mean := sum / float64(len(data))
	return mean
}

// CalculateMedian return the Median using low price
func CalculateMedianLowPrice(klines []*binance.Kline) float64 {
	data := ExtractLow(klines)
	sort.Float64s(data)
	length := len(data)
	if length%2 == 0 {
		median := (data[length/2-1] + data[length/2]) / 2
		return median
	}
	median := data[length/2]
	return median
}

// ExtractClose return all close price like float64 array
func ExtractClose(array []*binance.Kline) []float64 {
    result := make([]float64, len(array))

    for i, element := range array {
        result[i] = convertions.StringToFloat64(element.Close)
    }

    return result
}

// ExtractLow return all low price values like float64 array
func ExtractLow(array []*binance.Kline) []float64 {
    result := make([]float64, len(array))

    for i, element := range array {
        result[i] = convertions.StringToFloat64(element.Low)
    }

    return result
}