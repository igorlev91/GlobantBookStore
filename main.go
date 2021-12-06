package main

import (
	"fmt"

	"github.com/igorlev91/GlobantBookStore/source/database"
)

const (
	ServerAddress string = "8000"
)

var server = database.Database{}

func main() {
	fmt.Println("Start session")
	server.InitializeDatabase()
	server.RunServer(ServerAddress)

}
