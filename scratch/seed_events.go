package main

import (
	"log"
	"kec-backend/internal/config"
	"kec-backend/internal/models"
	"time"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	config.ConnectDB()

	events := []models.Event{
		{
			Title:            "Happy 79th Independence Day!",
			Date:             parseDate("2025-08-15"),
			Time:             "09:00 AM",
			Location:         "Krishna Engineering College, Bhilai",
			Category:         "National Celebration",
			Description:      "At Krishna Engineering College, we salute the spirit of freedom, innovation, and resilience that defines our great nation. May we continue to build a stronger India — with knowledge, technology, and unity at the core.",
			ImageURL:         "https://scontent.fidr4-2.fna.fbcdn.net/v/t39.30808-6/533121663_1160483856099139_6219316627004094283_n.jpg?_nc_cat=101&ccb=1-7&_nc_sid=833d8c&_nc_ohc=urUem3-V-SkQ7kNvwH-JRmS&_nc_oc=AdmeSApzH0VDlcFRcyCM1TuCmLK3DdwTiMlhKCJ0uJ-SHCAVMEi1-05XIZsJyrljUhi9e45a9m5AToH1dKoVZibM&_nc_zt=23&_nc_ht=scontent.fidr4-2.fna&_nc_gid=L4KidCdcsj9xTmeCOKTmQg&oh=00_AfYyoe-JcboXYnKF1ju1b4gtYYcOhgpAz0Tm3NXGmOQ7Sw&oe=68D9E2A4",
			Attendees:        500,
			Status:           "past",
			RegistrationLink: "#",
		},
		{
			Title:            "Welcome B.Tech Batch of 2025–26! 🎉",
			Date:             parseDate("2025-08-13"),
			Time:             "10:00 AM",
			Location:         "KEC, Bhilai",
			Category:         "Orientation",
			Description:      "Induction Program 2025–26🔥🔥\nWe’re thrilled to welcome the newest members of our Krishna Engineering College family! 💙 Join us in the Seminar Hall for an inspiring session with Prof. Rajiv Prakash, Director of IIT Bhilai, as our Chief Guest. Let’s celebrate the start of an exciting journey! 🚀",
			ImageURL:         "/events/event3.jpeg",
			Attendees:        400,
			Status:           "past",
			RegistrationLink: "#",
		},
		{
			Title:            "Campus Drive: Kalpataru Projects International Ltd.",
			Date:             parseDate("2025-09-25"),
			Time:             "11:00 AM",
			Location:         "KEC, Bhilai",
			Category:         "Placement Drive",
			Description:      "A proud day at KEC Bhilai as Kalpataru Projects International Ltd. conducted a successful campus drive. Final-year students had the opportunity to engage with recruiters from one of India’s leading infrastructure companies. A great step forward in building careers! 🏗️🎓",
			ImageURL:         "/events/event2.jpeg",
			Attendees:        200,
			Status:           "past",
			RegistrationLink: "#",
		},
	}

	for _, item := range events {
		var existing models.Event
		if err := config.DB.Where("title = ?", item.Title).First(&existing).Error; err != nil {
			config.DB.Create(&item)
			log.Printf("Created Event: %s\n", item.Title)
		} else {
			config.DB.Model(&existing).Updates(item)
			log.Printf("Updated Event: %s\n", item.Title)
		}
	}

	log.Println("Events seeding complete!")
}

func parseDate(dateStr string) time.Time {
	t, _ := time.Parse("2006-01-02", dateStr)
	return t
}
