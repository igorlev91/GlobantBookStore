package database

import (
	"encoding/json"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/igorlev91/GlobantBookStore/source/objects"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/gorm/logger"
)

type Database struct {
	Connetion *gorm.DB
	Router    *mux.Router
}

var (
	DB_DRIVER = "mysql"
)

func (server Database) InitializeDatabase() (*gorm.DB, error) {

	var err error
	err = LoadEnv("setting.env")
	if err != nil {
		panic("cannot load config: ")

	}
	db := Database{}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Setting.Database_username,
		Setting.Database_password, Setting.Database_host, Setting.Database_port, Setting.Database_name)

	fmt.Println(dsn)

	log.Println("Connection: ", dsn)
	var sqlDb *sql.DB

	db_count := StringToInt(Setting.Database_max_connection)
	db_timeout := StringToInt(Setting.Database_timeout)

	for i := 0; i <= db_count; i++ {
		sqlDb, err = sql.Open(DB_DRIVER, "book_manager:pseudo_pass@tcp(bookstore_database:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local")
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

	driver_connection, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", err)
	} else {
		fmt.Printf("We are connected to the %s database", err)
	}

	db.Connetion, err = SetupSchema(driver_connection)
	if err != nil {
		fmt.Printf("Failed to initialize database schema: %v", err)
	} else {
		fmt.Printf("good migrations")
	}

	fmt.Println("Successfully connection to database")

	books := []objects.Book{}
	if err := driver_connection.Find(&books).Error; err != nil {
		log.Fatal(err)
	}

	result, err := json.Marshal(books)

	fmt.Println(string(result))

	if server.Connetion == nil {
		return nil, err
	}

	return db.Connetion, nil
}

func SetupSchema(db *gorm.DB) (*gorm.DB, error) {

	log.Println("Updating database schema.")

	db.Debug().AutoMigrate(&objects.Book{})
	db.Debug().AutoMigrate(&objects.Genre{})

	log.Println("Database schema updated.")
	return db, db.Error
}

// Shutdown closes the database connection.
func (d *Database) CloseDatabase() (err error) {
	sqlDB, err := d.Connetion.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()

	sqlDB = nil
	return err
}

func (d Database) GetSession(db *gorm.DB) *gorm.DB {

	return d.Connetion
}

func StringToInt(val string) int {
	res, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return res
}
