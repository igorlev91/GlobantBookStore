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

	if database.GetSession() == nil {
		log.Fatal("Database session not created")
	}
	defer database.GetSession().Close()

	fmt.Println("Creating router")
	router := mux.NewRouter()
	router.HandleFunc("/books/{id:[0-9]+}", controllers.GetBookByIdMethod)

	fmt.Println("Creating server")
	server := http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server created.")
	fmt.Println("Listening started on", server.Addr)
	log.Fatal(server.ListenAndServe())
}
