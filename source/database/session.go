package database

import (
	"log"

	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Database_Connect(setting *Config) (*sql.DB, error) {
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		setting.Database.Sql_username,
		setting.Database.Sql_password,
		setting.Database.Sql_host,
		setting.Database.Port,
		setting.Database.Sql_name,
	)

	db, err := sql.Open("mysql", dbConfig)

	if err != nil {
		return nil, err
	}

	return db, err
}

func NewServer(setting *Config) {
	log.Println("Server listening on port", setting.Database.Port)
}
