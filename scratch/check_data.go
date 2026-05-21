package main

import (
	"fmt"
	"kec-backend/internal/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load("../.env")
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var count int64
	db.Model(&models.CampusFacility{}).Count(&count)
	fmt.Printf("Total Facilities: %d\n", count)

	var activeCount int64
	db.Model(&models.CampusFacility{}).Where("is_active = ?", true).Count(&activeCount)
	fmt.Printf("Active Facilities: %d\n", activeCount)

	var statsCount int64
	db.Model(&models.CampusStat{}).Count(&statsCount)
	fmt.Printf("Total Stats: %d\n", statsCount)
}
