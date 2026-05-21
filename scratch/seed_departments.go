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

	departments := []models.Department{
		{
			Name:          "Computer Science & Engineering",
			ShortName:     "CSE",
			Description:   "The Department of Computer Science & Engineering offers programs with a strong foundation in algorithms, programming, AI, data science, and modern software development.",
			Image:         "https://images.unsplash.com/photo-1517430816045-df4b7de11d1d?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1351&q=80",
			StudentCount:  450,
			CourseCount:   18,
			PlacementRate: "98%",
			FacultyCount:  25,
			LabCount:      8,
			Established:   2008,
			Highlights:    datatypes.JSON(`["AI & ML", "Cybersecurity", "Full Stack Development", "Data Science"]`),
			Vision:        "To be a center of excellence in computing and research, producing globally competent professionals who can innovate and lead in the ever-evolving field of technology.",
			Mission:       datatypes.JSON(`["Provide quality education in computer science with industry-relevant curriculum", "Promote research and innovation in emerging technologies like AI, ML, and IoT", "Prepare students for global IT challenges through industry collaborations", "Foster ethical practices and lifelong learning among students"]`),
			Programs:      datatypes.JSON(`[{"name": "B.Tech in CSE", "duration": "4 Years", "seats": 120}, {"name": "M.Tech in CSE", "duration": "2 Years", "seats": 30}, {"name": "B.Tech in AI & ML", "duration": "4 Years", "seats": 60}, {"name": "Ph.D in Computer Science", "duration": "3-5 Years", "seats": 15}]`),
			Facilities:    datatypes.JSON(`["Advanced Computing Lab", "AI & Robotics Lab", "Data Science Center", "IoT Innovation Lab", "High-Performance Computing Cluster"]`),
			Achievements:   datatypes.JSON(`["100% placement for last 3 years", "₹42 LPA highest package (Microsoft)", "15+ research papers published in 2023", "Winners of National Hackathon 2023"]`),
			Email:         "cse@krishnacollege.edu",
			Phone:         "+91-1234567890",
			Link:          "/departments/cse",
			IsActive:      true,
		},
		{
			Name:          "Civil Engineering",
			ShortName:     "CIVIL",
			Description:   "Civil Engineering focuses on infrastructure development, sustainable design, and construction technologies.",
			Image:         "https://images.unsplash.com/photo-1541888946425-d81bb19240f5?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1350&q=80",
			StudentCount:  280,
			CourseCount:   15,
			PlacementRate: "92%",
			FacultyCount:  32,
			LabCount:      6,
			Established:   2008,
			Highlights:    datatypes.JSON(`["Structural Design", "Environmental Engineering", "Transportation"]`),
			Vision:        "To produce skilled civil engineers for modern infrastructure development with sustainable practices.",
			Mission:       datatypes.JSON(`["Develop sustainable and eco-friendly solutions", "Encourage practical knowledge with industry exposure", "Foster innovation in construction techniques"]`),
			Programs:      datatypes.JSON(`[{"name": "B.Tech in Civil Engineering", "duration": "4 Years", "seats": 90}, {"name": "M.Tech in Structural Engineering", "duration": "2 Years", "seats": 25}]`),
			Facilities:    datatypes.JSON(`["Structural Lab", "Geotechnical Lab", "Surveying Lab"]`),
			Achievements:   datatypes.JSON(`["Consultancy projects for State Govt", "Research grants for eco-friendly bricks"]`),
			Email:         "hod.civil@kec.edu",
			Phone:         "+91 7000130300",
			Link:          "/departments/civil",
			IsActive:      true,
		},
		{
			Name:          "Mechanical Engineering",
			ShortName:     "ME",
			Description:   "Mechanical Engineering covers thermal sciences, robotics, design, and manufacturing technologies.",
			Image:         "https://images.unsplash.com/photo-1581091226033-d5c48150dbaa?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1350&q=80",
			StudentCount:  380,
			CourseCount:   16,
			PlacementRate: "94%",
			FacultyCount:  35,
			LabCount:      6,
			Established:   2008,
			Highlights:    datatypes.JSON(`["Thermodynamics", "Fluid Mechanics", "Materials Science"]`),
			Vision:        "To create leaders in mechanical and industrial design.",
			Mission:       datatypes.JSON(`["Deliver strong fundamentals in design and automation", "Collaborate with industries for real-world projects", "Encourage innovation in renewable energy systems"]`),
			Programs:      datatypes.JSON(`[{"name": "B.Tech in Mechanical Engineering", "duration": "4 Years", "seats": 120}, {"name": "M.Tech in Thermal Engineering", "duration": "2 Years", "seats": 20}]`),
			Facilities:    datatypes.JSON(`["Robotics Lab", "IC Engine Lab", "Workshop Practice"]`),
			Achievements:   datatypes.JSON(`["Student team won Formula Student India", "10+ patents filed by faculty"]`),
			Email:         "hod.mech@kec.edu",
			Phone:         "+91 7000130301",
			Link:          "/departments/mech",
			IsActive:      true,
		},
		{
			Name:          "Electrical Engineering",
			ShortName:     "EE",
			Description:   "Electrical Engineering emphasizes power systems, electronics, renewable energy, and automation.",
			Image:         "https://images.unsplash.com/photo-1504309092620-4d0ec726efa4?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1350&q=80",
			StudentCount:  320,
			CourseCount:   16,
			PlacementRate: "91%",
			FacultyCount:  38,
			LabCount:      7,
			Established:   2008,
			Highlights:    datatypes.JSON(`["Power Systems", "Renewable Energy", "Control Systems", "Electronics"]`),
			Vision:        "To contribute to innovations in energy and automation.",
			Mission:       datatypes.JSON(`["Provide strong knowledge in electronics & power systems", "Promote renewable energy research", "Equip students with industrial automation skills"]`),
			Programs:      datatypes.JSON(`[{"name": "B.Tech in Electrical Engineering", "duration": "4 Years", "seats": 90}, {"name": "M.Tech in Power Systems", "duration": "2 Years", "seats": 18}]`),
			Facilities:    datatypes.JSON(`["Power System Lab", "Electrical Machines Lab", "Control System Lab"]`),
			Achievements:   datatypes.JSON(`["Smart Grid project funded by DST", "Best Student Paper award in IEEE conference"]`),
			Email:         "hod.ee@kec.edu",
			Phone:         "+91 7000130302",
			Link:          "/departments/eee",
			IsActive:      true,
		},
	}

	for _, dept := range departments {
		var existing models.Department
		if err := config.DB.Where("short_name = ?", dept.ShortName).First(&existing).Error; err != nil {
			config.DB.Create(&dept)
			log.Printf("Created Department: %s\n", dept.Name)
		} else {
			config.DB.Model(&existing).Updates(map[string]interface{}{
				"name":           dept.Name,
				"description":    dept.Description,
				"image":          dept.Image,
				"student_count":  dept.StudentCount,
				"course_count":   dept.CourseCount,
				"placement_rate": dept.PlacementRate,
				"faculty_count":  dept.FacultyCount,
				"lab_count":      dept.LabCount,
				"established":    dept.Established,
				"highlights":     dept.Highlights,
				"vision":         dept.Vision,
				"mission":        dept.Mission,
				"programs":       dept.Programs,
				"facilities":     dept.Facilities,
				"achievements":   dept.Achievements,
				"email":          dept.Email,
				"phone":          dept.Phone,
				"link":           dept.Link,
				"is_active":      dept.IsActive,
			})
			log.Printf("Updated Department: %s\n", dept.Name)
		}
	}

	log.Println("Department seeding (EXACT from departments.js) complete!")
}
