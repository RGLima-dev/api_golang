package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver
)

func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConnection)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		db.Close()
	}
	return db, nil
}
