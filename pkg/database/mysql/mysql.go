package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yogarn/arten/pkg/config"
)

func ConnectDatabase() *sql.DB {
	dsn := config.LoadDataSourceName()
	db, err := openDatabase(dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func openDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
