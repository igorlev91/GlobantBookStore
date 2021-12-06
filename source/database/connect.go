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

	"net/http"

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

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Setting.Database_username,
		Setting.Database_password, Setting.Database_host, Setting.Database_port, Setting.Database_name)

	fmt.Println(connectionString)

	log.Println("Connection: ", connectionString)
	var sqlDb *sql.DB

	db_count := handers.StringToInt(Setting.Database_max_connection)
	db_timeout := handers.StringToInt(Setting.Database_timeout)

	for i := 0; i <= db_count; i++ {
		sqlDb, err = sql.Open(DB_DRIVER, connectionString)
		if err != nil {
			log.Fatal(err)
		}
		defer sqlDb.Close()
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

	db.Connetion, err = Migrate(driver_connection)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connection to database")

	log.Println("Creating router")
	db.Router = mux.NewRouter()

	db.initializeRoutes()

	if server.Connetion == nil {
		return nil, err
	}

	return db.Connetion, nil
}

func Migrate(db *gorm.DB) (*gorm.DB, error) {
	db.Debug().AutoMigrate(&objects.Book{})
	db.Debug().AutoMigrate(&objects.Genre{})
	db.Migrator()

	return db, db.Error
}

func (d *Database) CloseDatabase() error {
	sqlDB, err := d.Connetion.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		return err
	}
	return nil
}

func (d Database) GetSession(db *gorm.DB) *gorm.DB {

	return d.Connetion
}

func (server *Database) RunServer(addr string) {

	fmt.Println("Listening to port 8000")
	log.Fatal(http.ListenAndServe(addr, server.Router))

}

func (db *Database) initializeRoutes() {

	db.Router.HandleFunc("/", db.CreateBookMethod).Methods("GET")
	db.Router.HandleFunc("/books/{id:[0-9]+}", db.GetBookByIdMethod).Methods("GET")
	db.Router.HandleFunc("/books", db.GetBooksByFilterMethod).Methods("GET")
}
