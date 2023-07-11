package utils

import (
	"fmt"
	"time"

	"github.com/MauCastillo/alana/binance-api/intervals"
	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/operations/scalping/database"
	"github.com/MauCastillo/alana/operations/scalping/simultor"
	"github.com/MauCastillo/alana/shared/convertions"
)

var (
	Good     = 0
	Mistakes = 0
	Neutral  = 0
	Cycle    = 1
)

type Util struct {
	Accuracy int     `json:"accuracy"`
	Mistakes int     `json:"mistakes"`
	Neutral  int     `json:"neutral"`
	Earn     float64 `json:"earn"`
}

func Iterractor(coin *symbols.Symbols, limitKline int) (*simultor.Simulator, error) {
	simulation, err := simultor.NewSimulator(coin, *intervals.Minute, limitKline)
	if err != nil {
		return nil, err
	}

	currentPrice := simulation.CurrentPrice()
	fmt.Printf("=> Activo: %s \nPrecio: %s\n", currentPrice.Symbol, currentPrice.Price)
	fmt.Printf("=> Stochastic Oscillator: %f \nRelative Strenght Index: %f\n", simulation.StochasticOscillator, simulation.RelativeStrenghtIndex)

	if simulation.IsTOBuy() {
		fmt.Println("**********************************************")
		fmt.Println("Crear Orden de Compra")
		fmt.Printf("Stochastic Oscillator: %f \nRelative Strenght Index: %f\n", simulation.StochasticOscillator, simulation.RelativeStrenghtIndex)
		simulation.SetPriceBuy(convertions.StringToFloat64(currentPrice.Price))
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
		simulation.SetPriceBuy(convertions.StringToFloat64(currentPrice.Price))
	}

	return simulation, err
}

func GetBestValue(s *simultor.Simulator, coin *symbols.Symbols, limitKline int) (float64, error) {
	simulation, err := simultor.NewSimulator(coin, *intervals.Minute, limitKline)
	if err != nil || simulation.ObjectivePrice() < s.GetPriceBuy() {
		err := database.SavewareHouse(s, float64(0))
		if err != nil {
			return float64(0), nil
		}

		Mistakes++
		return float64(0), nil
	}

	Good++

	err = database.SavewareHouse(s, simulation.ObjectivePrice())
	if err != nil {
		return float64(0), nil
	}

	return simulation.ObjectivePrice(), nil
}

func countdown(minute int) {
	second := minute * 60
	for i := second; i >= 0; i-- {
		fmt.Printf("[Awating: %d Second] \n", i)
		time.Sleep(time.Second)
	}
}

func RunCollector(coin *symbols.Symbols, limitKline, waitingPeriod, cycles, periodSell int) (*Util, error) {
	earn := float64(0)
	for i := 0; i < cycles; i++ {
		simulation, err := Iterractor(coin, limitKline)
		if err != nil {
			return nil, err
		}

		countdown(waitingPeriod)
		best, err := GetBestValue(simulation, coin, periodSell)
		if err != nil {
			return nil, err
		}

		Cycle++
		if best > 0 {
			earn += best - simulation.GetPriceBuy()
		}

	}

	accuracy := (100 / Cycle) * Good
	neutral := (100 / Cycle) * Neutral
	mistakes := (100 / Cycle) * Mistakes

	util := &Util{
		Accuracy: accuracy,
		Mistakes: neutral,
		Neutral:  mistakes,
		Earn:     earn,
	}

	return util, nil
}
