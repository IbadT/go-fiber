package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file %s", err.Error())
	}

	fmt.Println("Hello Golang + fiber")
}
