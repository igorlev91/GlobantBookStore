package database

import (
	"os"

	"log"

	"github.com/joho/godotenv"
)

var Setting struct {
	Database_username string
	Database_password string
	Database_host     string
	Database_port     string
	Database_name     string

	Database_max_connection string
	Database_timeout        string
}

func SetEnvParams(settingDefault *string, env_variable string, default_value string) {
	var exists bool
	*settingDefault, exists = os.LookupEnv(env_variable)
	if !exists {
		*settingDefault = default_value
		log.Println(env_variable)
		return
	}
	log.Println(env_variable, *settingDefault)
}

func LoadEnv(path string) error {
	log.Println("Loading ", path)
	err := godotenv.Load(path)
	if err != nil {
		panic(err.Error())
	}

	SetEnvParams(&Setting.Database_host, "MYSQL_HOST", "bookstore_database")
	SetEnvParams(&Setting.Database_name, "MYSQL_DATABASE", "bookstore")
	SetEnvParams(&Setting.Database_username, "MYSQL_USER", "root")
	SetEnvParams(&Setting.Database_password, "MYSQL_PASSWORD", "")
	SetEnvParams(&Setting.Database_port, "SERVER_ADDRESS", "3336")

	SetEnvParams(&Setting.Database_max_connection, "DATABASE_MAX_CONN_COUNT", "3")
	SetEnvParams(&Setting.Database_timeout, "DATABASE_TIMEOUT", "50")

	return nil
}
