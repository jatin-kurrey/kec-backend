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
			Bio:            "At Krishna Engineering College, our mission is to foster innovation, excellence, and ethical leadership in our students. We are committed to providing world-class education and nurturing future-ready engineers who will make a positive impact on society. Our institution stands as a beacon of knowledge, where we blend traditional values with modern technological advancements to create a holistic learning environment. Our dedicated faculty, state-of-the-art infrastructure, and industry partnerships ensure that our students receive not just education, but an experience that shapes their character and professional capabilities. We take pride in seeing our graduates excel in diverse fields across the globe, carrying forward the legacy of excellence that defines Krishna Engineering College. I invite you to join our vibrant academic community and embark on a journey of discovery, innovation, and personal growth.",
			Priority:       1,
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
			Priority:       2,
		},
		{
			Name:           "Prabhat Kumar Patel",
			Role:           "HOD",
			Department:     "Civil Engineering",
			Image:          "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=400&q=80",
			Qualification:  "B.E (Civil), M.Tech (CTM - Civil Engineering), MBA (Finance), MBA (HR & Marketing), Pursuing AMIE",
			Experience:     "7 years",
			Email:          "hod.civil.prabhat@kec.edu",
			Phone:          "+91 9876543215",
			Specialization: "Construction Technology and Management",
			Achievements:   datatypes.JSON(`["6 research papers published in UGC-approved journals", "Presented papers in peer-reviewed journals", "Best Young Faculty Award - Novel Research Academy", "Best Research Scholar Award - Bhartiya Vikas Sansthan", "Bharatiya Gaurav Samman - Bhartiya Kala Sanskriti Academy", "Best Young Researcher Award - Institute of Scholars", "Top 10 Motivated Faculty of India - Engineering Graphics with Timoshenko", "Featured on Health & Success Magazine cover (June 2020)"]`),
			Bio:            "A highly accomplished academician with multiple qualifications and numerous awards. Passionate about research and student mentorship.",
			Priority:       3,
		},
		{
			Name:           "Dr. Joy Sonashalol",
			Role:           "HOD",
			Department:     "Computer Science & Engineering",
			Image:          "https://images.unsplash.com/photo-1522075469751-3a6694fb2f61?ixlib=rb-4.0.3&auto=format&fit=crop&w=400&q=80",
			Qualification:  "Ph.D. in Artificial Intelligence",
			Experience:     "20 years",
			Email:          "hod.cse@kec.edu",
			Phone:          "+91 9876543212",
			Specialization: "Artificial Intelligence, Machine Learning, Data Science",
			Achievements:   datatypes.JSON(`["Developed AI-driven learning platforms", "Author of 3 textbooks in Machine Learning and AI", "Consultant for multiple software product companies"]`),
			Bio:            "A leading academician in computer science, known for pioneering research in artificial intelligence and machine learning. Passionate about shaping future tech leaders through innovation and mentorship.",
			Priority:       4,
		},
		{
			Name:           "Mr. Tarachand Sahu",
			Role:           "HOD",
			Department:     "Electrical Engineering",
			Image:          "https://images.unsplash.com/photo-1522075469751-3a6694fb2f61?ixlib=rb-4.0.3&auto=format&fit=crop&w=400&q=80",
			Qualification:  "B.E. (Electrical & Electronics Engineering), M.Tech (Electrical Engineering)",
			Experience:     "10 years",
			Email:          "tarachand.sahu@kec.edu",
			Phone:          "+91 9876543216",
			Specialization: "Electrical Machines, Power Systems",
			Achievements:   datatypes.JSON(`["Published 2 research papers in reputed journals"]`),
			Bio:            "Mr. Tarachand Sahu is an Assistant Professor in the Department of Electrical Engineering. His teaching and academic interests lie in the areas of power systems and electrical machines. He is committed to student development and technical excellence.",
			Priority:       5,
		},
	}

	for _, member := range leadership {
		var existing models.Leadership
		if err := config.DB.Where("name = ?", member.Name).First(&existing).Error; err != nil {
			config.DB.Create(&member)
			log.Printf("Created Leadership: %s\n", member.Name)
		} else {
			// Use map to ensure all fields are updated (even if they were empty)
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

	log.Println("Leadership seeding (EXACT) complete!")
}
