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
	limitKline    = int(env.GetInt64("LIMIT_KLINE", 100))
	waitingPeriod = int(env.GetInt64("WAITING", 15))
	cycles        = int(env.GetInt64("Cycles", 15))
	PriceBuy      = float64(9999999999)
	Good          = 0
	Cycle         = 0
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
		fmt.Println("_________________________________________________")
		fmt.Println("Crear Orden de Compra")
		fmt.Printf("Stochastic Oscillator: %f \nRelative Strenght Index: %f\n", simulation.StochasticOscillator, simulation.RelativeStrenghtIndex)
		PriceBuy = convertions.StringToFloat64(currentPrice.Price)
	}

	if simulation.IsTOSale() {
		fmt.Println("_________________________________________________")
		fmt.Println("No Comprar ni por el Putas")

		fmt.Printf("Stochastic Oscillator: %f \n Relative Strenght Index: %f\n", simulation.StochasticOscillator, simulation.RelativeStrenghtIndex)
	}

	best, err := GetBestValueo()
	if err != nil {
		return err
	}

	fmt.Printf("Best Value to Buy: %f\n Best Value to Sale: %f", simulation.StochasticOscillator, best)

	return nil
}

func GetBestValueo() (float64, error) {
	sleepTimer := 15 * time.Minute
	time.Sleep(sleepTimer)

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

func main() {
	var err error
	for i := 0; i < cycles; i++ {

		fmt.Printf("===> Index: %d \n", i)
		err = Iterractor()
		if err != nil {
			fmt.Println(err)
			continue
		}

		sleepTimer := 15 * time.Minute
		time.Sleep(sleepTimer)

		accuracy := (100/Cycle) * Good
		fmt.Printf("===> accuracy: %d \n", accuracy)
	}
}
