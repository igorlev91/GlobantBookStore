package database

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jinzhu/gorm"
)

func migrate_data(db *gorm.DB) error {

	fmt.Println("Start migrations")
	driver, err := mysql.WithInstance(db.DB(), &mysql.Config{})
	if err != nil {
		return errors.New("mysql.WithInstance: " + err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("file://database/migrations/", "mysql", driver)
	if err != nil {
		return errors.New("migrate.NewWithDatabaseInstance: " + err.Error())
	}

	err = m.Up()
	switch err = m.Up(); err {
	case nil:
		fmt.Println("Migrations executed")
	case migrate.ErrNoChange:
		fmt.Println("No migrations to execute")
	default:
		return errors.New("m.Steps: " + err.Error())
	}

	return nil
}
