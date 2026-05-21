package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Course struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title         string         `gorm:"not null"`
	ShortName     string         `json:"short_name"`
	Icon          string         `json:"icon"`
	Duration      string         `json:"duration"`
	Seats         int            `json:"seats"`
	Credits       int            `json:"credits"`
	Description   string         `json:"description"`
	Highlights    datatypes.JSON `json:"highlights"`
	Career        string         `json:"career"`
	Color         string         `json:"color"`
	IconColor     string         `json:"icon_color"`
	BgColor       string         `json:"bg_color"`
	Department    string         `json:"department"`
	Eligibility   string         `json:"eligibility"`
	Fees          string         `json:"fees"`
	IsActive      bool           `gorm:"default:true"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	courses := []Course{
		{
			Title:       "Artificial Intelligence",
			ShortName:   "AI",
			Icon:        "Brain",
			Duration:    "4 years",
			Seats:       60,
			Credits:     160,
			Description: "Learn to create intelligent systems that can perform tasks requiring human intelligence.",
			Highlights:  jsonToRaw([]string{"Machine Learning algorithms", "Natural Language Processing", "Computer Vision", "Neural Networks", "AI Ethics"}),
			Career:      "AI Engineer, Data Scientist, Research Scientist",
			Color:       "from-purple-500 to-pink-500",
			IconColor:   "text-purple-500",
			BgColor:     "bg-purple-50",
			Department:  "Computer Science",
			Eligibility: "10+2 with Physics, Chemistry, and Mathematics",
			Fees:        "₹95,000 per year",
			IsActive:    true,
		},
		{
			Title:       "Machine Learning",
			ShortName:   "ML",
			Icon:        "Cpu",
			Duration:    "4 years",
			Seats:       60,
			Credits:     160,
			Description: "Master algorithms and statistical models that computer systems use to perform tasks without explicit instructions.",
			Highlights:  jsonToRaw([]string{"Supervised & Unsupervised Learning", "Deep Learning", "Reinforcement Learning", "Predictive Analytics", "ML Operations"}),
			Career:      "ML Engineer, AI Specialist, Data Analyst",
			Color:       "from-blue-500 to-cyan-500",
			IconColor:   "text-blue-500",
			BgColor:     "bg-blue-50",
			Department:  "Computer Science",
			Eligibility: "10+2 with Physics, Chemistry, and Mathematics",
			Fees:        "₹92,000 per year",
			IsActive:    true,
		},
		{
			Title:       "Cyber Security",
			ShortName:   "CS",
			Icon:        "Shield",
			Duration:    "4 years",
			Seats:       60,
			Credits:     160,
			Description: "Protect systems, networks, and programs from digital attacks and unauthorized access.",
			Highlights:  jsonToRaw([]string{"Network Security", "Cryptography", "Ethical Hacking", "Digital Forensics", "Security Operations"}),
			Career:      "Security Analyst, Ethical Hacker, Security Architect",
			Color:       "from-green-500 to-emerald-500",
			IconColor:   "text-green-500",
			BgColor:     "bg-green-50",
			Department:  "Computer Science",
			Eligibility: "10+2 with Physics, Chemistry, and Mathematics",
			Fees:        "₹90,000 per year",
			IsActive:    true,
		},
		{
			Title:       "Cloud Computing",
			ShortName:   "CC",
			Icon:        "Cloud",
			Duration:    "4 years",
			Seats:       60,
			Credits:     160,
			Description: "Design, implement and manage cloud-based systems and infrastructure.",
			Highlights:  jsonToRaw([]string{"Cloud Architecture", "Virtualization Technologies", "AWS/Azure/GCP", "DevOps", "Containerization"}),
			Career:      "Cloud Architect, Cloud Engineer, DevOps Engineer",
			Color:       "from-orange-500 to-amber-500",
			IconColor:   "text-orange-500",
			BgColor:     "bg-orange-50",
			Department:  "Computer Science",
			Eligibility: "10+2 with Physics, Chemistry, and Mathematics",
			Fees:        "₹88,000 per year",
			IsActive:    true,
		},
		{
			Title:       "Data Science",
			ShortName:   "DS",
			Icon:        "Database",
			Duration:    "4 years",
			Seats:       60,
			Credits:     160,
			Description: "Extract insights and knowledge from structured and unstructured data using scientific methods.",
			Highlights:  jsonToRaw([]string{"Statistical Analysis", "Data Visualization", "Big Data Technologies", "Predictive Modeling", "Business Intelligence"}),
			Career:      "Data Scientist, Data Analyst, Business Analyst",
			Color:       "from-red-500 to-pink-500",
			IconColor:   "text-red-500",
			BgColor:     "bg-red-50",
			Department:  "Computer Science",
			Eligibility: "10+2 with Physics, Chemistry, and Mathematics",
			Fees:        "₹93,000 per year",
			IsActive:    true,
		},
		{
			Title:       "Electric Vehicle Technology",
			ShortName:   "EV",
			Icon:        "Zap",
			Duration:    "4 years",
			Seats:       60,
			Credits:     160,
			Description: "Design and develop electric vehicles and their components with focus on sustainability.",
			Highlights:  jsonToRaw([]string{"Electric Powertrains", "Battery Technology", "Vehicle Dynamics", "Charging Infrastructure", "Sustainable Mobility"}),
			Career:      "EV Design Engineer, Battery Engineer, Automotive Engineer",
			Color:       "from-teal-500 to-cyan-500",
			IconColor:   "text-teal-500",
			BgColor:     "bg-teal-50",
			Department:  "Mechanical Engineering",
			Eligibility: "10+2 with Physics, Chemistry, and Mathematics",
			Fees:        "₹96,000 per year",
			IsActive:    true,
		},
		{
			Title:       "Fast Charging Station Technology",
			ShortName:   "FCST",
			Icon:        "BatteryCharging",
			Duration:    "4 years",
			Seats:       60,
			Credits:     160,
			Description: "Specialize in the design and implementation of rapid charging infrastructure for electric vehicles.",
			Highlights:  jsonToRaw([]string{"Power Electronics", "Grid Integration", "Energy Storage Systems", "Smart Charging", "Renewable Integration"}),
			Career:      "Charging Infrastructure Engineer, Power Systems Engineer",
			Color:       "from-yellow-500 to-amber-500",
			IconColor:   "text-yellow-500",
			BgColor:     "bg-yellow-50",
			Department:  "Electrical Engineering",
			Eligibility: "10+2 with Physics, Chemistry, and Mathematics",
			Fees:        "₹94,000 per year",
			IsActive:    true,
		},
		{
			Title:       "Robotics & Automation",
			ShortName:   "RA",
			Icon:        "Bot",
			Duration:    "4 years",
			Seats:       60,
			Credits:     160,
			Description: "Design, build, and program robotic systems for industrial and consumer applications.",
			Highlights:  jsonToRaw([]string{"Robot Mechanics", "Control Systems", "Computer Vision", "Industrial Automation", "AI in Robotics"}),
			Career:      "Robotics Engineer, Automation Specialist, Control Systems Engineer",
			Color:       "from-indigo-500 to-blue-500",
			IconColor:   "text-indigo-500",
			BgColor:     "bg-indigo-50",
			Department:  "Mechanical Engineering",
			Eligibility: "10+2 with Physics, Chemistry, and Mathematics",
			Fees:        "₹97,000 per year",
			IsActive:    true,
		},
	}

	for _, course := range courses {
		if err := db.Create(&course).Error; err != nil {
			log.Printf("Error creating course %s: %v", course.Title, err)
		} else {
			fmt.Printf("Created course: %s\n", course.Title)
		}
	}

	fmt.Println("Seeding completed!")
}

func jsonToRaw(v interface{}) datatypes.JSON {
	b, _ := json.Marshal(v)
	return datatypes.JSON(b)
}
