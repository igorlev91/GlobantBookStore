package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"

	"database/sql"
)

var session db.Session

// open database session test
func init() {
	conn := mysql.ConnectionURL{
		Database: os.Getenv("BOOKSTORE_DBDRIVER"),
		Host:     os.Getenv("BOOKSTORE_HOST"),
		User:     os.Getenv("BOOKSTORE_USER"),
		Password: os.Getenv("BOOKSTORE_PASSWORD"),
		Options: map[string]string{
			"multiStatements": "true",
		},
	}

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
