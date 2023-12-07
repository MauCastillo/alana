package database

import (
	"context"
	"fmt"
	"time"

	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/operations/scalping/models"
	"github.com/MauCastillo/alana/operations/scalping/simultor"
	"github.com/MauCastillo/alana/shared/dynamodb"
	"github.com/MauCastillo/alana/shared/google/analizistrend"
)

const (
	dateFormat = "2006-01-02 15:04:05"
)

var (
	databaseDynamoDB = dynamodb.NewDynamoDB()
)

func SavewareHouse(coin *symbols.Symbols, simulation *simultor.Simulator, analizis *analizistrend.AnalizisTrend, goodPrice, hightPrice, targetPrice float64) error {
	now := time.Now().UTC()
	formatted := now.Format(dateFormat)

	err := analizis.Refresh(context.Background())
	if err != nil {
		print("=> error refresh: ", err.Error())
	}

	balanceCryptocurrency := analizis.RealtimeArticleBalance.Cryptocurrency
	balanceEconomic := analizis.RealtimeArticleBalance.Economic

	op := models.Operation{
		Pass:                       fmt.Sprintf("%s_%s", coin.Name, formatted),
		Name:                       coin.Name,
		Date:                       formatted,
		FearAndGreedScore:          simulation.FearAndGreedCNN.FearAndGreed.Score,
		FearAndGreedPreviousClose:  simulation.FearAndGreedCNN.FearAndGreed.PreviousClose,
		FearAndGreedPrevious1Month: simulation.FearAndGreedCNN.FearAndGreed.Previous1Month,
		FearAndGreedPrevious1Year:  simulation.FearAndGreedCNN.FearAndGreed.Previous1Year,
		MarketMomentumSp500Score:   simulation.FearAndGreedCNN.MarketMomentumSp500.Score,
		MarketMomentumSp125Score:   simulation.FearAndGreedCNN.MarketMomentumSp125.Score,
		JunkBondDemandScore:        simulation.FearAndGreedCNN.JunkBondDemand.Score,
		SafeHavenDemandScore:       simulation.FearAndGreedCNN.SafeHavenDemand.Score,
		PriceBuy:                   simulation.GetPriceBuy(),
		MarketCurrency:             simulation.RawDataDatabase(),
		VolumenInfoBTC:             simulation.VolumenCurrencyBTC(),
		VolumenCurrency:            simulation.VolumenCurrency(),
		GoodPrice:                  goodPrice,
		Status:                     goodPrice > 0,
		TargetPrice:                targetPrice,
		Economic:                   balanceEconomic,
		Cryptocurrency:             balanceCryptocurrency,
	}

	return databaseDynamoDB.SaveRow(op)
}
