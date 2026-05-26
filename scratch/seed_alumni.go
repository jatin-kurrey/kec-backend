package main

import (
	"log"
	"kec-backend/internal/config"
	"kec-backend/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/datatypes"
)

func main() {
	godotenv.Load(".env")
	config.ConnectDB()

	alumniStories := []models.Alumni{
		{
			Name:            "SAZIYA NAAZ",
			Batch:           "2025",
			Branch:          "CSE",
			Degree:          "B.Tech CSE",
			Location:        "Raipur",
			Company:         "Codenicely",
			CurrentPosition: "Software Engineer",
			Story:           "My journey at KEC gave me confidence to crack Codenicely interviews.",
			ImageURL:        "/alumini/SAZIYA NAAZ.png",
			Achievements:    datatypes.JSON(`["Secured placement in Codenicely before graduation", "Excellent academic track record"]`),
			IsFeatured:      true,
		},
		{
			Name:            "TARA CHAND DEWANGAN",
			Batch:           "2025",
			Branch:          "CSE",
			Degree:          "B.Tech CSE",
			Location:        "Raipur",
			Company:         "Sthanve Software",
			CurrentPosition: "Software Engineer",
			Story:           "Faculty guidance helped me grab a role at Sthanve Software.",
			ImageURL:        "/alumini/TARA CHAND DEWANGAN.png",
			Achievements:    datatypes.JSON(`["Placed at Sthanve Software during campus drive"]`),
			IsFeatured:      false,
		},
		{
			Name:            "NIDHI CHANDRAWANDI",
			Batch:           "2024",
			Branch:          "CSE",
			Degree:          "B.Tech CSE",
			Location:        "Raipur",
			Company:         "Nullclass Technology",
			CurrentPosition: "Engineer",
			Story:           "Hands-on projects at KEC shaped my problem-solving skills.",
			ImageURL:        "/alumini/nidhi chandrawanshi.png",
			Achievements:    datatypes.JSON(`["Started my professional journey at Nullclass Technology"]`),
			IsFeatured:      false,
		},
		{
			Name:            "BHUPENDRA",
			Batch:           "2022",
			Branch:          "Civil",
			Degree:          "B.Tech Civil",
			Location:        "Bhilai",
			Company:         "Bodex System Pvt Ltd",
			CurrentPosition: "Engineer",
			Story:           "The Civil department gave me industry-level exposure.",
			ImageURL:        "/alumini/bhupendra.png",
			Achievements:    datatypes.JSON(`["Contributed to key projects at Bodex System Pvt Ltd"]`),
			IsFeatured:      false,
		},
		{
			Name:            "NUPUR",
			Batch:           "2022",
			Branch:          "CSE",
			Degree:          "B.Tech CSE",
			Location:        "Bhilai",
			Company:         "Linterbiz Consulting Pvt Ltd",
			CurrentPosition: "Consultant",
			Story:           "The analytical skills I developed at KEC helped me excel in consulting.",
			ImageURL:        "/alumini/nupur.png",
			Achievements:    datatypes.JSON(`["Joined Linterbiz Consulting Pvt Ltd as a key team member"]`),
			IsFeatured:      false,
		},
		{
			Name:            "TRILOK DHRUW",
			Batch:           "2022",
			Branch:          "EE",
			Degree:          "B.Tech EE",
			Location:        "Bhilai",
			Company:         "Prixso Software",
			CurrentPosition: "Software Developer",
			Story:           "KEC's strong technical foundation helped me pivot into the IT industry confidently.",
			ImageURL:        "/alumini/trilok dhruw.png",
			Achievements:    datatypes.JSON(`["Transitioned from Electrical Engineering to a successful IT career"]`),
			IsFeatured:      false,
		},
		{
			Name:            "YURAJ KHARE",
			Batch:           "2020",
			Branch:          "CSE",
			Degree:          "B.Tech CSE",
			Location:        "Bhilai",
			Company:         "Empyra Software Solutions",
			CurrentPosition: "Software Engineer",
			Story:           "KEC gave me the technical foundation and confidence to thrive in the IT industry.",
			ImageURL:        "/alumini/yuraj Khare.png",
			Achievements:    datatypes.JSON(`["Started career as a full-stack developer at Empyra"]`),
			IsFeatured:      false,
		},
		{
			Name:            "SANJAY KUMAR",
			Batch:           "2022",
			Branch:          "EE",
			Degree:          "B.Tech EE",
			Location:        "Bhilai",
			Company:         "Kalpataru Power Transmission Ltd.",
			CurrentPosition: "Project Engineer",
			Story:           "The practical training and industry exposure at KEC laid the foundation for my career.",
			ImageURL:        "/alumini/sanjay kumar.png",
			Achievements:    datatypes.JSON(`["Successfully contributed to multiple national infrastructure projects"]`),
			IsFeatured:      false,
		},
	}

	for _, alum := range alumniStories {
		var existing models.Alumni
		if err := config.DB.Where("name = ?", alum.Name).First(&existing).Error; err != nil {
			config.DB.Create(&alum)
			log.Printf("Created Alumni Story: %s\n", alum.Name)
		} else {
			config.DB.Model(&existing).Updates(map[string]interface{}{
				"batch":            alum.Batch,
				"branch":           alum.Branch,
				"degree":           alum.Degree,
				"location":         alum.Location,
				"company":          alum.Company,
				"current_position": alum.CurrentPosition,
				"story":            alum.Story,
				"image_url":        alum.ImageURL,
				"achievements":     alum.Achievements,
				"is_featured":      alum.IsFeatured,
			})
			log.Printf("Updated Alumni Story: %s\n", alum.Name)
		}
	}

	// Seed stats
	stats := []models.AlumniStat{
		{
			Label: "Alumni Worldwide",
			Value: "50000",
			Icon:  "Users",
			Color: "#22c55e",
		},
		{
			Label: "Career Satisfaction",
			Value: "90",
			Icon:  "TrendingUp",
			Color: "#16a34a",
		},
		{
			Label: "Average Salary",
			Value: "85",
			Icon:  "Award",
			Color: "#15803d",
		},
		{
			Label: "Countries Represented",
			Value: "150",
			Icon:  "MapPin",
			Color: "#166534",
		},
	}

	for _, stat := range stats {
		var existing models.AlumniStat
		if err := config.DB.Where("label = ?", stat.Label).First(&existing).Error; err != nil {
			config.DB.Create(&stat)
			log.Printf("Created Alumni Stat: %s\n", stat.Label)
		} else {
			config.DB.Model(&existing).Updates(map[string]interface{}{
				"value": stat.Value,
				"icon":  stat.Icon,
				"color": stat.Color,
			})
			log.Printf("Updated Alumni Stat: %s\n", stat.Label)
		}
	}

	log.Println("Alumni seeding complete!")
}
