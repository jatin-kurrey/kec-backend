package main

import (
	"kec-backend/internal/config"
	"kec-backend/internal/models"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgresql://neondb_owner:npg_KOXcR16raGnN@ep-empty-river-anw1hgfi-pooler.c-6.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	config.DB = db

	// Clean existing exams to start fresh as requested
	db.Exec("DELETE FROM questions")
	db.Exec("DELETE FROM exams")

	now := time.Now()

	// 1. KEC Super 40 Scholarship Test
	scholarshipTest := models.Exam{
		ID:                    uuid.New(),
		Title:                 "KEC Super 40 Scholarship Entrance 2026",
		Description:           "Official scholarship entrance examination for the prestigious Super 40 batch. High-fidelity assessment covering Physics, Chemistry, Mathematics and Logical Reasoning.",
		Duration:              120,
		NegativeMarking:       1.0,
		ShuffleQuestions:      true,
		BrowserLockdown:       true,
		ShowResultImmediately: false,
		StartTime:             now,
		EndTime:               now.Add(24 * 7 * time.Hour),
		IsActive:              true,
		Questions: []models.Question{
			{
				Text:          "Which of the following represents the correct unit of Electric Potential?",
				Type:          "MCQ",
				Options:       []byte(`["Volt", "Ampere", "Ohm", "Watt"]`),
				CorrectAnswer: "Volt",
				Points:        4,
			},
			{
				Text:          "The value of (sin^2 θ + cos^2 θ) is always equal to:",
				Type:          "MCQ",
				Options:       []byte(`["0", "1", "-1", "Infinity"]`),
				CorrectAnswer: "1",
				Points:        4,
			},
			{
				Text:          "Which gas is primarily responsible for the greenhouse effect?",
				Type:          "MCQ",
				Options:       []byte(`["Oxygen", "Nitrogen", "Carbon Dioxide", "Helium"]`),
				CorrectAnswer: "Carbon Dioxide",
				Points:        4,
			},
			{
				Text:          "Complete the series: 2, 4, 8, 16, ...",
				Type:          "MCQ",
				Options:       []byte(`["24", "30", "32", "64"]`),
				CorrectAnswer: "32",
				Points:        4,
			},
		},
	}

	// 2. KEC B.Tech Entrance Assessment (Lateral Entry)
	lateralTest := models.Exam{
		ID:                    uuid.New(),
		Title:                 "B.Tech Lateral Entry Entrance 2026",
		Description:           "Institutional entrance assessment for Diploma students seeking admission into 2nd year B.Tech. Focuses on Engineering Mechanics and Basic Programming.",
		Duration:              90,
		NegativeMarking:       0.25,
		ShuffleQuestions:      true,
		BrowserLockdown:       true,
		ShowResultImmediately: true,
		StartTime:             now,
		EndTime:               now.Add(24 * 30 * time.Hour),
		IsActive:              true,
		Questions: []models.Question{
			{
				Text:          "Newton's First Law of motion is also known as:",
				Type:          "MCQ",
				Options:       []byte(`["Law of Inertia", "Law of Acceleration", "Law of Action-Reaction", "Law of Gravity"]`),
				CorrectAnswer: "Law of Inertia",
				Points:        2,
			},
			{
				Text:          "In C language, what is the size of 'int' data type on a 32-bit system?",
				Type:          "MCQ",
				Options:       []byte(`["1 byte", "2 bytes", "4 bytes", "8 bytes"]`),
				CorrectAnswer: "4 bytes",
				Points:        2,
			},
		},
	}

	db.Create(&scholarshipTest)
	db.Create(&lateralTest)

	log.Println("Authentic Institutional Exams seeded successfully.")
}
