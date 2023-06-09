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


