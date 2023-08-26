package sqlite

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/MauCastillo/alana/operations/scalping/models"
	"github.com/MauCastillo/alana/shared/env"
	_ "github.com/mattn/go-sqlite3" //nosec
)

const (
	FileExtention = "sqlite3"
	// path          = "../../../../"
	dateFormat = "2006-01-02"
)

var (
	dataBaseName = env.GetString("DATABASE_NAME", "data_warehouse")
	// databaseVersion  = env.GetString("VERSION", "01")
	databaseFileName = getDatabaseNameFile()
)

type DataBase struct {
	Database *sql.DB
}

func getDatabaseNameFile() string {
	t := time.Now().UTC()
	s2 := t.Format(dateFormat)

	database := fmt.Sprintf("%s_v%s.%s", dataBaseName, s2, FileExtention)

	return database
}

func NewDatabase() (*DataBase, error) {
	db, err := sql.Open("sqlite3", databaseFileName)
	if err != nil {
		return nil, err
	}

	database := &DataBase{Database: db}

	return database, nil
}

func (d *DataBase) CreateNewTableForce(tableName string) error {
	sts := `DROP TABLE IF EXISTS %s;
		CREATE TABLE %s(id INTEGER PRIMARY KEY, operation TEXT, good_price float, hight_price float);`

	queryCreation := fmt.Sprintf(sts, tableName, tableName)

	_, err := d.Database.Exec(queryCreation)
	if err != nil {
		return err
	}

	return nil
}

func (d *DataBase) CreateNewTable(tableName string) error {
	sts := `CREATE TABLE IF NOT EXISTS %s(id INTEGER PRIMARY KEY, operation TEXT, good_price float, hight_price float);`

	queryCreation := fmt.Sprintf(sts, tableName)

	_, err := d.Database.Exec(queryCreation)
	if err != nil {
		return err
	}

	return nil
}

func (d *DataBase) InsertOperations(tableName string, goodPrice, hightPrice float64, op []models.Operation) error {
	var queryInsert string
	for _, operation := range op {
		out, err := json.Marshal(operation)
		if err != nil {
			continue
		}

		queryInsert += fmt.Sprintf("INSERT INTO %s(operation, good_price, hight_price) VALUES('%s',%f,%f);", tableName, string(out), goodPrice, hightPrice)
	}

	_, err := d.Database.Exec(queryInsert)

	if err != nil {
		return err
	}

	return nil
}
