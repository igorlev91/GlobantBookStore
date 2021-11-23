package database

import (
	"github.com/joeshaw/envdecode"
	"github.com/subosito/gotenv"
)

type Config struct {
	Port     int `env:"PORT,default=8080"`
	Database struct {
		Sql_username string `env:"MYSQL_USERNAME,required"`
		Sql_password string `env:"MYSQL_ROOT_PASSWORD,required"`
		Sql_host     string `env:"MYSQL_HOST,default=localhost"`
		Port         string `env:"MYSQL_PORT,required"`
		Sql_name     string `env:"MYSQL_DATABASE,required"`
	}
}

func NewConfig() Config {
	var cfg Config
	gotenv.Load(".env")
	err := envdecode.Decode(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func StartServer() {
	cfg := NewConfig()
	db, err := Database_Connect(&cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
