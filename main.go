package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/igorlev91/GlobantBookStore/source/controllers"

	"github.com/upper/db/v4/adapter/mysql"

	"github.com/upper/db/v4"

	"database/sql"
)

var session db.Session

func GetSession() db.Session {
	return session
}

func main() {
	fmt.Println("Start session")

	db_settings := mysql.ConnectionURL{
		Database: `book_store`,
		Host:     `127.0.0.1`,
		User:     `root`,
		Password: `29394959abc`,
		Options: map[string]string{
			"multiStatements": "true",
		},
	}

	// open db session
	fmt.Println("Try open session: ", db_settings)
	var err error
	session, err = mysql.Open(db_settings)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Session created")

	// try migrate
	internalSQLDriver := session.Driver().(*sql.DB)
	err = database.migrate_data(internalSQLDriver)
	if err != nil {
		fmt.Print(err)
	}

	//TODO Creating router logic
	log.Println("Creating router")
	router := mux.NewRouter()
	router.HandleFunc("/books/new", controllers.CreateBookMethod).Methods("POST")
	router.HandleFunc("/books/{id:[0-9]+}", controllers.GetBookByIdMethod).Methods("GET")

	//Creating server

	server := http.Server{
		Handler:      router,
		Addr:         "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Server created.")
	log.Println("Listening started on: ", "localhost:8000")
	log.Fatal(server.ListenAndServe())

}
