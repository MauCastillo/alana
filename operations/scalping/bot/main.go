package main

import (
	"fmt"

	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/operations/scalping/bot/utils"
	"github.com/MauCastillo/alana/operations/scalping/models"
)

var (
	index  = 0
	inputs = []models.ExecutionParams{
		{LimitKline: 60, WaitingPeriod: 1, PeriodSell: 1, Cycles: 1, Coin: *symbols.EthUsdt},
		{LimitKline: 60, WaitingPeriod: 1, PeriodSell: 1, Cycles: 1, Coin: *symbols.BtcUsdt},
		{LimitKline: 60, WaitingPeriod: 1, PeriodSell: 1, Cycles: 1, Coin: *symbols.BnbUsdt},
	}
)

func main() {
	for _, s := range inputs {
		fmt.Println("Name coin: ", s.Coin.Name)
		u, err := utils.RunCollector(&s.Coin, s.LimitKline, s.WaitingPeriod, s.Cycles, s.PeriodSell)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(u)
		index++
	}
}
