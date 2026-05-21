package main

import (
	"log"
	"kec-backend/internal/config"
	"kec-backend/internal/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	godotenv.Load("../../.env")
	config.ConnectDB()

	log.Println("Dropping departments table...")
	config.DB.Migrator().DropTable(&models.Department{})
	
	log.Println("Re-migrating departments table...")
	config.DB.AutoMigrate(&models.Department{})
	
	log.Println("Table recreated successfully")
}
