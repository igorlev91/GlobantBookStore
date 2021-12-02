package main

import (
	"fmt"

	"github.com/igorlev91/GlobantBookStore/source/database"
)

const (
	ServerAddress string = "8080"
)

var server = database.Database{}

func main() {
	fmt.Println("Start session")
	server.InitializeDatabase()
	server.RunServer(ServerAddress)
	/*// create database session
	database.InitializeDatabase()

	// Creating router logic
	log.Println("Creating router")
	router := mux.NewRouter()
	/*
		session, err := database.InitializeDatabase()
		if err != nil {
			panic("cannot init database ")

		}
		/*
			session := &controllers.BookORM{
				DB: database.GetSession(),
			}

			router.HandleFunc("/books/new", session.CreateBookMethod).Methods("POST")
			router.HandleFunc("/books/{id:[0-9]+}", session.GetBookByIdMethod).Methods("GET")
			router.HandleFunc("/books/{id:[0-9]+}", session.GetBooksByFilterMethod).Methods("GET")

	server := http.Server{
		Handler:      router,
		Addr:         ServerAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server created.")
	log.Println("Listening started on: ", server.Addr)
	log.Fatal(server.ListenAndServe())
	*/

}
