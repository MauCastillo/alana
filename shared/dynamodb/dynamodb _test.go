package dynamodb

import (
	"testing"
	"time"

	"github.com/MauCastillo/alana/operations/scalping/models"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/stretchr/testify/require"
)

var (
	now       = time.Now().UTC()
	formatted = now.Format(dateFormat)
)

func TestNewDynamoDB(t *testing.T) {
	c := require.New(t)

	database := NewDynamoDB()
	c.NotEmpty(database)

	item := models.Operation{
		Date: formatted,
		Name: "Claro",
	}

	err := database.SaveRow(item)
	c.NoError(err)

}

func TestSaveDynamoDBError(t *testing.T) {
	c := require.New(t)

	database := NewDynamoDB()
	c.NotEmpty(database)

	item := models.Operation{}

	err := database.SaveRow(item)
	c.Contains(err.Error(), "ValidationException: One or more parameter values were invalid:")

}

func TestGetDynamoDB(t *testing.T) {
	c := require.New(t)

	database := NewDynamoDB()
	c.NotEmpty(database)

	now := time.Now().UTC()
	formatted := now.Format(dateFormat)

	item := models.Operation{
		Name: "claro",
		Date:                       formatted,
		FearAndGreedPrevious1Month: 3943,
	}

	err := database.SaveRow(item)
	c.NoError(err)

	result, err := database.GetRow("name", "claro")
	c.NoError(err)

	// Convertir el resultado en la estructura original
	var op models.Operation
	err = dynamodbattribute.UnmarshalMap(result, &op)
	c.NoError(err)

	c.Equal(op.FearAndGreedPrevious1Month, float64(3943))

}

func TestGetDynamoDBNotFound(t *testing.T) {
	c := require.New(t)

	database := NewDynamoDB()
	c.NotEmpty(database)
	_, err := database.GetRow("name", "sky_test_Gato")

	c.EqualError(err, ErrorNotFoundSatellite.Error())
}
