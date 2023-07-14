package technicalanalysis

import (
	"github.com/MauCastillo/alana/shared/convertions"
	"github.com/adshao/go-binance/v2"
)

func CalculateStochasticOscillator(kline []*binance.Kline, period int) (float64, float64) {
	// Verificar si hay suficientes precios para el período
	if len(kline) < period {
		return 0.0, 0.0
	}

	// Calcular el mínimo y máximo del período seleccionado
	minPrice := minFloat64(kline)
	maxPrice := maxFloat64(kline)
	currentPrice := convertions.StringToFloat64(kline[len(kline)-1].Close)

	// Calcular el valor del %K

	k := ((currentPrice - minPrice) / (maxPrice - minPrice)) * 100

	// Paso 2: Calcular el Valor %D (Media Móvil de %K)
	dValue := simpleMovingAverage(kline, period)

	return k, dValue
}

// Función para encontrar el máximo valor en un slice de float64
func maxFloat64(numbers []*binance.Kline) float64 {
	max := convertions.StringToFloat64(numbers[0].Close)
	var numValue float64

	for _, value := range numbers {
		numValue = convertions.StringToFloat64(value.Close)
		if numValue > max {
			max = numValue
		}
	}
	return max
}

// Función para encontrar el mínimo valor en un slice de float64
func minFloat64(numbers []*binance.Kline) float64 {
	min := convertions.StringToFloat64(numbers[0].Close)
	var numValue float64

	for _, value := range numbers {
		numValue = convertions.StringToFloat64(value.Close)
		if numValue < min {
			min = numValue
		}
	}
	return min
}

// Función para calcular la media móvil simple de un slice de float64
func simpleMovingAverage(numbers []*binance.Kline, period int) float64 {
	sum := 0.0
	count := 0
	for i := len(numbers) - period; i < len(numbers); i++ {
		sum += convertions.StringToFloat64(numbers[i].Close)
		count++
	}
	return sum / float64(count)
}

func CalculateRSI(data []*binance.Kline) float64 {
	var firstPoint float64
	var secondPointPoint float64

	gains := make([]float64, 0)
	losses := make([]float64, 0)

	// Calculate gains and losses
	for i := 1; i < len(data); i++ {
		firstPoint = convertions.StringToFloat64(data[i].Close)
		secondPointPoint = convertions.StringToFloat64(data[i-1].Close)
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
	if sum <= float64(0) {
		return float64(0.0001)
	}

	average := sum / float64(len(data))
	return average
}
