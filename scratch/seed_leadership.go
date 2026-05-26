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

	leadership := []models.Leadership{
		{
			Name:           "Mr. Anand Kumar Tripathi",
			Role:           "Chairman, Krishna Engineering College & Vice-Chairman, KES",
			Department:     "Management",
			Image:          "/leadership/anandKumar.jpeg",
			Qualification:  "Graduate",
			Experience:     "15+ Years",
			Email:          "chairman@krishnacollege.edu",
			Phone:          "+91-788-2286662",
			Specialization: "Governance & Strategy",
			Achievements:   datatypes.JSON(`["Chairman of Krishna Engineering College", "Vice-Chairman of Krishna Education Society", "Driving academic excellence and innovation in education"]`),
			Bio:            "A visionary leader shaping the growth of Krishna Engineering College and contributing significantly as Vice-Chairman of Krishna Education Society.",
			Priority:       1,
		},
		{
			Name:           "Mr. M. M. Tripathi",
			Role:           "Chairman, Krishna Education Society",
			Department:     "Management",
			Image:          "https://www.kecbhilai.com/images/MANAGEMENT%20&%20HIGHER%20AUTHORITIES/mmtripathi.jpg",
			Qualification:  "Graduate",
			Experience:     "20+ Years",
			Email:          "president@krishnacollege.edu",
			Phone:          "+91-788-2286661",
			Specialization: "Academic Governance",
			Achievements:   datatypes.JSON(`["Chairman of Krishna Education Society", "Promoting quality education and values", "Champion of community-driven education"]`),
			Bio:            "Providing strong leadership and guidance as Chairman of Krishna Education Society, fostering an environment of learning and growth.",
			Priority:       2,
		},
		{
			Name:           "Mr. Pramod Kumar Tripathi",
			Role:           "Secretary, Krishna Education Society",
			Department:     "Management",
			Image:          "https://www.kecbhilai.com/images/MANAGEMENT%20&%20HIGHER%20AUTHORITIES/pramodtripathi.jpg",
			Qualification:  "Graduate",
			Experience:     "12+ Years",
			Email:          "secretary@krishnacollege.edu",
			Phone:          "+91-788-2286663",
			Specialization: "Administration & Development",
			Achievements:   datatypes.JSON(`["Secretary of Krishna Education Society", "Strengthening academic and cultural initiatives", "Committed to holistic student development"]`),
			Bio:            "Dedicated to the mission of Krishna Education Society, ensuring high standards in education and administration.",
			Priority:       3,
		},
		{
			Name:           "Dr. Ajay Tiwari",
			Role:           "Principal",
			Department:     "Administration",
			Image:          "https://image-static.collegedunia.com/public/image/19-09:08-Ajay_Tiwari_01.jpeg",
			Qualification:  "PhD Mechanical (Robotics)",
			Experience:     "33 years",
			Email:          "drajay2806@gmail.com",
			Phone:          "9893510942",
			Specialization: "Mechanical (Robotics)",
			Achievements:   datatypes.JSON(`["21 Publications (10 National, 7 International, 8 Journals/Conferences)", "2 PhD Scholars Guided", "1 Patent Registered", "Area of Specialization: Mechanical (Robotics)"]`),
			Bio:            "At Krishna Engineering College, our mission is to foster innovation, excellence, and ethical leadership in our students. We are committed to providing world-class education and nurturing future-ready engineers who will make a positive impact on society.",
			Priority:       4,
		},
		{
			Name:           "Ash Kumar Soni",
			Role:           "HOD",
			Department:     "Mechanical Engineering",
			Image:          "https://images.unsplash.com/photo-1522075469751-3a6694fb2f61?ixlib=rb-4.0.3&auto=format&fit=crop&w=400&q=80",
			Qualification:  "M.Tech (Production Engg), B.E (Mechanical Engg)",
			Experience:     "14 years",
			Email:          "hod.mech.ash@kec.edu",
			Phone:          "+91 9876543214",
			Specialization: "Production Engineering",
			Achievements:   datatypes.JSON(`["5 publications", "Guided 5 student projects"]`),
			Bio:            "An experienced mechanical engineering professional with expertise in production engineering. Committed to academic excellence and student development.",
			Priority:       5,
		},
	}

	for _, member := range leadership {
		var existing models.Leadership
		if err := config.DB.Where("name = ?", member.Name).First(&existing).Error; err != nil {
			config.DB.Create(&member)
			log.Printf("Created Leadership: %s\n", member.Name)
		} else {
			config.DB.Model(&existing).Updates(map[string]interface{}{
				"role":           member.Role,
				"department":     member.Department,
				"image":          member.Image,
				"qualification":  member.Qualification,
				"experience":     member.Experience,
				"email":          member.Email,
				"phone":          member.Phone,
				"specialization": member.Specialization,
				"achievements":   member.Achievements,
				"bio":            member.Bio,
				"priority":       member.Priority,
			})
			log.Printf("Updated Leadership: %s\n", member.Name)
		}
	}

	log.Println("Leadership seeding complete!")
}
