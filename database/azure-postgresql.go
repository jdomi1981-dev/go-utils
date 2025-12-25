package database

import (
	"database/sql"
	"fmt"
)

const (
	DRIVER_NAME = "postgres"
)

func ConnectDB(dbHost string, dbUser string, dbPassword string, dbName string, resourceName string) (*sql.DB, error) {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", dbHost, dbUser, dbPassword, dbName)

	db, err := sql.Open(DRIVER_NAME, connectionString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ExecuteQuery(db *sql.DB, sqlStatement string, resourceName string) (*sql.Rows, error) {
	defer db.Close()
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func ExecuteTransaction(db *sql.DB, sqlStatement string, resourceName string) (int64, error) {
	defer db.Close()
	result, err := db.Exec(sqlStatement)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
