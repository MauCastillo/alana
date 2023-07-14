package main

import (
	"fmt"

	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/operations/scalping/bot/utils"
	"github.com/MauCastillo/alana/operations/scalping/database"
	"github.com/MauCastillo/alana/operations/scalping/models"
	"github.com/MauCastillo/alana/shared/env"
)

var (
	index  = 0
	inputs = []models.ExecutionParams{
		{LimitKline: 60, WaitingPeriod: 1, PeriodSell: 1, Cycles: 2, Coin: *symbols.EthUsdt},
		{LimitKline: 60, WaitingPeriod: 1, PeriodSell: 1, Cycles: 2, Coin: *symbols.BtcUsdt},
		{LimitKline: 60, WaitingPeriod: 1, PeriodSell: 1, Cycles: 2, Coin: *symbols.BnbUsdt},
	}
)

const (
	TableFormat = "oscillator_strenght_%s"
)

var (
	IsCreateTable = env.GetBool("CREATE_TABLE", true)
)

func main() {
	var tableName string
	for _, s := range inputs {
		if IsCreateTable {
			tableName = fmt.Sprintf(TableFormat, s.Coin.Name)
			fmt.Println("Creando Tablas => ", tableName)
			err := database.Init(tableName)
			if err != nil {
				fmt.Println(err)
			}
		}

		fmt.Println("Name coin: ", s.Coin.Name)
		u, err := utils.RunCollector(&s.Coin, s.LimitKline, s.WaitingPeriod, s.Cycles, s.PeriodSell)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(u)
		index++
	}
}
