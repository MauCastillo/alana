package technicalanalysis

import (
	"github.com/adshao/go-binance/v2"
	"github.com/MauCastillo/alana/shared/convertions"
	
)

func CalculateStochasticOscillator(kline []*binance.Kline, period int) float64 {
	var dayPrice float64
	// Verificar si hay suficientes precios para el período
	if len(kline) < period {
		return 0.0
	}

	// Calcular el mínimo y máximo del período seleccionado
	minPrice := convertions.StringToFloat64(kline[0].Close)
	maxPrice := convertions.StringToFloat64(kline[0].Close)
	for i := 1; i < period; i++ {
		dayPrice = convertions.StringToFloat64(kline[i].Close)
		if dayPrice < minPrice {
			minPrice = dayPrice
		}
		if dayPrice > maxPrice {
			maxPrice = dayPrice
		}
	}

	// Calcular el valor del %K
	currentPrice := convertions.StringToFloat64(kline[period-1].Close)
	k := (currentPrice - minPrice) / (maxPrice - minPrice) 

	return k * 100
}

func CalculateRSI(data []*binance.Kline) float64 {
	var firstPoint float64
	var secondPointPoint float64

	gains := make([]float64, 0)
	losses := make([]float64, 0)


	// Calculate gains and losses
	for i := 1; i < len(data); i++ {
		firstPoint  = convertions.StringToFloat64(data[i].Close)
		secondPointPoint  = convertions.StringToFloat64(data[i-1].Close) 
		diff := firstPoint - secondPointPoint
		if diff >= 0 {
			gains = append(gains, diff)
		} else {
			losses = append(losses, -diff)
		}
	}

	// Calculate average gains and losses
	avgGain := calculateAverage(gains)
	avgLoss := calculateAverage(losses)

	// Calculate RSI
	rs := avgGain / avgLoss
	rsi := 100 - (100 / (1 + rs))

	return rsi
}

func calculateAverage(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	average := sum / float64(len(data))
	return average
}

