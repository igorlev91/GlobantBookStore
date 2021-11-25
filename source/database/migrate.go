package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateDatabase(db *sql.DB) error {

	fmt.Println("Start migrations")
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Println("Error mysql.WithInstance")
		panic(err)
	}

	migration, err := migrate.NewWithDatabaseInstance("file://database/migrations/", "mysql", driver)
	if err != nil {
		log.Println("Error migrate.NewWithDatabaseInstance")
		panic(err)
	}

	err = migration.Up()
	if migration != nil {
		err := migration.Steps(1)
		if err != nil {
			panic(err)
		}
	}


	log.Println("Database migrated")

	return nil
}
