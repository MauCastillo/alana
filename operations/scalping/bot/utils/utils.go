package utils

import (
	"fmt"
	"time"

	"github.com/MauCastillo/alana/binance-api/intervals"
	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/operations/scalping/database"
	"github.com/MauCastillo/alana/operations/scalping/simultor"
	"github.com/MauCastillo/alana/shared/cnn"
	"github.com/MauCastillo/alana/shared/convertions"
	"github.com/MauCastillo/alana/shared/google/analizistrend"
)

const (
	MinuteInSeconds = 60
)

type Util struct {
	Earn float64         `json:"earn"`
	Coin symbols.Symbols `json:"coin"`
}

func Iterractor(coin *symbols.Symbols, limitKline int, cnnReport *cnn.FearAndGreedCNN) (*simultor.Simulator, error) {
	simulation, err := simultor.NewSimulator(coin, *intervals.Minute, limitKline, cnnReport)
	if err != nil {
		return nil, err
	}

	currentPrice := simulation.CurrentPrice()
	fmt.Printf("=> Activo: %s \nPrecio: %s\n", currentPrice.Symbol, currentPrice.Price)
	fmt.Printf("=> Stochastic Oscillator K: %f \n=> Stochastic Oscillator D: %f \nRelative Strenght Index: %f\n", simulation.StochasticOscillatorK, simulation.StochasticOscillatorD, simulation.RelativeStrenghtIndex)

	simulation.SetPriceBuy(convertions.StringToFloat64(currentPrice.Price))
	return simulation, err
}

func GetBestValue(s *simultor.Simulator, coin *symbols.Symbols, limitKline int, analizis *analizistrend.AnalizisTrend, cnnReport *cnn.FearAndGreedCNN) (float64, error) {
	simulation, err := simultor.NewSimulator(coin, *intervals.Minute, limitKline, cnnReport)
	if err != nil {
		return float64(0), nil
	}

	targetPrice := simulation.TargetPrice(s.GetPriceBuy())

	goodPrice := simulation.GoodPrice()

	err = database.SavewareHouse(coin, s, analizis, goodPrice, simulation.BestPriceCoin(), targetPrice)
	if err != nil {
		return float64(0), nil
	}

	fmt.Println(coin.Name)
	fmt.Println("*** Saving data in DynamoDB Table ***")

	return simulation.GoodPrice(), nil
}

func countdown(minute int) {
	second := minute * MinuteInSeconds
	for i := second; i >= 0; i-- {
		time.Sleep(time.Second)
	}
}

func RunCollector(coin *symbols.Symbols, limitKline, waitingPeriod, cycles, periodSell int, analizis *analizistrend.AnalizisTrend, cnnReport *cnn.FearAndGreedCNN) (*Util, error) {
	earn := float64(0)

	for i := 0; i < cycles; i++ {
		simulation, err := Iterractor(coin, limitKline, cnnReport)
		if err != nil {
			return nil, err
		}

		countdown(waitingPeriod)
		best, err := GetBestValue(simulation, coin, periodSell, analizis, cnnReport)
		if err != nil {
			return nil, err
		}

		if best > 0 {
			earn += best - simulation.GetPriceBuy()
		}

	}

	util := &Util{
		Earn: earn,
		Coin: *coin,
	}

	return util, nil
}
