package dynamodb

import (
	"errors"

	"github.com/MauCastillo/alana/shared/env"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	// ErrorTimeOutSatellite timeout after save information satellite
	ErrorTimeOutSatellite = errors.New("satellite information is no longer available timeout")
	// ErrorNotFoundSatellite satellite not found
	ErrorNotFoundSatellite = errors.New("satellite not found")
	TableName             = env.GetString("TABLE_NAME", "basic_training")
	TimeOut               = env.GetInt64("TIMEOUT", 5)
)

const (
	dateFormat = "2006-01-02 15:04:05"
)


type DynamoDB struct {
	DataBase *dynamodb.DynamoDB
}

func NewDynamoDB() *DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)
	return &DynamoDB{DataBase: svc}
}

func (d *DynamoDB) SaveRow(item interface{}) error {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(TableName),
	}

	_, err = d.DataBase.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

func (d *DynamoDB) GetRow(key, value string) (map[string]*dynamodb.AttributeValue, error) {
	result, err := d.DataBase.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key: map[string]*dynamodb.AttributeValue{
			key: {
				S: aws.String(value),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, ErrorNotFoundSatellite
	}

	return result.Item, err
}
