package main

import (
	"fmt"

	"github.com/MauCastillo/alana/binance-api/intervals"
	"github.com/MauCastillo/alana/binance-api/services"
	"github.com/MauCastillo/alana/binance-api/symbols"
	technicalanalysis "github.com/MauCastillo/alana/binance-api/technical-analysis"
	"github.com/MauCastillo/alana/shared/env"
)

var (
	limitKline  = int(env.GetInt64("LIMIT_KLINE", 4))
	limitRSIBuy = env.GetFloat64("LIMIT_RSI_BUY", 30)
	limitOSBuy  = env.GetFloat64("LIMIT_OS_BUY", 20)

	limitRSISale = env.GetFloat64("LIMIT_RSI_SALE", 30)
	limitOSSale  = env.GetFloat64("LIMIT_OS_SALE", 20)
)

func main() {
	LocalKlineToBTC, err := services.NewKlineService(*symbols.BtcBusd, *intervals.FifteenMinutes, limitKline)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	stochasticOscillator := technicalanalysis.CalculateStochasticOscillator(LocalKlineToBTC.Kline, 4)
	relativeStrenghtIndex := technicalanalysis.CalculateRSI(LocalKlineToBTC.Kline)

	if stochasticOscillator <= limitOSBuy && relativeStrenghtIndex <= limitRSIBuy {
		fmt.Println("_________________________________________________")
		fmt.Println("Crear Orden de Compra")
		fmt.Printf("Stochastic Oscillator: %f \n Relative Strenght Index: %f", stochasticOscillator, relativeStrenghtIndex)

	}

	if stochasticOscillator >= limitOSSale && relativeStrenghtIndex >= limitRSISale {
		fmt.Println("_________________________________________________")
		fmt.Println("Crear Orden de Venta")
		fmt.Printf("Stochastic Oscillator: %f \n Relative Strenght Index: %f", stochasticOscillator, relativeStrenghtIndex)

	}

}
