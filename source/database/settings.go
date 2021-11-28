package database

import (
	"os"
)

type Setting struct {
	Database_username string `env:"BOOKSTORE_USER,required"`
	Database_password string `env:"BOOKSTORE_PASSWORD,required"`
	Database_Host     string `env:"BOOKSTORE_HOST,default=localhost"`
	Database_Port     string `env:"BOOKSTORE_PORT,required"`
	Database_name     string `env:"BOOKSTORE_NAME,required"`
}

func getSetting() *Setting {
	return &Setting{
		Database_username: os.Getenv("BOOKSTORE_USER"),
		Database_password: os.Getenv("BOOKSTORE_PASSWORD"),
		Database_Host:     os.Getenv("BOOKSTORE_HOST"),
		Database_Port:     os.Getenv("BOOKSTORE_PORT"),
		Database_name:     os.Getenv("BOOKSTORE_NAME"),
	}
}
