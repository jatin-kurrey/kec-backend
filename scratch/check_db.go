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
	var count int64
	config.DB.Model(&models.Leadership{}).Count(&count)
	log.Printf("Leadership count: %d", count)
	
	var members []models.Leadership
	config.DB.Find(&members)
	for _, m := range members {
		log.Printf("Member: %s (%s)", m.Name, m.Role)
	}
}
