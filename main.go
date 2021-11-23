package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/igorlev91/GlobantBookStore/source/database"
)

func main() {
	database.StartServer()
}
