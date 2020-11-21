package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

const path = "../.env"

var data *Data

func init() {
	if err := godotenv.Load(path); err != nil {
		fmt.Println("not used .env")
	}

	data = &Data{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Database:    os.Getenv("DATABASE"),
		LogLevel:    os.Getenv("LOG_LEVEL"),
	}
}

type Data struct {
	DatabaseURL string
	Database    string
	LogLevel    string
}

func Get() *Data {
	return data
}
