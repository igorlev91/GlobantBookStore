package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/igorlev91/GlobantBookStore/source/database"
	"github.com/igorlev91/GlobantBookStore/source/handlers"
)

const (
	ServerAddress string = "8000"
)

var server = database.Database{}

func main() {

	//defer database.CloseDatabase(db)
	fmt.Println("Start session")
	server.InitializeDatabase()

	// init router
	log.Println("Creating router")
	router := mux.NewRouter()

	handler := &handlers.ORM{
		DB: server.Connetion,
	}
	router.HandleFunc("/books/{id:[0-9]+}", handler.GetBookByIdMethod).Methods("GET")
	router.HandleFunc("/books/{id:[0-9]+}", handler.UpdateBookMethod).Methods("PUT")
	router.HandleFunc("/books/{id:[0-9]+}", handler.DeleteBookMethod).Methods("DELETE")
	router.HandleFunc("/books", handler.GetBooksByFilterMethod).Methods("GET")
	router.HandleFunc("/books/new", handler.CreateBookMethod).Methods("POST")

	// listen & serve
	log.Println("Creating server")
	server := http.Server{
		Handler:      router,
		Addr:         ServerAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Server created.")
	log.Println("Listening started on: ", server.Addr)
	log.Fatal(server.ListenAndServe())

}
