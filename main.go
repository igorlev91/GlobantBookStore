package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/igorlev91/GlobantBookStore/source/controllers"
	"github.com/igorlev91/GlobantBookStore/source/database"
)

func main() {
	fmt.Println("Start session")
	// create database session
	database.InitializeDatabase()

	// Creating router logic
	log.Println("Creating router")
	router := mux.NewRouter()
	session := &controllers.BookORM{
		DB: database.GetSession(),
	}
	router.HandleFunc("/books/new", session.CreateBookMethod).Methods("POST")
	router.HandleFunc("/books/{id:[0-9]+}", session.GetBookByIdMethod).Methods("GET")
	router.HandleFunc("/books/{id:[0-9]+}", session.GetBooksByFilterMethod).Methods("GET")

	//Creating server
	server := http.Server{
		Handler:      router,
		Addr:         database.Setting.ServerAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Server created.")
	log.Println("Listening started on: ", server.Addr)
	log.Fatal(server.ListenAndServe())

}
