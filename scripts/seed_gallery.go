package main

import (
	"log"
	"os"
	"time"

	"kec-backend/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	items := []models.Gallery{
		{
			Title:       "Annual Cultural Fest 2024",
			Category:    "Cultural",
			Description: "A vibrant celebration of art, music, and dance featuring students from all departments. The highlight was the inter-collegiate dance competition.",
			ImageURL:    "https://images.unsplash.com/photo-1514525253361-bee438d59ef7",
			Date:        time.Now().AddDate(0, -2, 0),
			Time:        "10:00 AM",
			Location:    "Main Auditorium",
			Likes:       245,
			Downloads:   89,
			Attendees:   1200,
			Color:       "#EC4899",
		},
		{
			Title:       "Tech Innovation Summit",
			Category:    "Academic",
			Description: "Showcasing the latest engineering projects and research breakthroughs by our brilliant students. Includes AI, Robotics, and Sustainable Tech.",
			ImageURL:    "https://images.unsplash.com/photo-1540575861501-7ad0582373f2",
			Date:        time.Now().AddDate(0, -1, -5),
			Time:        "09:00 AM",
			Location:    "Science Block",
			Likes:       182,
			Downloads:   124,
			Attendees:   850,
			Color:       "#3B82F6",
		},
		{
			Title:       "Traditional Day Celebrations",
			Category:    "Religious",
			Description: "Celebrating our diverse heritage through traditional attire and cultural performances. Unity in diversity was the theme of the day.",
			ImageURL:    "https://images.unsplash.com/photo-1511795409834-ef04bbd61622",
			Date:        time.Now().AddDate(0, -4, 0),
			Time:        "11:00 AM",
			Location:    "Central Plaza",
			Likes:       312,
			Downloads:   56,
			Attendees:   1500,
			Color:       "#F59E0B",
		},
		{
			Title:       "Robotics Workshop 2024",
			Category:    "Academic",
			Description: "Hands-on training session on advanced robotics and automation technologies. Students built and programmed their own mini-bots.",
			ImageURL:    "https://images.unsplash.com/photo-1485827404703-89b55fcc595e",
			Date:        time.Now().AddDate(0, 0, -10),
			Time:        "02:00 PM",
			Location:    "E-Block Lab",
			Likes:       156,
			Downloads:   210,
			Attendees:   300,
			Color:       "#10B981",
		},
		{
			Title:       "Independence Day Parade",
			Category:    "National",
			Description: "Annual flag hoisting ceremony and patriotic parade by the NCC cadets and faculty members.",
			ImageURL:    "https://images.unsplash.com/photo-1532375810709-75b1da00537c",
			Date:        time.Now().AddDate(0, -9, 0),
			Time:        "08:30 AM",
			Location:    "Sports Ground",
			Likes:       420,
			Downloads:   115,
			Attendees:   2500,
			Color:       "#EA580C",
		},
	}

	for _, item := range items {
		db.Create(&item)
	}

	log.Println("Successfully seeded gallery items!")
}
