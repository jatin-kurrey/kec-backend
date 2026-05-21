package main

import (
	"kec-backend/internal/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load("../.env")
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.Model(&models.CampusFacility{}).Update("is_active", true)
	db.Model(&models.CampusStat{}).Update("is_active", true)
	log.Println("All facilities and stats set to active!")
}
