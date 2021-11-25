package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	objects "github.com/igorlev91/GlobantBookStore/source/objects"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetConnectionDatabase() string {
	config := getConfig()
	fmt.Println(config)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Database_username,
		config.Database_password, config.Database_ServerName, config.Database_Port, config.Database_name)
}

func Database_Connect() {

	connection := GetConnectionDatabase()
	fmt.Println(connection)

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic("connected db failed")
	}

	println("connected on db")

	db.AutoMigrate(&objects.Book{}, &objects.Genre{})
	if err != nil {
		panic("Failed to migrate tables")
	}

	DB = db
}
