package main

import (
	"fmt"
	"os"
)

var dbUrl string

func init() {
	fmt.Println("Runing init function")
	os.Setenv("DB_URL", "Test DB_URL")
	dbUrl = os.Getenv("DB_URL")
	if dbUrl == "" {
		dbUrl = "Default DB URL string"
	}
}

func main() {
	fmt.Println("Running main function")
	fmt.Println("DB URL : ", dbUrl)
}
