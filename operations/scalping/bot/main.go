package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/operations/scalping/bot/utils"
	"github.com/MauCastillo/alana/operations/scalping/models"
	"github.com/MauCastillo/alana/shared/cnn"
	"github.com/MauCastillo/alana/shared/env"
	"github.com/MauCastillo/alana/shared/google/analizistrend"
)

const (
	dateFormat = "January 02, 2006  15:04:05"
)

var (
	limitKline    = int(env.GetInt64("LIMIT_KLINE", 6))
	waitingPeriod = int(env.GetInt64("WAITING_PERIOD", 1))
	periodSell    = waitingPeriod
	cycles        = int(env.GetInt64("CYCLES", 4))

	wg sync.WaitGroup

	inputs = []models.ExecutionParams{
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.EthUsdt},
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.BtcUsdt},
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.BnbUsdt},
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.DotUsdt},
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.AvaxUsdt},
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.SolUsdt},
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.MaticUsdt},
	}
)

func collector(s models.ExecutionParams, analizis *analizistrend.AnalizisTrend, cnn *cnn.FearAndGreedCNN) {
	defer wg.Done()
	fmt.Println("=> Started collector to : ", s.Coin.Name)
	_, err := utils.RunCollector(&s.Coin, s.LimitKline, s.WaitingPeriod, s.Cycles, s.PeriodSell, analizis, cnn)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("=> Ended collector to : ", s.Coin.Name)

}

func main() {
	analizis, err := analizistrend.NewAnalizisTrend(context.Background(), "EN", "US", "b")
	if err != nil {
		print(err)
		panic(err)
	}

	requestCNN, err := cnn.NewFearAndGreedCNN()
	if err != nil {
		print(err)
		panic(err)
	}

	start := time.Now()
	for _, s := range inputs {
		wg.Add(1)
		go func(index models.ExecutionParams) {
			collector(index, analizis, requestCNN)
		}(s)

		fmt.Println(" ***************************** ")
		fmt.Println(s)
		fmt.Print(" *** Collector Information *** ")
	}

	wg.Wait()
	currentTime := time.Now()

	fmt.Println("=> Starting collector Time: ", start.Format(dateFormat))
	fmt.Println("_____________________________________________________")
	fmt.Printf(" => Ciclos: %d Periodo de espera: %d Minutos de Klines: %d \n", cycles, waitingPeriod, limitKline)
	fmt.Println("_____________________________________________________")
	fmt.Println("=> Todas las funciones asíncronas han finalizado: ", currentTime.Format(dateFormat))
}
