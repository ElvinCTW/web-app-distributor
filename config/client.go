package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"reflect"
	"sync"
)

const devPath = "dev.env"

var (
	data = &Data{}
	once sync.Once
)

type Data struct {
	DATABASE_URL string
	DATABASE     string
	LOG_LEVEL    string
}

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

	insert()
}

func insert() {
	typeofConfig := reflect.TypeOf(*data)

	for i := 0; i < typeofConfig.NumField(); i++ {
		fName := typeofConfig.Field(i).Name
		v := os.Getenv(fName)
		value := reflect.ValueOf(v)
		reflect.ValueOf(data).Elem().FieldByName(fName).Set(value)
		fmt.Println(fName, "=", value)
	}
}
