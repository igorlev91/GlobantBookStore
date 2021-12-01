package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/igorlev91/GlobantBookStore/source/handers"
	"github.com/igorlev91/GlobantBookStore/source/objects"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"database/sql"
	"log"
	"time"
)

var (
	Db *gorm.DB
)

func InitializeDatabase() {

	var err error
	err = LoadEnv("setting.env")
	if err != nil {
		panic("cannot load config: ")

	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Setting.Database_username,
		Setting.Database_password, Setting.Database_host, Setting.Database_port, Setting.Database_name)

	fmt.Println(connectionString)
	// open db session
	log.Println("Connection: ", connectionString)
	var sqlDb *sql.DB

	db_count := handers.StringToInt(Setting.Database_max_connection)
	db_timeout := handers.StringToInt(Setting.Database_timeout)

	for i := 0; i <= db_count; i++ {
		sqlDb, err = sql.Open("mysql", connectionString)
		if err != nil {
			log.Fatal(err)
		}
		if i == 2 {
			fmt.Println("Waiting opening database")
			break
		}
		if err := sqlDb.Ping(); err != nil {
			fmt.Printf("Try open session %d\n", i)
			time.Sleep(time.Duration(db_timeout) * time.Second)
		} else {
			fmt.Println("Success")
			break
		}
	}
	fmt.Println("Successfully connection to database")

	driver_connection, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                "mysql",
		Conn:                      sqlDb,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic(err.Error())
	}

	Db, err = Migrate(driver_connection)

}

func Migrate(db *gorm.DB) (*gorm.DB, error) {
	db.AutoMigrate(&objects.Book{}, &objects.Genre{})
	return db, db.Error
}

func GetSession() *gorm.DB {
	return Db
}
