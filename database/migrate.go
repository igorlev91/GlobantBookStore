package database

import (
	"database/sql"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func try_migrate(db *sql.DB) error {

	fmt.Println("Starting migrations")

	return nil
}
