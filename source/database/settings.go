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
		//log.Fatal("Error loading .env file")
	}

	SetEnvParams(&Setting.Database_host, "BOOKSTORE_HOST", "localhost:3306")
	SetEnvParams(&Setting.Database_name, "BOOKSTORE_DATABASE", "mysql")
	SetEnvParams(&Setting.Database_username, "BOOKSTORE_USER", "root")
	SetEnvParams(&Setting.Database_password, "BOOKSTORE_PASSWORD", "")
	SetEnvParams(&Setting.Database_port, "BOOKSTORE_PORT", "3000")

	return nil
}
