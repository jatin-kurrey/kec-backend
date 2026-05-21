package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")
	fmt.Printf("DB_URL: '%s'\n", os.Getenv("DB_URL"))
	fmt.Printf("DATABASE_URL: '%s'\n", os.Getenv("DATABASE_URL"))
}
