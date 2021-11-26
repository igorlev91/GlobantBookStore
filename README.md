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

  # Using MySQL

1. Run MySql database by using Docker. Use the next command:

docker run --name mysql-bookstore -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=true -e MYSQL_DATABASE=bookstore -d mysql:latest

If docker error come up: Error response from daemon: Ports are not available: listen tcp 0.0.0.0:3306: bind: address already in use. Use the next command:
  docker run -e MYSQL_ROOT_PASSWORD=root  --name localMysql -d  -p 3366:3306  mysql:8.0.23


// learning
1. We will need a MySQL driver
  - go get github.com/go-sql-driver/mysql
  - [read the documentation](https://github.com/go-sql-driver/mysql#installation)
  - [see all SQL drivers](https://github.com/golang/go/wiki/SQLDrivers)
  - [Astaxie's book](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.2.html)
1. Include the driver in your imports
  - _ "github.com/go-sql-driver/mysql"
  - [Read the documentation](https://github.com/go-sql-driver/mysql#usage)
1. Determine the Data Source Name
  - user:password@tcp(localhost:5555)/dbname?charset=utf8
  - [Read the documentation](https://github.com/go-sql-driver/mysql#dsn-data-source-name)
1. Open a connection
  -  "user:password@tcp(localhost:5555)/dbname?charset=utf8")
