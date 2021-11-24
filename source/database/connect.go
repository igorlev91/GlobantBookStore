package database

import (
	"database/sql"
	"fmt"
	"time"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceDB struct {
	DB *gorm.DB
}

func GetConnectionDatabase() string {
	config := getConfig()
	fmt.Println(config)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Database_username,
		config.Database_password, config.Database_ServerName, config.Database_Port, config.Database_name)
}

func Database_Connect(s *ServiceDB) {
	fmt.Println("Start Database_Connect")
	connection := GetConnectionDatabase()
	fmt.Println(connection)

	var client *sql.DB
	client, err := sql.Open("mysql", connection)

	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: client,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	client.SetConnMaxIdleTime(time.Minute * 3)
	client.SetMaxIdleConns(10)
	client.SetMaxIdleConns(10)
	fmt.Println("Successfully connect to MySql")

	// try migrate
	//err = migrate_data(db)
	//if err != nil {
	//	fmt.Print(err)
	//}
	s.DB = db

}
