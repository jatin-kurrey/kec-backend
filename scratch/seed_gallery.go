package main

import (
	"log"
	"kec-backend/internal/config"
	"kec-backend/internal/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	config.ConnectDB()

	gallery := []models.Gallery{
		{
			Title:       "Hanuman Jayanti Celebration 2025",
			Category:    "Religious",
			ImageURL:    "https://scontent.fbho3-1.fna.fbcdn.net/v/t39.30808-6/441305399_122752391367473984_3686828832087925275_n.jpg",
			Description: "Hanuman Jayanti was celebrated with great devotion and enthusiasm at Krishna Engineering College. Both students and staff collaborated to organize a vibrant and spiritual event.",
		},
		{
			Title:       "Gulal 2.0: Freshers & Farewell Party 2025",
			Category:    "Cultural",
			ImageURL:    "https://scontent.fbho3-1.fna.fbcdn.net/v/t39.30808-6/441230145_122752391347473984_7869424194947334913_n.jpg",
			Description: "A colorful and joyful celebration as Krishna Engineering College hosted Gulal 2.0, combining Freshers' Welcome and Farewell party. Students celebrated Holi and bid farewell to seniors with music, dance, and fun.",
		},
		{
			Title:       "Happy 79th Independence Day!",
			Category:    "National",
			ImageURL:    "/events/event4.jpg",
			Description: "At Krishna Engineering College, we salute the spirit of freedom, innovation, and resilience that defines our great nation. May we continue to build a stronger India — with knowledge, technology, and unity at the core.",
		},
		{
			Title:       "Welcome, B.Tech Batch of 2025–26!",
			Category:    "Academic",
			ImageURL:    "/events/event3.jpeg",
			Description: "Induction Program 2025–26🔥🔥 We're thrilled to welcome the newest members of our Krishna Engineering College family! Join us in the Seminar Hall for an inspiring session with Prof. Rajiv Prakash, Director of IIT Bhilai, as our Chief Guest.",
		},
		{
			Title:       "Cultural Festivals",
			Category:    "Cultural",
			ImageURL:    "/campus/Cultural.jpeg",
			Description: "Celebrate diversity with cultural events and international festivals showcasing the rich heritage and traditions from across the globe.",
		},
	}

	for _, item := range gallery {
		var existing models.Gallery
		if err := config.DB.Where("title = ?", item.Title).First(&existing).Error; err != nil {
			config.DB.Create(&item)
			log.Printf("Created Gallery Item: %s\n", item.Title)
		} else {
			config.DB.Model(&existing).Updates(map[string]interface{}{
				"description": item.Description,
				"category":    item.Category,
				"image_url":   item.ImageURL,
			})
			log.Printf("Updated Gallery Item: %s\n", item.Title)
		}
	}

	log.Println("Gallery seeding complete!")
}
