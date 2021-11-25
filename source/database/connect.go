package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"

	"database/sql"
)

var session db.Session

const (
	DBHost     = "127.0.0.1"
	DBPort     = ":3306"
	DBUser     = "root"
	DBPassword = ""
	DBDbase    = "book_store"
)

// open database session test
func init() {
	conn := mysql.ConnectionURL{
		Database: `book_store`,
		Host:     "localhost:3306",
		User:     `root`,
		Password: `29394959abc`,
		Options: map[string]string{
			"multiStatements": "true",
		},
	}
	//var conn = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", DBUser, DBPassword, DBHost, DBDbase)

	// open db session
	fmt.Println("Start open session: ", conn)
	var err error
	session, err = mysql.Open(conn)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Session created")

	// migrate tables
	SQLDriver := session.Driver().(*sql.DB)
	err = MigrateDatabase(SQLDriver)
	if err != nil {
		fmt.Print(err)
	}
}

func GetConnection() db.Session {
	return session
}
