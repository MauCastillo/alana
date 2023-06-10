package main

import (
	"fmt"

	"github.com/MauCastillo/alana/binance-api/intervals"
	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/shared/env"
	"github.com/MauCastillo/alana/operations/scalping/simultor"
)

var (
	limitKline  = int(env.GetInt64("LIMIT_KLINE", 500))
)



func main() {
	simulation, err := simultor.NewSimulator(*symbols.BtcBusd, *intervals.Minute, limitKline)
	if err != nil{
		fmt.Print(err)
	}

	currentPrice := simulation.CurrentPrice()
	fmt.Printf("=> Activo: %s \nPrecio: %s\n",currentPrice.Symbol, currentPrice.Price)
	fmt.Printf("=> Stochastic Oscillator: %f \nRelative Strenght Index: %f\n", simulation.StochasticOscillator, simulation.RelativeStrenghtIndex)

	if simulation.IsTOBuy() {
		fmt.Println("_________________________________________________")
		fmt.Println("Crear Orden de Compra")
		fmt.Printf("Stochastic Oscillator: %f \nRelative Strenght Index: %f\n", simulation.StochasticOscillator, simulation.RelativeStrenghtIndex)

	}

	if simulation.IsTOSale() {
		fmt.Println("_________________________________________________")
		fmt.Println("Crear Orden de Venta")
		fmt.Printf("Stochastic Oscillator: %f \n Relative Strenght Index: %f\n", simulation.StochasticOscillator, simulation.RelativeStrenghtIndex)
	}

}
