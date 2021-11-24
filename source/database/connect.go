package database

import (
	"database/sql"
	"fmt"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func GetConnectionDatabase() string {
	config := getConfig()
	fmt.Println(config)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Database_username,
		config.Database_password, config.Database_ServerName, config.Database_Port, config.Database_name)
}

func Database_Connect() (*gorm.DB, error) {
	connection := GetConnectionDatabase()

	sqlDb, err := sql.Open("mysql", connection)

	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db, err
}
