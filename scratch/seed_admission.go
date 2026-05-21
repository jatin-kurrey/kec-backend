package main

import (
	"log"
	"kec-backend/internal/config"
	"kec-backend/internal/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	config.ConnectDB()

	// Clear existing data
	config.DB.Exec("DELETE FROM admission_guides")
	config.DB.Exec("DELETE FROM admission_steps")
	config.DB.Exec("DELETE FROM admission_eligibilities")
	config.DB.Exec("DELETE FROM admission_documents")
	config.DB.Exec("DELETE FROM admission_fees")

	// 1. Seed Admission Guide
	guide := models.AdmissionGuide{
		Name:          "Mr. Durga Das",
		Position:      "Head of Admissions",
		Qualification: "Ph.D in Education Management, MBA",
		Experience:    "15+ years in academic administration",
		Email:         "admission.head@kecbhilai.edu.in",
		Phone:         "+91-9876543210",
		Image:         "https://images.unsplash.com/photo-1560250097-0b93528c311a?ixlib=rb-4.0.3&auto=format&fit=crop&w=200&q=80",
		Message:       "I welcome all prospective students to Krishna Engineering College. Our admission process is designed to be transparent and student-friendly. Feel free to reach out to me or my team for any assistance throughout your admission journey.",
	}
	config.DB.Create(&guide)

	// 2. Seed Admission Steps
	steps := []models.AdmissionStep{
		{Title: "Fill Online Application Form", Description: "Complete our user-friendly online application form with your details", Icon: "FileText", SortOrder: 1},
		{Title: "Submit Required Documents", Description: "Upload all necessary documents for verification", Icon: "Download", SortOrder: 2},
		{Title: "Appear for Entrance/Counseling", Description: "Participate in our admission process (if applicable)", Icon: "User", SortOrder: 3},
		{Title: "Pay Admission & Tuition Fee", Description: "Complete the payment process to secure your seat", Icon: "DollarSign", SortOrder: 4},
		{Title: "Confirm Enrollment", Description: "Receive your admission confirmation and welcome package", Icon: "CheckCircle", SortOrder: 5},
	}
	for _, s := range steps {
		config.DB.Create(&s)
	}

	// 3. Seed Eligibility
	eligibility := []models.AdmissionEligibility{
		{Criteria: "B.Tech: 10+2 with Physics, Chemistry, and Mathematics with minimum 45% aggregate (40% for SC/ST).", SortOrder: 1},
		{Criteria: "M.Tech: Bachelor's degree in Engineering/Technology with relevant specialization.", SortOrder: 2},
		{Criteria: "MBA: Graduation in any discipline with valid CAT/MAT/UPSEE score.", SortOrder: 3},
	}
	for _, e := range eligibility {
		config.DB.Create(&e)
	}

	// 4. Seed Documents
	documents := []string{
		"10th & 12th Marksheet/Certificate",
		"Graduation Marksheet (for PG courses)",
		"Transfer & Migration Certificate",
		"Passport Size Photographs",
		"Caste/Category Certificate (if applicable)",
		"Entrance Exam Score Card (if applicable)",
		"CG PET Score Card",
		"Domicile certificate",
		"Residence Certificate",
		"Freedom fighter certificate (if required)",
		"Physically Handicapped certificate (if required)",
	}
	for i, d := range documents {
		config.DB.Create(&models.AdmissionDocument{Name: d, SortOrder: i + 1})
	}

	// 5. Seed Fees
	fees := []models.AdmissionFee{
		{Program: "B.Tech (CSE/AI/IT)", TuitionFee: "₹85,000", DevelopmentFee: "₹15,000", Total: "₹1,00,000", SortOrder: 1},
		{Program: "B.Tech (Mechanical/Civil)", TuitionFee: "₹75,000", DevelopmentFee: "₹15,000", Total: "₹90,000", SortOrder: 2},
		{Program: "B.Tech (ECE/EEE)", TuitionFee: "₹80,000", DevelopmentFee: "₹15,000", Total: "₹95,000", SortOrder: 3},
		{Program: "M.Tech", TuitionFee: "₹60,000", DevelopmentFee: "₹10,000", Total: "₹70,000", SortOrder: 4},
		{Program: "MBA", TuitionFee: "₹55,000", DevelopmentFee: "₹10,000", Total: "₹65,000", SortOrder: 5},
	}
	for _, f := range fees {
		config.DB.Create(&f)
	}

	log.Println("Admission data seeded successfully!")
}
