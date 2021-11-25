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
	if database.GetConnection() == nil {
		log.Fatal("Database session not created")
	}
	defer database.GetConnection().Close()

	//TODO Creating router logic
	log.Println("Creating router")
	router := mux.NewRouter()
	router.HandleFunc("/books/new", controllers.CreateBookMethod).Methods("POST")
	router.HandleFunc("/books/{id:[0-9]+}", controllers.GetBookByIdMethod).Methods("GET")

	//Creating server

	server := http.Server{
		Handler:      router,
		Addr:         "localhost:3306",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Server created.")
	log.Println("Listening started on: ", "localhost:3306")
	log.Fatal(server.ListenAndServe())

}
