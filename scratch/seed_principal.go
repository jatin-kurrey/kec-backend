package main

import (
	"log"
	"kec-backend/internal/config"
	"kec-backend/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/datatypes"
)

func main() {
	godotenv.Load()
	godotenv.Load("../../.env")
	config.ConnectDB()

	principal := models.Leadership{
		Name:           "Dr. Ajay Tiwari",
		Role:           "Principal",
		Department:     "Administration",
		Image:          "https://image-static.collegedunia.com/public/image/19-09:08-Ajay_Tiwari_01.jpeg",
		Qualification:  "PhD Mechanical (Robotics)",
		Experience:     "33 years",
		Email:          "drajay2806@gmail.com",
		Phone:          "9893510942",
		Specialization: "Mechanical (Robotics)",
		Achievements:   datatypes.JSON(`["21 Publications (10 National, 7 International, 8 Journals/Conferences)", "2 PhD Scholars Guided", "1 Patent Registered"]`),
		Bio:            "At Krishna Engineering College, our mission is to foster innovation, excellence, and ethical leadership in our students. We are committed to providing world-class education and nurturing future-ready engineers who will make a positive impact on society. Our institution stands as a beacon of knowledge, where we blend traditional values with modern technological advancements to create a holistic learning environment. Our dedicated faculty, state-of-the-art infrastructure, and industry partnerships ensure that our students receive not just education, but an experience that shapes their character and professional capabilities. We take pride in seeing our graduates excel in diverse fields across the globe, carrying forward the legacy of excellence that defines Krishna Engineering College. I invite you to join our vibrant academic community and embark on a journey of discovery, innovation, and personal growth.",
		Priority:       0,
	}

	var existing models.Leadership
	if err := config.DB.Where("role = ?", "Principal").First(&existing).Error; err != nil {
		config.DB.Create(&principal)
		log.Printf("Created Principal: %s\n", principal.Name)
	} else {
		config.DB.Model(&existing).Updates(principal)
		log.Printf("Updated Principal: %s\n", principal.Name)
	}

	log.Println("Principal seeding complete")
}
