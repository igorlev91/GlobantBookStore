# GlobantBookStore

## Description 
Book Store 

## Setup and Development 
### Prerequisite
- git
- go 1.17 or Later

### Setup 
- Install Go <br>
  See [Go Installation](https://golang.org/doc/install)

 ## Structure
```
main.go -> Entry point of application
source/database/connect.go -> mysql connection established Here  
settings -> folder to store all connection and  related logic      

``` 


  # Using MySQL

1. Run MySql database by using Docker. Use the next command:

docker run --name mysql-bookstore -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=true -e MYSQL_DATABASE=bookstore -d mysql:latest



  # Using Docker
docker-compose build

// learning
1. We will need a MySQL driver
  - go get github.com/go-sql-driver/mysql
  - [read the documentation](https://github.com/go-sql-driver/mysql#installation)
 
1. Include the driver in your imports
  - _ "github.com/go-sql-driver/mysql"
  - [Read the documentation](https://github.com/go-sql-driver/mysql#usage)
1. Determine the Data Source Name
  - user:password@tcp(localhost:5555)/dbname?charset=utf8
  - [Read the documentation](https://github.com/go-sql-driver/mysql#dsn-data-source-name)
1. Open a connection
  -  "user:password@tcp(localhost:5555)/dbname?charset=utf8")
