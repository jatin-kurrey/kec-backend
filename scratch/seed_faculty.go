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

	faculty := []models.Faculty{
		{SNo: 1, Name: "Dr. SUDESHNA SENGUPTA", Role: "ASST PROFESSOR", Department: "HUMANITIES", Qualification: "Ph.D., M.A.", Experience: "12 years", Specialization: "English Literature", Publications: 5, Projects: 2, Bio: "Dedicated professor with a passion for humanities and literature."},
		{SNo: 2, Name: "Dr. CHAND RAM", Role: "ASST PROFESSOR", Department: "APPLIED MATHEMATICS", Qualification: "Ph.D., M.Sc.", Experience: "15 years", Specialization: "Advanced Calculus", Publications: 8, Projects: 3, Bio: "Expert in mathematical modeling and applied sciences."},
		{SNo: 3, Name: "Mrs. RICHA SHARMA", Role: "ASST PROFESSOR", Department: "APPLIED CHEMISTRY", Qualification: "M.Sc., B.Ed.", Experience: "10 years", Specialization: "Organic Chemistry", Publications: 3, Projects: 1, Bio: "Focused on chemical research and industrial applications."},
		{SNo: 4, Name: "Mr. AMIT PANDEY", Role: "ASST PROFESSOR", Department: "ELECTRICAL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "8 years", Specialization: "Power Systems", Publications: 4, Projects: 2, Bio: "Passionate about renewable energy and smart grids."},
		{SNo: 5, Name: "Mr. DEVENDRA DEWANGAN", Role: "ASST PROFESSOR", Department: "MECHANICAL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "9 years", Specialization: "Thermal Engineering", Publications: 2, Projects: 2, Bio: "Expert in thermodynamics and heat transfer."},
		{SNo: 6, Name: "Mr. MUKESH TAMBOLI", Role: "ASST PROFESSOR", Department: "APPLIED PHYSICS", Qualification: "M.Sc., M.Phil.", Experience: "11 years", Specialization: "Solid State Physics", Publications: 6, Projects: 1},
		{SNo: 7, Name: "Mrs. NAMITA MISHRA", Role: "ASST PROFESSOR", Department: "APPLIED PHYSICS", Qualification: "M.Sc., NET", Experience: "7 years", Specialization: "Optics", Publications: 2, Projects: 1},
		{SNo: 8, Name: "Mr. NEERAJ CHANDRAKAR", Role: "ASST PROFESSOR", Department: "MECHANICAL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "10 years", Specialization: "Manufacturing Technology", Publications: 3, Projects: 2},
		{SNo: 9, Name: "Mr. RAMESHWAR MISHRA", Role: "ASST PROFESSOR", Department: "MECHANICAL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "12 years", Specialization: "Machine Design", Publications: 4, Projects: 3},
		{SNo: 10, Name: "Ms. VANDANA YADAV", Role: "ASST PROFESSOR", Department: "COMPUTER SCIENCE & ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "6 years", Specialization: "Data Structures", Publications: 2, Projects: 2},
		{SNo: 11, Name: "Mr. SAKET CHANDRAKAR", Role: "ASST PROFESSOR", Department: "CIVIL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "7 years", Specialization: "Structural Engineering", Publications: 3, Projects: 1},
		{SNo: 12, Name: "Ms SHIVAN ANISHA", Role: "LECTURER", Department: "CIVIL ENGINEERING", Qualification: "B.E.", Experience: "3 years", Specialization: "Surveying", Publications: 0, Projects: 1},
		{SNo: 13, Name: "Mrs. VEENA VERMA", Role: "ASST PROFESSOR", Department: "ELECTRICAL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "9 years", Specialization: "Control Systems", Publications: 2, Projects: 1},
		{SNo: 14, Name: "Mr. PIYUSH KUMAR", Role: "ASST PROFESSOR", Department: "ELECTRICAL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "8 years", Specialization: "Electrical Machines", Publications: 3, Projects: 2},
		{SNo: 15, Name: "Ms. MAMTA DEWANGAN", Role: "ASST PROFESSOR", Department: "COMPUTER SCIENCE & ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "5 years", Specialization: "Web Technologies", Publications: 1, Projects: 3},
		{SNo: 16, Name: "Mr. JOY SONA", Role: "ASST PROFESSOR", Department: "COMPUTER SCIENCE & ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "10 years", Specialization: "Network Security", Publications: 5, Projects: 2},
		{SNo: 17, Name: "Mr. PRABHAT PRASAD", Role: "ASST PROFESSOR", Department: "CIVIL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "6 years", Specialization: "Transportation Engineering", Publications: 2, Projects: 1},
		{SNo: 18, Name: "Mr. HOMENDRA KUMAR", Role: "ASST PROFESSOR", Department: "ELECTRICAL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "7 years", Specialization: "Digital Electronics", Publications: 2, Projects: 1},
		{SNo: 19, Name: "Mr. RAVINDRA SHARMA", Role: "ASSOCIATE PROFESSOR", Department: "CIVIL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "18 years", Specialization: "Geotechnical Engineering", Publications: 10, Projects: 5},
		{SNo: 20, Name: "Mr. ASIF NIZAM", Role: "ASST PROFESSOR", Department: "MECHANICAL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "11 years", Specialization: "Automobile Engineering", Publications: 4, Projects: 2},
		{SNo: 21, Name: "Ms. AANCHAL LAL", Role: "ASST PROFESSOR", Department: "ELECTRICAL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "6 years", Specialization: "Power Electronics", Publications: 1, Projects: 1},
		{SNo: 22, Name: "Mr. ABINASH SWAIN", Role: "ASST PROFESSOR", Department: "MECHANICAL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "7 years", Specialization: "CAD/CAM", Publications: 2, Projects: 1},
		{SNo: 23, Name: "Ms. PRIYANKA SHARMA", Role: "ASST PROFESSOR", Department: "ELECTRICAL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "5 years", Specialization: "Signals and Systems", Publications: 1, Projects: 1},
		{SNo: 24, Name: "Mr. JAGENDRA NARANG", Role: "ASST PROFESSOR", Department: "COMPUTER SCIENCE & ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "8 years", Specialization: "Software Engineering", Publications: 3, Projects: 2},
		{SNo: 25, Name: "Ms. JUHI CHATTORAJ", Role: "ASST PROFESSOR", Department: "ELECTRICAL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "4 years", Specialization: "Instrumentation", Publications: 0, Projects: 1},
		{SNo: 26, Name: "Dr. NIRBHAY SINGH", Role: "ASST PROFESSOR", Department: "PHYSICS", Qualification: "Ph.D., M.Sc.", Experience: "13 years", Specialization: "Quantum Mechanics", Publications: 7, Projects: 2},
		{SNo: 27, Name: "Mr. NITESH KUMAR", Role: "ASST PROFESSOR", Department: "COMPUTER SCIENCE & ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "9 years", Specialization: "Operating Systems", Publications: 3, Projects: 2},
		{SNo: 28, Name: "Mr. AJAY SAHU", Role: "ASST PROFESSOR", Department: "MECHANICAL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "10 years", Specialization: "Fluid Mechanics", Publications: 2, Projects: 1},
		{SNo: 29, Name: "Mr. JITENDRA VERMA", Role: "ASST PROFESSOR", Department: "ELECTRICAL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "11 years", Specialization: "Control Systems", Publications: 4, Projects: 2},
		{SNo: 30, Name: "Ms. MANPREET KAUR", Role: "ASST PROFESSOR", Department: "COMPUTER SCIENCE & ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "6 years", Specialization: "Algorithms", Publications: 2, Projects: 3},
		{SNo: 31, Name: "Mrs. PRATIBHA DONGRE", Role: "ASST PROFESSOR", Department: "COMPUTER SCIENCE & ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "7 years", Specialization: "Database Management", Publications: 1, Projects: 2},
		{SNo: 32, Name: "Mr. TARACHAND SAHU", Role: "ASST PROFESSOR", Department: "ELECTRICAL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "10 years", Specialization: "Power Systems", Publications: 2, Projects: 1},
		{SNo: 33, Name: "Mr. ABDUL ANSARI", Role: "ASST PROFESSOR", Department: "CIVIL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "8 years", Specialization: "Environmental Engineering", Publications: 3, Projects: 2},
		{SNo: 34, Name: "Mr. PRATIK GOYAL", Role: "ASST PROFESSOR", Department: "CIVIL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "6 years", Specialization: "Concrete Technology", Publications: 1, Projects: 1},
		{SNo: 35, Name: "Mr. SYED ULLAH", Role: "ASST PROFESSOR", Department: "MECHANICAL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "9 years", Specialization: "Thermodynamics", Publications: 2, Projects: 1},
		{SNo: 36, Name: "Dr. AJAY TIWARI", Role: "PRINCIPAL/DIRECTOR", Department: "MECHANICAL ENGINEERING", Qualification: "Ph.D., M.Tech", Experience: "33 years", Specialization: "Robotics", Publications: 21, Projects: 5, Bio: "Visionary leader and robotics expert."},
		{SNo: 37, Name: "Mr. AMIT DAS", Role: "ASST PROFESSOR", Department: "CIVIL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "7 years", Specialization: "Hydrology", Publications: 2, Projects: 1},
		{SNo: 38, Name: "Mr. DEVENDRA SAHU", Role: "ASST PROFESSOR", Department: "CIVIL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "8 years", Specialization: "Urban Planning", Publications: 1, Projects: 2},
		{SNo: 39, Name: "Ms. KRITIKA SAHU", Role: "ASST PROFESSOR", Department: "ELECTRICAL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "5 years", Specialization: "Electrical Measurements", Publications: 1, Projects: 1},
		{SNo: 40, Name: "Ms. NEELAM AGRAWAL", Role: "ASST PROFESSOR", Department: "COMPUTER SCIENCE & ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "6 years", Specialization: "Machine Learning", Publications: 2, Projects: 2},
		{SNo: 41, Name: "Ms. PRIYANKA THAKUR", Role: "ASST PROFESSOR", Department: "COMPUTER SCIENCE & ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "7 years", Specialization: "Artificial Intelligence", Publications: 3, Projects: 3},
		{SNo: 42, Name: "Mr. VINAY SHRIVASTAVA", Role: "ASST PROFESSOR", Department: "COMPUTER SCIENCE AND ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "12 years", Specialization: "Cloud Computing", Publications: 6, Projects: 4},
		{SNo: 43, Name: "Ms. MINEETA KHANUJA", Role: "ASST PROFESSOR", Department: "COMPUTER SCIENCE & ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "6 years", Specialization: "Python Programming", Publications: 1, Projects: 2},
		{SNo: 44, Name: "Mrs. GARIMA SINHA", Role: "ASST PROFESSOR", Department: "COMPUTER SCIENCE & ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "8 years", Specialization: "Software Testing", Publications: 2, Projects: 1},
		{SNo: 45, Name: "Ms. SEEMA SAHU", Role: "ASST PROFESSOR", Department: "COMPUTER SCIENCE & ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "5 years", Specialization: "Object Oriented Programming", Publications: 1, Projects: 1},
		{SNo: 46, Name: "Mr. SHAILENDRA KHARE", Role: "ASST PROFESSOR", Department: "APPLIED MATHEMATICS", Qualification: "M.Sc., M.Phil.", Experience: "14 years", Specialization: "Numerical Methods", Publications: 5, Projects: 2},
		{SNo: 47, Name: "Mrs. SEEPIE JOHN", Role: "ASST PROFESSOR", Department: "CIVIL ENGINEERING", Qualification: "M.Tech, B.E.", Experience: "10 years", Specialization: "Bridge Engineering", Publications: 3, Projects: 1},
		{SNo: 48, Name: "Mr. NEERAJ RAJPUT", Role: "ASST PROFESSOR", Department: "APPLIED CHEMISTRY", Qualification: "M.Sc.", Experience: "9 years", Specialization: "Physical Chemistry", Publications: 2, Projects: 1},
		{SNo: 49, Name: "Mrs. PARINA VARMA", Role: "ASST PROFESSOR", Department: "PHYSICS", Qualification: "M.Sc.", Experience: "11 years", Specialization: "Electronics", Publications: 4, Projects: 2},
		{SNo: 50, Name: "Dr. ONKAR DIXIT", Role: "ASST PROFESSOR", Department: "APPLIED MATHEMATICS", Qualification: "Ph.D., M.Sc.", Experience: "16 years", Specialization: "Statistical Analysis", Publications: 9, Projects: 4},
	}

	for _, member := range faculty {
		if member.Image == "" {
			member.Image = "https://ui-avatars.com/api/?background=random&color=fff&name=Faculty"
		}
		var existing models.Faculty
		if err := config.DB.Where("name = ?", member.Name).First(&existing).Error; err != nil {
			config.DB.Create(&member)
			log.Printf("Created Faculty: %s\n", member.Name)
		} else {
			config.DB.Model(&existing).Updates(map[string]interface{}{
				"sno":            member.SNo,
				"role":           member.Role,
				"dept":           member.Department,
				"qualification":  member.Qualification,
				"experience":     member.Experience,
				"specialization": member.Specialization,
				"publications":   member.Publications,
				"projects":       member.Projects,
				"bio":            member.Bio,
				"is_active":      true,
			})
			log.Printf("Updated Faculty: %s\n", member.Name)
		}
	}

	log.Println("Faculty seeding (EXACT with profiles) complete!")
}
