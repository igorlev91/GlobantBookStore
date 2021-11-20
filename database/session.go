package database

import (
	"fmt"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
)

var session db.Session

func init() {
	db_settings := mysql.ConnectionURL{
		Database: `book_store`,
		Host:     `localhost`,
		User:     `default`,
		Password: `default`,
	}

	// open database session
	fmt.Println("Try open session: ", db_settings)
	var err error
	session, err = mysql.Open(db_settings)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Session created")
}

func GetSession() db.Session {
	return session
}
