package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func main() {
	http.HandleFunc("/", homePage)

	// Open up our database connection.
	db, err := sql.Open("mysql", "test_user:secret@tcp(db:3306)/test")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	log.Fatal(http.ListenAndServe(":8080", nil))

}
