package main

import (
	"encoding/json"
	"kec-backend/internal/config"
	"kec-backend/internal/models"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgresql://neondb_owner:npg_KOXcR16raGnN@ep-empty-river-anw1hgfi-pooler.c-6.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	config.DB = db

	opt1, _ := json.Marshal([]string{"Tesla", "Weber", "Henry", "Farad"})
	optEmpty, _ := json.Marshal([]string{})
	optLogic, _ := json.Marshal([]string{"1", "2", "3", "4"})

	// 1. JEE Mains Practice Test
	exam1 := models.Exam{
		ID:                    uuid.New(),
		Title:                 "JEE Mains 2026: Physics & Math",
		Description:           "A comprehensive practice test covering Mechanics, Optics, and Calculus. Mandatory for Super 40 aspirants.",
		Duration:              60,
		NegativeMarking:       0.25,
		ShuffleQuestions:      true,
		BrowserLockdown:       true,
		ShowResultImmediately: true,
		StartTime:             time.Now(),
		EndTime:               time.Now().Add(720 * time.Hour),
		IsActive:              true,
		Questions: []models.Question{
			{
				Text:          "What is the SI unit of Magnetic Flux?",
				Type:          "MCQ",
				Options:       datatypes.JSON(opt1),
				CorrectAnswer: "Weber",
				Points:        4,
			},
			{
				Text:          "Find the value of lim(x->0) sin(x)/x.",
				Type:          "INTEGER",
				Options:       datatypes.JSON(optEmpty),
				CorrectAnswer: "1",
				Points:        4,
			},
		},
	}

	// 2. Logic & General Aptitude
	exam2 := models.Exam{
		ID:                    uuid.New(),
		Title:                 "Super 40: Aptitude & IQ Test",
		Description:           "Testing logical reasoning, pattern recognition, and numerical ability.",
		Duration:              30,
		NegativeMarking:       0,
		ShuffleQuestions:      false,
		BrowserLockdown:       false,
		ShowResultImmediately: false,
		StartTime:             time.Now(),
		EndTime:               time.Now().Add(720 * time.Hour),
		IsActive:              true,
		Questions: []models.Question{
			{
				Text:          "If A = 1, B = 2, what is C?",
				Type:          "MCQ",
				Options:       datatypes.JSON(optLogic),
				CorrectAnswer: "3",
				Points:        1,
			},
			{
				Text:          "What comes next in the series: 2, 4, 8, 16, ...?",
				Type:          "INTEGER",
				Options:       datatypes.JSON(optEmpty),
				CorrectAnswer: "32",
				Points:        2,
			},
		},
	}

	db.Create(&exam1)
	db.Create(&exam2)
	log.Println("Sample exams created successfully!")
}
