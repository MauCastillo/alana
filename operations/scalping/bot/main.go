package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/operations/scalping/bot/utils"
	"github.com/MauCastillo/alana/operations/scalping/database"
	"github.com/MauCastillo/alana/operations/scalping/models"
	"github.com/MauCastillo/alana/shared/env"
	
)
const(
	dateFormat = "January 02, 2006  15:04:05"
)

var (
	limitKline    = int(env.GetInt64("LIMIT_KLINE", 6))
	waitingPeriod = int(env.GetInt64("WAITING_PERIOD", 1))
	periodSell    = waitingPeriod
	cycles        = int(env.GetInt64("CYCLES", 1))

	inputs = []models.ExecutionParams{
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.EthUsdt},
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.BtcUsdt},
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.BnbUsdt},
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.AdaUsdt},
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.SolUsdt},
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.MaticUsdt},
	}
)

func collector(s models.ExecutionParams) {

	fmt.Println("=> Started collector to : ", s.Coin.Name)
	_, err := utils.RunCollector(&s.Coin, s.LimitKline, s.WaitingPeriod, s.Cycles, s.PeriodSell)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("=> Ended collector to : ", s.Coin.Name)

}

func main() {
	start := time.Now()

	var wg sync.WaitGroup

	for _, s := range inputs {
		err := database.CreateNewTable(s.Coin.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}

		wg.Add(1)
		go func(index models.ExecutionParams) {
			defer wg.Done()
			collector(index)
		}(s)
	}

	wg.Wait()
	currentTime := time.Now()

	fmt.Println("=> Start Time: ", start.Format(dateFormat))
	fmt.Println("___________")
	fmt.Printf(" => Ciclos: %d Periodo de espera: %d Minutos de Klines: %d \n", cycles, waitingPeriod, limitKline)
	fmt.Println("___________")
	fmt.Println("=> Todas las funciones asíncronas han finalizado: ", currentTime.Format(dateFormat))
}
