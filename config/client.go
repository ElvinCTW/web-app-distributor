package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"sync"
)

const devPath = "dev.env"

var (
	data *Data
	once sync.Once
)

func Get() *Data {
	once.Do(func() {
		load()
	})
	return data
}

func load() {
	if err := godotenv.Load(devPath); err != nil {
		fmt.Println("not used .env")
		panic(err)
	}

	data = &Data{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Database:    os.Getenv("DATABASE"),
		LogLevel:    os.Getenv("LOG_LEVEL"),
	}

	fmt.Println("DATABASE_URL:", data.DatabaseURL)
	fmt.Println("DATABASE:", data.Database)
	fmt.Println("LOG_LEVEL:", data.LogLevel)
}

type Data struct {
	DatabaseURL string
	Database    string
	LogLevel    string
}
