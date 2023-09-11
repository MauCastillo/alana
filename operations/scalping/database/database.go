package database

import (
	"fmt"
	"time"

	"github.com/MauCastillo/alana/binance-api/symbols"
	"github.com/MauCastillo/alana/operations/scalping/models"
	"github.com/MauCastillo/alana/operations/scalping/simultor"
	"github.com/MauCastillo/alana/shared/dynamodb"
)

const (
	dateFormat = "2006-01-02 15:04:05"
)

var (
	databaseDynamoDB = dynamodb.NewDynamoDB()
)

func SavewareHouse(coin *symbols.Symbols, simulation *simultor.Simulator, goodPrice, hightPrice float64) error {
	now := time.Now().UTC()
	formatted := now.Format(dateFormat)

	op := models.Operation{
		Pass:                       fmt.Sprintf("%s_%s", coin.Name, formatted),
		Name:                       coin.Name,
		Date:                       formatted,
		Coin:                       *coin,
		FearAndGreedScore:          simulation.FearAndGreedCNN.FearAndGreed.Score,
		FearAndGreedPreviousClose:  simulation.FearAndGreedCNN.FearAndGreed.PreviousClose,
		FearAndGreedPrevious1Month: simulation.FearAndGreedCNN.FearAndGreed.Previous1Month,
		FearAndGreedPrevious1Year:  simulation.FearAndGreedCNN.FearAndGreed.Previous1Year,
		MarketMomentumSp500Score:   simulation.FearAndGreedCNN.MarketMomentumSp500.Score,
		MarketMomentumSp125Score:   simulation.FearAndGreedCNN.MarketMomentumSp125.Score,
		JunkBondDemandScore:        simulation.FearAndGreedCNN.JunkBondDemand.Score,
		SafeHavenDemandScore:       simulation.FearAndGreedCNN.SafeHavenDemand.Score,
		StochasticOscillatorK:      simulation.StochasticOscillatorK,
		StochasticOscillatorD:      simulation.StochasticOscillatorD,
		RelativeStrenghtIndex:      simulation.RelativeStrenghtIndex,
		PriceBuy:                   simulation.GetPriceBuy(),
		MarketInfo:                 simulation.RawDataDatabase(),
		MarketInfoBTC:              simulation.RawDataDatabaseBTC(),
		MarketInfoETH:              simulation.RawDataDatabaseETH(),
		Status:                     goodPrice > 0,
		GoodPrice:                  goodPrice,
		HightPrice:                 hightPrice,
	}

	return databaseDynamoDB.SaveRow(op)
}
