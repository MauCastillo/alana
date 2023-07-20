package sqlite

import (
	"math/rand"
	"os"
	"testing"

	"github.com/MauCastillo/alana/operations/scalping/models"
	"github.com/stretchr/testify/require"
)

func TestNewDatabase(t *testing.T) {
	c := require.New(t)

	data, err := NewDatabase()
	c.NoError(err)
	c.NotNil(data)

	tableName := "testing"

	defer os.Remove(databaseFileName)

	err = data.CreateNewTable(tableName)
	c.NoError(err)
	op := []models.Operation{
		{FearAndGreedScore: float64(34.42), FearAndGreedPreviousClose: float64(156.126), SafeHavenDemandScore: rand.Float64() * 100},
		{FearAndGreedScore: float64(214.14), FearAndGreedPreviousClose: float64(256.236), SafeHavenDemandScore: rand.Float64() * 100},
		{FearAndGreedScore: float64(345.674), FearAndGreedPreviousClose: float64(356.186), SafeHavenDemandScore: rand.Float64() * 100},
		{FearAndGreedScore: float64(243.454), FearAndGreedPreviousClose: float64(456.126), SafeHavenDemandScore: rand.Float64() * 100},
		{FearAndGreedScore: float64(23.564), FearAndGreedPreviousClose: float64(556.106), SafeHavenDemandScore: rand.Float64() * 100},
	}
	err = data.InsertOperations(tableName, rand.Float64()*10, rand.Float64()*10, op)
	c.NoError(err)
}
