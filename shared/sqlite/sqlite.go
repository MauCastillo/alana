package sqlite

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/MauCastillo/alana/operations/scalping/models"
	"github.com/MauCastillo/alana/shared/env"
	_ "github.com/mattn/go-sqlite3" //nosec
)

const (
	FileExtention = "sqlite3"
)

var (
	dataBaseName = env.GetString("DATABASE_NAME", "data-warehouse")
)

type DataBase struct {
	Database *sql.DB
}

func NewDatabase() (*DataBase, error) {
	nameFile := fmt.Sprintf("%s.%s", dataBaseName, FileExtention)

	db, err := sql.Open("sqlite3", nameFile)
	if err != nil {
		return nil, err
	}

	database := &DataBase{Database: db}

	return database, nil
}

func (d *DataBase) CreateNewTable(tableName string) error {
	sts := `DROP TABLE IF EXISTS %s;
		CREATE TABLE %s(id INTEGER PRIMARY KEY, operation TEXT, good_price float);`

	queryCreation := fmt.Sprintf(sts, tableName, tableName)

	_, err := d.Database.Exec(queryCreation)
	if err != nil {
		return err
	}

	return nil
}

func (d *DataBase) InsertOperations(tableName string, goodPrice float64, op []models.Operation) error {
	var queryInsert string
	for _, operation := range op {
		out, err := json.Marshal(operation)
		if err != nil {
			continue
		}

		queryInsert += fmt.Sprintf("INSERT INTO %s(operation, good_price) VALUES('%s',%f);", tableName, string(out), goodPrice)
	}

	_, err := d.Database.Exec(queryInsert)

	if err != nil {
		return err
	}

	return nil
}
