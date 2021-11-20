package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/igorlev91/GlobantBookStore/controllers"
	"github.com/igorlev91/GlobantBookStore/database"
)

func main() {
	// create database session
	if database.GetSession() == nil {
		log.Fatal("Database session not created")
	}
	defer database.GetSession().Close()

	// create router
	fmt.Println("Creating router")
	router := mux.NewRouter()
	router.HandleFunc("/books/{id:[0-9]+}", controllers.GetBookByIdMethod)
	// listen & serve
	fmt.Println("Creating server")
	server := http.Server{
		Handler:      router,
		Addr:         ":3000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	fmt.Println("Server created.")
	fmt.Println("Listening started on", server.Addr)
	log.Fatal(server.ListenAndServe())
}
