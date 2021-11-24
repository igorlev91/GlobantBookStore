package database

import (
	"os"
)

type Config struct {
	Database_username   string `env:"MYSQL_USERNAME,required"`
	Database_password   string `env:"MYSQL_ROOT_PASSWORD,required"`
	Database_ServerName string `env:"MYSQL_HOST,default=localhost"`
	Database_Port       string `env:"MYSQL_PORT,required"`
	Database_name       string `env:"MYSQL_DATABASE,required"`
}

func getConfig() *Config {
	return &Config{
		Database_username:   os.Getenv("MYSQL_USER"),
		Database_password:   os.Getenv("MYSQL_ROOT_PASSWORD"),
		Database_ServerName: os.Getenv("MYSQL_SERVERNAME"),
		Database_Port:       os.Getenv("MYSQL_PORT"),
		Database_name:       os.Getenv("MYSQL_DATABASE"),
	}
}
