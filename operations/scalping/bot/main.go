package main

import (
	"fmt"

	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/operations/scalping/bot/utils"
)

var (
	index = 0
)

func main() {
	limitKline := 60
	waitingPeriod := 0
	periodSell := 5
	cycles := 1
	coin := symbols.BtcBusd

	u, err := utils.RunCollector(coin, limitKline, waitingPeriod, cycles, periodSell)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(u)
	index++
}
