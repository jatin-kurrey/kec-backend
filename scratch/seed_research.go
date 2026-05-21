package main

import (
	"log"
	"kec-backend/internal/config"
	"kec-backend/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/datatypes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	config.ConnectDB()
	
	// Clear existing data to avoid duplicates and ensure "exact" state
	config.DB.Exec("DELETE FROM research_stats")
	config.DB.Exec("DELETE FROM research_areas")
	config.DB.Exec("DELETE FROM research_projects")
	config.DB.Exec("DELETE FROM research_facilities")

	// 1. Seed Research Stats
	stats := []models.ResearchStat{
		{Label: "Ongoing Projects", Value: "56+", Icon: "FlaskRound", Color: "from-blue-500 to-blue-600", SortOrder: 1},
		{Label: "Research Funding", Value: "₹1Cr+", Icon: "BookOpen", Color: "from-purple-500 to-purple-600", SortOrder: 2},
		{Label: "Publications", Value: "225+", Icon: "GraduationCap", Color: "from-green-500 to-green-600", SortOrder: 3},
		{Label: "Patents Filed", Value: "18", Icon: "Award", Color: "from-orange-500 to-orange-600", SortOrder: 4},
	}

	for _, s := range stats {
		config.DB.Create(&s)
	}

	// 2. Seed Research Areas
	areas := []models.ResearchArea{
		{
			Title: "Artificial Intelligence & Machine Learning",
			Description: "Advanced research in neural networks, deep learning, and AI applications",
			Icon: "Cpu",
			Projects: 12,
			Publications: 45,
			Color: "from-blue-500 to-blue-700",
		},
		{
			Title: "Renewable Energy Systems",
			Description: "Developing sustainable energy solutions and storage technologies",
			Icon: "Zap",
			Projects: 8,
			Publications: 32,
			Color: "from-green-500 to-green-700",
		},
		{
			Title: "Biotechnology & Genetic Engineering",
			Description: "Cutting-edge research in genomics, proteomics and medical biotechnology",
			Icon: "Dna",
			Projects: 10,
			Publications: 38,
			Color: "from-purple-500 to-purple-700",
		},
		{
			Title: "Nanotechnology & Materials Science",
			Description: "Design and synthesis of novel nanomaterials for various applications",
			Icon: "Atom",
			Projects: 7,
			Publications: 29,
			Color: "from-orange-500 to-orange-700",
		},
		{
			Title: "Data Science & Big Data Analytics",
			Description: "Research on data mining, pattern recognition, and predictive analytics",
			Icon: "Database",
			Projects: 9,
			Publications: 41,
			Color: "from-indigo-500 to-indigo-700",
		},
		{
			Title: "Internet of Things & Cyber-Physical Systems",
			Description: "Developing smart connected systems for industrial and domestic use",
			Icon: "Network",
			Projects: 11,
			Publications: 36,
			Color: "from-cyan-500 to-cyan-700",
		},
	}

	for _, a := range areas {
		config.DB.Create(&a)
	}

	// 3. Seed Research Projects
	projects := []models.ResearchProject{
		{
			Title: "Electric Vehicle Development with Sodium-Ion Battery",
			Department: "Mechanical Engineering",
			Funding: "₹55,00,000",
			Duration: "2024-2026",
			Status: "Ongoing",
			Description: "India's first initiative in EV development using sodium-ion batteries to enhance vehicle range in a single charge. Focus on energy efficiency and cost reduction.",
			Team: datatypes.JSON([]byte(`["Dr. Joy Sonashalol", "Mr. Ash Kumar Soni", "6 Research Scholars"]`)),
			Outcomes: "Prototype Developed, 2 Publications",
			Image: "https://images.unsplash.com/photo-1614436165834-50e93f5cd79e?ixlib=rb-4.0.3&auto=format&fit=crop&w=800&q=80",
		},
		{
			Title: "Fast Charging Station Infrastructure with GPS Integration",
			Department: "Electrical Engineering",
			Funding: "₹48,00,000",
			Duration: "2023-2025",
			Status: "Ongoing",
			Description: "Development of fast charging station modules for EVs with GPS-enabled tracking for easy location access by users and service providers.",
			Team: datatypes.JSON([]byte(`["Mr. Tarachand Sahu", "4 Research Scholars"]`)),
			Outcomes: "Mobile App Developed, Pilot Deployment",
			Image: "https://images.unsplash.com/photo-1581091012184-5c46a1cd2765?ixlib=rb-4.0.3&auto=format&fit=crop&w=800&q=80",
		},
		{
			Title: "Intelligent Robotics with Artificial Intelligence",
			Department: "Computer Science & Engineering",
			Funding: "₹62,00,000",
			Duration: "2023-2026",
			Status: "Ongoing",
			Description: "Building advanced robotics systems capable of adaptive behavior and decision-making through machine learning and AI integration.",
			Team: datatypes.JSON([]byte(`["Dr. Priya Singh", "5 Research Scholars"]`)),
			Outcomes: "3 Publications, 1 Patent Filed",
			Image: "https://images.unsplash.com/photo-1581090700227-1e8f9e1f2c44?ixlib=rb-4.0.3&auto=format&fit=crop&w=800&q=80",
		},
		{
			Title: "Semiconductor Chip Research & Manufacturing",
			Department: "Electronics and Communication",
			Funding: "₹78,00,000",
			Duration: "2024-2027",
			Status: "Ongoing",
			Description: "Research on design, material optimization, and indigenous manufacturing of semiconductor chips aimed at reducing import dependency.",
			Team: datatypes.JSON([]byte(`["Dr. V. Sharma", "Dr. S. Joshi", "6 Research Scholars"]`)),
			Outcomes: "2 Prototypes, Industry Collaboration Initiated",
			Image: "https://images.unsplash.com/photo-1581092334600-5fc179ee4c35?ixlib=rb-4.0.3&auto=format&fit=crop&w=800&q=80",
		},
		{
			Title: "Hyperloop Transportation System Research",
			Department: "Mechanical Engineering",
			Funding: "₹1,00,00,000",
			Duration: "2024-2028",
			Status: "Ongoing",
			Description: "Exploratory research on implementing Hyperloop technology for ultra-fast transportation, focusing on magnetic levitation and vacuum tube dynamics.",
			Team: datatypes.JSON([]byte(`["Dr. Joy Sonashalol", "8 Research Scholars"]`)),
			Outcomes: "Design Simulation Completed, 1 Journal Publication",
			Image: "https://images.unsplash.com/photo-1504384764586-bb4cdc1707b0?ixlib=rb-4.0.3&auto=format&fit=crop&w=800&q=80",
		},
	}

	for _, p := range projects {
		config.DB.Create(&p)
	}

	// 4. Seed Research Facilities
	facilities := []models.ResearchFacility{
		{
			Name: "Advanced Computing Lab",
			Description: "Equipped with high-performance computing clusters for complex simulations and AI research",
			Image: "https://images.unsplash.com/photo-1517430816045-df4b7de11d1d?ixlib=rb-4.0.3&auto=format&fit=crop&w=800&q=80",
			Features: datatypes.JSON([]byte(`["HPC Clusters", "GPU Workstations", "Cloud Computing", "AI Training Infrastructure"]`)),
			Capacity: "40 researchers",
		},
		{
			Name: "Nano Research Center",
			Description: "State-of-the-art facilities for nanomaterials synthesis and characterization",
			Image: "https://images.unsplash.com/photo-1581092580497-e0d23cbdf1dc?ixlib=rb-4.0.3&auto=format&fit=crop&w=800&q=80",
			Features: datatypes.JSON([]byte(`["SEM/TEM Microscopy", "Clean Room", "Thin Film Deposition", "Material Testing"]`)),
			Capacity: "25 researchers",
		},
		{
			Name: "Biotechnology Research Center",
			Description: "Modern labs for genetic engineering, cell culture, and bioprocessing research",
			Image: "https://images.unsplash.com/photo-1581092918056-0c4c3acd3789?ixlib=rb-4.0.3&auto=format&fit=crop&w=800&q=80",
			Features: datatypes.JSON([]byte(`["PCR Labs", "Cell Culture", "Fermentation", "Chromatography"]`)),
			Capacity: "30 researchers",
		},
		{
			Name: "Renewable Energy Lab",
			Description: "Testing facilities for solar, wind and other renewable energy systems",
			Image: "https://images.unsplash.com/photo-1466611653911-95081537e5b7?ixlib=rb-4.0.3&auto=format&fit=crop&w=800&q=80",
			Features: datatypes.JSON([]byte(`["Solar Simulators", "Wind Tunnels", "Battery Testing", "Smart Grid Systems"]`)),
			Capacity: "20 researchers",
		},
		{
			Name: "Data Science Center",
			Description: "Advanced tools for big data analytics, visualization, and machine learning research",
			Image: "https://images.unsplash.com/photo-1551288049-bebda4e38f71?ixlib=rb-4.0.3&auto=format&fit=crop&w=800&q=80",
			Features: datatypes.JSON([]byte(`["Data Visualization", "ML Workbenches", "Database Servers", "Collaboration Spaces"]`)),
			Capacity: "35 researchers",
		},
		{
			Name: "Electronics Prototyping Lab",
			Description: "Facilities for IoT development, embedded systems, and electronics prototyping",
			Image: "https://images.unsplash.com/photo-1558442074-3c19857bc1dc?ixlib=rb-4.0.3&auto=format&fit=crop&w=800&q=80",
			Features: datatypes.JSON([]byte(`["3D Printing", "PCB Fabrication", "IoT Testbeds", "Sensor Networks"]`)),
			Capacity: "28 researchers",
		},
	}

	for _, f := range facilities {
		config.DB.Create(&f)
	}

	log.Println("Research data seeded successfully!")
}
