package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/igorlev91/GlobantBookStore/source/controllers"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// init router
	//TODO Creating router logic
	log.Println("Creating router")
	router := mux.NewRouter()
	router.HandleFunc("/books/{id:[0-9]+}", controllers.GetBookByIdMethod).Methods("GET")

	//TODO Creating server
	log.Println("Creating server")
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
