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

func (d *DataBase) CreateNewTable() error {
	sts := `DROP TABLE IF EXISTS training_data;
		CREATE TABLE training_data(id INTEGER PRIMARY KEY, data TEXT, good_price float);`

	_, err := d.Database.Exec(sts)
	if err != nil {
		return err
	}

	return nil
}

func (d *DataBase) InsertOperations(op []models.Operation, goodPrice float64) error {
	var queryInsert string
	for _, operation := range op {
		out, err := json.Marshal(operation)
		if err != nil {
			continue
		}

		queryInsert += fmt.Sprintf("INSERT INTO training_data(data, good_price) VALUES('%s',%f);", string(out), goodPrice)
	}

	_, err := d.Database.Exec(queryInsert)

	if err != nil {
		return err
	}

	return nil
}
