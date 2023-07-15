package main

import (
	"fmt"
	"sync"

	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/operations/scalping/bot/utils"
	"github.com/MauCastillo/alana/operations/scalping/database"
	"github.com/MauCastillo/alana/operations/scalping/models"
	"github.com/MauCastillo/alana/shared/env"
)

const (
	TableFormat   = "oscillator_strenght_%s"
	limitKline    = 120
	waitingPeriod = 30
	periodSell    = 30
	cycles        = 192
)

var (
	IsCreateTable = env.GetBool("CREATE_TABLE", true)
	inputs        = []models.ExecutionParams{
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.EthUsdt},
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.BtcUsdt},
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.BnbUsdt},
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.AdaUsdt},
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.SolUsdt},
		{LimitKline: limitKline, WaitingPeriod: waitingPeriod, PeriodSell: periodSell, Cycles: cycles, Coin: *symbols.MaticUsdt},
	}
)

func asyncFunction(s models.ExecutionParams) {
	var tableName string

	if IsCreateTable {
		tableName = fmt.Sprintf(TableFormat, s.Coin.Name)
		fmt.Println("Creando Tablas => ", tableName)
		err := database.Init(tableName)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("=> Started collector to : ", s.Coin.Name)
	_, err := utils.RunCollector(&s.Coin, s.LimitKline, s.WaitingPeriod, s.Cycles, s.PeriodSell)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("=> Ended collector to : ", s.Coin.Name)

}

func main() {
	var wg sync.WaitGroup

	for _, s := range inputs {
		wg.Add(1)
		go func(index models.ExecutionParams) {
			defer wg.Done()
			asyncFunction(index)
		}(s)
	}

	wg.Wait()
	fmt.Println("Todas las funciones asíncronas han finalizado")
}
