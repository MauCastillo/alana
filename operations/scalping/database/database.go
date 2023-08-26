package database

import (
	"fmt"

	"github.com/MauCastillo/alana/operations/scalping/models"
	"github.com/MauCastillo/alana/operations/scalping/simultor"
	"github.com/MauCastillo/alana/shared/env"
	"github.com/MauCastillo/alana/shared/sqlite"
)

var (
	IsCreateTable = env.GetBool("CREATE_TABLE", true)
	TableFormat   = env.GetString("TABLE_NAME_FORMAT", "coinmarket_image_%s")
)

func SavewareHouse(simulation *simultor.Simulator, goodPrice, hightPrice float64) error {
	op := models.Operation{
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
	}

	database, err := sqlite.NewDatabase()
	if err != nil {
		return err
	}

	tableName := fmt.Sprintf(TableFormat, simulation.Symbol.Name)

	listOp := []models.Operation{op}
	err = database.InsertOperations(tableName, goodPrice, hightPrice, listOp)
	if err != nil {
		return err
	}

	return nil
}

func Init(tableName string) error {
	database, err := sqlite.NewDatabase()
	if err != nil {
		return err
	}

	if IsCreateTable {
		err = database.CreateNewTable(tableName)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateNewTable(table string) error {
	tableName := fmt.Sprintf(TableFormat, table)

	err := Init(tableName)
	if err != nil {
		return err
	}

	return nil
}
