package db

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver
)

// Conect Open connection with database and return it
func Conect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConnectionDataBase)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
