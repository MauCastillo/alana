package main

import (
	"fmt"
	"time"

	"github.com/MauCastillo/alana/binance-api/intervals"
	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/operations/scalping/simultor"
	"github.com/MauCastillo/alana/shared/convertions"
	"github.com/MauCastillo/alana/shared/env"
)

var (
	limitKline    = int(env.GetInt64("LIMIT_KLINE", 500))
	waitingPeriod = int(env.GetInt64("WAITING", 2))
	cycles        = int(env.GetInt64("Cycles", 15))
	ganancia      = float64(0)
	PriceBuy      = float64(9999999999)
	Good          = 0
	Neutral       = 0
	Cycle         = 1
)

func Iterractor() error {
	simulation, err := simultor.NewSimulator(*symbols.BtcBusd, *intervals.Minute, limitKline)
	if err != nil {
		return err
	}

	currentPrice := simulation.CurrentPrice()
	fmt.Printf("=> Activo: %s \nPrecio: %s\n", currentPrice.Symbol, currentPrice.Price)
	fmt.Printf("=> Stochastic Oscillator: %f \nRelative Strenght Index: %f\n", simulation.StochasticOscillator, simulation.RelativeStrenghtIndex)

	if simulation.IsTOBuy() {
		fmt.Println("**********************************************")
		fmt.Println("Crear Orden de Compra")
		fmt.Printf("Stochastic Oscillator: %f \nRelative Strenght Index: %f\n", simulation.StochasticOscillator, simulation.RelativeStrenghtIndex)
		PriceBuy = convertions.StringToFloat64(currentPrice.Price)
	}

	if simulation.IsTOSale() {
		fmt.Println("------------------------------------------------")
		fmt.Println("No Comprar ni por el Putas")

		fmt.Printf("Stochastic Oscillator: %f \n Relative Strenght Index: %f\n", simulation.StochasticOscillator, simulation.RelativeStrenghtIndex)
	}

	if !simulation.IsTOBuy() && !simulation.IsTOSale() {
		Neutral++
		fmt.Println("================================================")
		fmt.Println("Simplemente No Se que hacer Necesito Mas Data XD")
		fmt.Printf("Stochastic Oscillator: %f \nRelative Strenght Index: %f\n", simulation.StochasticOscillator, simulation.RelativeStrenghtIndex)
		PriceBuy = convertions.StringToFloat64(currentPrice.Price)
	}

	return nil
}

func GetBestValue() (float64, error) {
	simulation, err := simultor.NewSimulator(*symbols.BtcBusd, *intervals.Minute, waitingPeriod)
	if err != nil {
		return float64(0), err
	}

	if simulation.ObjectivePrice() < PriceBuy {
		return float64(0), nil
	}

	Good++
	return simulation.ObjectivePrice(), nil
}

func countdown(minute int) {
	second := minute * 60
	for i := second; i >= 0; i-- {
		fmt.Printf("[Awating: %d Second] \n", i)
		time.Sleep(time.Second)
	}
}

func main() {
	var err error
	var accuracy int
	var neutral int
	var best float64

	for i := 0; i < cycles; i++ {
		accuracy = (100 / Cycle) * Good
		neutral = (100 / Cycle) * Neutral
		fmt.Printf("===> Accuracy: %d%% \n", accuracy)
		fmt.Printf("===> Neutral: %d%% \n", neutral)
		fmt.Printf("===> Index: %d \n", i)
		err = Iterractor()
		if err != nil {
			fmt.Println(err)
			continue
		}

		countdown(waitingPeriod)
		best, err := GetBestValue()
		if err != nil {
			fmt.Println(err)
			continue
		}

		Cycle++

		fmt.Printf("Best Value to Buy: %f\n Best Value to Sale: %f", PriceBuy, best)
		if best > 0 {
			ganancia += best - PriceBuy
		}

		fmt.Printf("===> Current is await Ganancias USD:%f", ganancia)
		fmt.Printf("===> Current is await Ganancias REAL:%f", ganancia/1000)

	}

	fmt.Println("=> This is the end! <=")
	accuracy = (100 / Cycle) * Good
	neutral = (100 / Cycle) * Neutral
	fmt.Printf(":::: Accuracy: %d%% ::::\n", accuracy)
	fmt.Printf(":::: Neutral: %d%% ::::\n", neutral)

	fmt.Printf("Best Value to Buy: %f\n Best Value to Sale: %f", PriceBuy, best)
	fmt.Printf("===> Current is await")
}
