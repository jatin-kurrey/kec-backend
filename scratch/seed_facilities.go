package main

import (
	"encoding/json"
	"fmt"
	"kec-backend/internal/config"
	"kec-backend/internal/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load()
	}
	if err != nil {
		log.Println("No .env file found")
	}

	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	config.DB = db

	// Auto migrate to ensure tables exist
	db.AutoMigrate(&models.CampusFacility{}, &models.CampusStat{})

	// Seed Stats
	stats := []models.CampusStat{
		{Label: "Laboratories", Value: "49+", Icon: "Beaker", Color: "blue", SortOrder: 1},
		{Label: "Departments", Value: "6", Icon: "Building", Color: "green", SortOrder: 2},
		{Label: "Essential Amenities", Value: "18", Icon: "CheckCircle", Color: "purple", SortOrder: 3},
		{Label: "Desirable Amenities", Value: "12", Icon: "Wifi", Color: "orange", SortOrder: 4},
		{Label: "Sq.M. Total Area", Value: "3283+", Icon: "MapPin", Color: "red", SortOrder: 5},
	}

	for _, s := range stats {
		db.Where(models.CampusStat{Label: s.Label}).FirstOrCreate(&s)
	}

	// Seed Amenities
	essential := []string{
		"Stand alone language laboratory", "Potable water supply", "Electric supply",
		"Sewage disposal", "Telephone and fax", "Vehicle Parking",
		"Institution website", "Barrier free built environment", "Safety provisions",
		"Digital Library", "Classification of books", "NPTEL facility",
		"General insurance", "Motorised Road", "Notice boards",
		"Medical Facilities", "Grievance Redressal Committee", "Anti-sexual harassment Committee",
	}

	for i, name := range essential {
		fac := models.CampusFacility{
			Name:      name,
			Category:  "Essential",
			Available: true,
			SortOrder: i,
			IsActive:  true,
		}
		db.Where(models.CampusFacility{Name: name, Category: "Essential"}).FirstOrCreate(&fac)
	}

	desirable := []string{
		"Alumni Association", "Industry Institute Interaction", "Placement and Training",
		"Back up Electric supply", "ERP Software", "Transport facility",
		"Post/Bank facility/ATM", "CCTV System", "Staff quarters",
		"Display of courses", "Public announcement system", "Group insurance",
	}

	for i, name := range desirable {
		fac := models.CampusFacility{
			Name:      name,
			Category:  "Desirable",
			Available: true,
			SortOrder: i,
			IsActive:  true,
		}
		db.Where(models.CampusFacility{Name: name, Category: "Desirable"}).FirstOrCreate(&fac)
	}

	// Infrastructure
	infra := []struct {
		Name     string
		Desc     string
		Icon     string
		Features []string
	}{
		{
			Name: "Central Library",
			Desc: "Extensive collection of books, journals, and digital resources with NPTEL facility",
			Icon: "Library",
			Features: []string{"Digital Resources", "NPTEL Facility", "Reading Hall"},
		},
		{
			Name: "Computing Facilities",
			Desc: "High-speed internet with Wi-Fi connectivity across campus and computer labs",
			Icon: "Cpu",
			Features: []string{"High Speed WiFi", "Advanced Labs", "24/7 Access"},
		},
		{
			Name: "Safety & Security",
			Desc: "24/7 security, CCTV surveillance, and fire safety systems throughout campus",
			Icon: "Shield",
			Features: []string{"CCTV Surveillance", "Fire Safety", "Security Guard"},
		},
	}

	for i, item := range infra {
		featuresJSON, _ := json.Marshal(item.Features)
		fac := models.CampusFacility{
			Name:        item.Name,
			Description: item.Desc,
			Category:    "Infrastructure",
			Icon:        item.Icon,
			Features:    datatypes.JSON(featuresJSON),
			SortOrder:   i,
			IsActive:    true,
		}
		db.Where(models.CampusFacility{Name: item.Name, Category: "Infrastructure"}).FirstOrCreate(&fac)
	}

	fmt.Println("Campus Facilities seeded successfully!")
}
