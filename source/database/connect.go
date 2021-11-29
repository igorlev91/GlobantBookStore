package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"database/sql"
)

var (
	Db *gorm.DB
)

// open database session test
func InitializeDatabase() {

	var err error
	err = LoadEnv("setting.env")
	if err != nil {
		panic(err.Error())
		//log.Fatal("cannot load config: ", err)
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Setting.Database_username,
		Setting.Database_password, Setting.Database_host, Setting.Database_port, Setting.Database_name)

	fmt.Println(connectionString)
	sqlDb, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}
	Db = db
	//db.AutoMigrate(&objects.Book{}, &objects.Genre{})
}
