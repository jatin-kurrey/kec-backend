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

	notices := []models.Notice{
		{
			Title:     "Mid-Semester Examination Schedule Released",
			Content:   "The schedule for mid-semester examinations has been published. All students are required to check their respective department notices for exact dates and timings. Examinations will commence from September 5, 2025.",
			Type:      "academic",
			Link:      "#",
			Important: true,
			IsActive:  true,
		},
		{
			Title:     "Placement Drive – Infosys & TCS Recruitment",
			Content:   "Infosys and TCS will be conducting campus recruitment drives on September 10-12, 2025. Eligible students from CSE, IT, ECE, and EEE departments must register by August 30, 2025.",
			Type:      "placement",
			Link:      "#",
			Important: true,
			IsActive:  true,
		},
		{
			Title:     "Holiday Notice – Independence Day Celebration",
			Content:   "College will remain closed on August 15, 2025, on account of Independence Day. The celebration ceremony will be held on August 14, 2025, at the college grounds from 9:00 AM onwards.",
			Type:      "general",
			Link:      "#",
			Important: false,
			IsActive:  true,
		},
		{
			Title:     "Workshop on AI & Machine Learning",
			Content:   "A two-day workshop on Artificial Intelligence and Machine Learning will be conducted by industry experts from TechMind Solutions on August 25-26, 2025. Limited seats available.",
			Type:      "event",
			Link:      "#",
			Important: false,
			IsActive:  true,
		},
		{
			Title:     "Last Date for Scholarship Applications",
			Content:   "The last date for submitting scholarship applications for the academic year 2025-26 is August 20, 2025. Students are advised to submit their forms to the administration office before the deadline.",
			Type:      "student",
			Link:      "#",
			Important: true,
			IsActive:  true,
		},
		{
			Title:     "Inter-College Sports Tournament",
			Content:   "Annual inter-college sports tournament will be held from September 15-20, 2025. Students interested in participating should register with the sports department by September 5, 2025.",
			Type:      "event",
			Link:      "#",
			Important: false,
			IsActive:  true,
		},
	}

	for _, notice := range notices {
		var existing models.Notice
		if err := config.DB.Where("title = ?", notice.Title).First(&existing).Error; err != nil {
			config.DB.Create(&notice)
			log.Printf("Created Notice: %s\n", notice.Title)
		} else {
			config.DB.Model(&existing).Updates(map[string]interface{}{
				"content":   notice.Content,
				"type":      notice.Type,
				"link":      notice.Link,
				"important": notice.Important,
				"is_active": true,
			})
			log.Printf("Updated Notice: %s\n", notice.Title)
		}
	}

	log.Println("Notice seeding complete!")
}
