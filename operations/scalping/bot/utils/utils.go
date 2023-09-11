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
	Good            = 0
	Mistakes        = 0
	Neutral         = 0
)

const (
	MinuteInSeconds = 60
)

type Util struct {
	Accuracy int             `json:"accuracy"`
	Mistakes int             `json:"mistakes"`
	Neutral  int             `json:"neutral"`
	Earn     float64         `json:"earn"`
	Coin     symbols.Symbols `json:"coin"`
}

func Iterractor(coin *symbols.Symbols, limitKline int) (*simultor.Simulator, error) {
	simulation, err := simultor.NewSimulator(coin, *intervals.Minute, limitKline)
	if err != nil {
		return nil, err
	}

	currentPrice := simulation.CurrentPrice()
	fmt.Printf("=> Activo: %s \nPrecio: %s\n", currentPrice.Symbol, currentPrice.Price)
	fmt.Printf("=> Stochastic Oscillator K: %f \n=> Stochastic Oscillator D: %f \nRelative Strenght Index: %f\n", simulation.StochasticOscillatorK, simulation.StochasticOscillatorD, simulation.RelativeStrenghtIndex)

	simulation.SetPriceBuy(convertions.StringToFloat64(currentPrice.Price))
	return simulation, err
}

func GetBestValue(s *simultor.Simulator, coin *symbols.Symbols, limitKline int) (float64, error) {
	simulation, err := simultor.NewSimulator(coin, *intervals.Minute, limitKline)
	if err != nil {
		Mistakes++
		return float64(0), nil
	}

	Good++

	objetivePrice := 0.0
	if simulation.ObjectivePrice() > s.GetPriceBuy() {
		objetivePrice = simulation.ObjectivePrice()
	}

	err = database.SavewareHouse(coin, s, objetivePrice, simulation.BestPriceCoin())
	if err != nil {
		return float64(0), nil
	}
	fmt.Println(coin.Name)
	fmt.Println("============================= Guardando en Dynamodb ==============================")

	return simulation.ObjectivePrice(), nil
}

func countdown(minute int) {
	second := minute * MinuteInSeconds
	for i := second; i >= 0; i-- {
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

		if best > 0 {
			earn += best - simulation.GetPriceBuy()
		}

	}

	accuracy := (100 / cycles) * Good
	neutral := (100 / cycles) * Neutral
	mistakes := (100 / cycles) * Mistakes

	util := &Util{
		Accuracy: accuracy,
		Mistakes: mistakes,
		Neutral:  neutral,
		Earn:     earn,
		Coin:     *coin,
	}

	return util, nil
}
