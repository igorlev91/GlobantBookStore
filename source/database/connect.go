package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/igorlev91/GlobantBookStore/source/objects"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"database/sql"
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
	sqlDb, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
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
	Db = db
	//Db, err = Migrate(db)

}

func Migrate(db *gorm.DB) (*gorm.DB, error) {
	db.AutoMigrate(&objects.Book{}, &objects.Genre{})
	return db, db.Error
}

func GetSession() *gorm.DB {
	return Db
}
