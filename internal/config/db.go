package config

import (
	"log"
	"os"

	"kec-backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	url := os.Getenv("DB_URL")
	if url == "" {
		log.Fatal("DB_URL is not set in .env")
	}

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection established")

	// Auto-migrate models
	err = db.AutoMigrate(
		&models.Admin{},
		&models.Application{},
		&models.Notice{},
		&models.Department{},
		&models.Gallery{},
		&models.Exam{},
		&models.Question{},
		&models.ExamResponse{},
		&models.Leadership{},
		&models.Faculty{},
		&models.Course{},
		&models.ResearchArea{},
		&models.ResearchProject{},
		&models.ResearchFacility{},
		&models.ResearchStat{},
		&models.AdmissionGuide{},
		&models.AdmissionStep{},
		&models.AdmissionEligibility{},
		&models.AdmissionDocument{},
		&models.AdmissionFee{},
		&models.CampusFacility{},
		&models.CampusStat{},
		&models.Alumni{},
		&models.AlumniStat{},
		&models.PlacementStat{},
		&models.Recruiter{},
		&models.PlacementTestimonial{},
		&models.PressMedia{},
	)
	if err != nil {
		log.Fatal("Failed to auto-migrate:", err)
	}

	DB = db
}
