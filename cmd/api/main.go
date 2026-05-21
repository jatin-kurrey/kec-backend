package main

import (
	"log"
	"os"

	"kec-backend/internal/config"
	"kec-backend/internal/handlers"
	"kec-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		err = godotenv.Load("../../.env")
	}
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Connect to DB
	config.ConnectDB()

	// Initialize Fiber
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Adjust for production
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("KEC Backend is running!")
	})

	// Public Routes
	api := app.Group("/api/v1")
	
	// Auth
	api.Post("/auth/login", handlers.Login)
	api.Post("/auth/init", handlers.CreateInitialAdmin)

	// Applications (Public)
	api.Post("/applications", handlers.CreateApplication)

	// Content (Public)
	api.Get("/notices", handlers.GetNotices)
	api.Get("/departments", handlers.GetDepartments)
	api.Get("/departments/:shortName", handlers.GetDepartmentByShortName)
	api.Get("/courses", handlers.GetCourses)
	api.Get("/gallery", handlers.GetGallery)
	api.Post("/gallery/:id/like", handlers.IncrementGalleryLikes)
	api.Post("/gallery/:id/download", handlers.IncrementGalleryDownloads)
	api.Get("/leadership", handlers.GetLeadership)
	api.Get("/faculty", handlers.GetFaculty)
	api.Get("/research/areas", handlers.GetResearchAreas)
	api.Get("/research/projects", handlers.GetResearchProjects)
	api.Get("/research/facilities", handlers.GetResearchFacilities)
	api.Get("/research/stats", handlers.GetResearchStats)
	api.Get("/facilities", handlers.GetCampusFacilities)
	api.Get("/facilities/stats", handlers.GetCampusStats)
	
	// Admission (Public)
	api.Get("/admission/guide", handlers.GetAdmissionGuide)
	api.Get("/admission/steps", handlers.GetAdmissionSteps)
	api.Get("/admission/eligibility", handlers.GetAdmissionEligibility)
	api.Get("/admission/documents", handlers.GetAdmissionDocuments)
	api.Get("/admission/fees", handlers.GetAdmissionFees)
	api.Get("/alumni", handlers.GetAlumni)
	api.Get("/alumni/stats", handlers.GetAlumniStats)
	api.Get("/placements/stats", handlers.GetPlacementStats)
	api.Get("/placements/recruiters", handlers.GetRecruiters)
	api.Get("/placements/testimonials", handlers.GetPlacementTestimonials)
	api.Get("/press-media", handlers.GetPressMedia)
	
	// Protected Admin Routes
	admin := api.Group("/admin", middleware.AuthRequired)
	
	// Application Management
	admin.Get("/applications", handlers.GetApplications)
	admin.Put("/applications/:id/status", handlers.UpdateApplicationStatus)

	// Admin Account Management
	admin.Get("/accounts", handlers.GetAdmins)
	admin.Post("/accounts", handlers.CreateAdmin)
	admin.Delete("/accounts/:id", handlers.DeleteAdmin)
	
	// Content Management
	admin.Post("/notices", handlers.CreateNotice)
	admin.Post("/gallery", handlers.AddToGallery)
	admin.Put("/gallery/:id", handlers.UpdateGalleryItem)
	admin.Delete("/gallery/:id", handlers.DeleteGalleryItem)

	// Leadership Management
	admin.Post("/leadership", handlers.CreateLeadership)
	admin.Put("/leadership/:id", handlers.UpdateLeadership)
	admin.Delete("/leadership/:id", handlers.DeleteLeadership)
	admin.Get("/leadership", handlers.GetAllLeadershipAdmin)
	
	// Faculty Management
	admin.Get("/faculty", handlers.GetAllFacultyAdmin)
	admin.Post("/faculty", handlers.CreateFaculty)
	admin.Put("/faculty/:id", handlers.UpdateFaculty)
	admin.Delete("/faculty/:id", handlers.DeleteFaculty)

	// Department Management
	admin.Get("/departments", handlers.GetAllDepartmentsAdmin)
	admin.Post("/departments", handlers.CreateDepartment)
	admin.Put("/departments/:id", handlers.UpdateDepartment)
	admin.Delete("/departments/:id", handlers.DeleteDepartment)

	// Course Management
	admin.Get("/courses", handlers.GetAllCoursesAdmin)
	admin.Post("/courses", handlers.CreateCourse)
	admin.Put("/courses/:id", handlers.UpdateCourse)
	admin.Delete("/courses/:id", handlers.DeleteCourse)

	// Research Management
	admin.Get("/research/areas", handlers.GetAllResearchAreasAdmin)
	admin.Post("/research/areas", handlers.CreateResearchArea)
	admin.Put("/research/areas/:id", handlers.UpdateResearchArea)
	admin.Delete("/research/areas/:id", handlers.DeleteResearchArea)

	admin.Get("/research/projects", handlers.GetAllResearchProjectsAdmin)
	admin.Post("/research/projects", handlers.CreateResearchProject)
	admin.Put("/research/projects/:id", handlers.UpdateResearchProject)
	admin.Delete("/research/projects/:id", handlers.DeleteResearchProject)

	admin.Get("/research/facilities", handlers.GetAllResearchFacilitiesAdmin)
	admin.Post("/research/facilities", handlers.CreateResearchFacility)
	admin.Put("/research/facilities/:id", handlers.UpdateResearchFacility)
	admin.Delete("/research/facilities/:id", handlers.DeleteResearchFacility)

	admin.Get("/research/stats", handlers.GetResearchStats)
	admin.Post("/research/stats", handlers.CreateResearchStat)

	// Campus Facilities Management
	admin.Post("/facilities", handlers.CreateCampusFacility)
	admin.Put("/facilities/:id", handlers.UpdateCampusFacility)
	admin.Delete("/facilities/:id", handlers.DeleteCampusFacility)
	admin.Post("/facilities/stats", handlers.CreateCampusStat)
	admin.Put("/facilities/stats/:id", handlers.UpdateCampusStat)
	admin.Delete("/facilities/stats/:id", handlers.DeleteCampusStat)
	admin.Put("/research/stats/:id", handlers.UpdateResearchStat)
	admin.Delete("/research/stats/:id", handlers.DeleteResearchStat)

	// Admission Management
	admin.Post("/admission/guide", handlers.UpdateAdmissionGuide)
	admin.Post("/admission/steps", handlers.CreateAdmissionStep)
	admin.Put("/admission/steps/:id", handlers.UpdateAdmissionStep)
	admin.Delete("/admission/steps/:id", handlers.DeleteAdmissionStep)
	admin.Post("/admission/eligibility", handlers.CreateAdmissionEligibility)
	admin.Put("/admission/eligibility/:id", handlers.UpdateAdmissionEligibility)
	admin.Delete("/admission/eligibility/:id", handlers.DeleteAdmissionEligibility)
	admin.Post("/admission/documents", handlers.CreateAdmissionDocument)
	admin.Put("/admission/documents/:id", handlers.UpdateAdmissionDocument)
	admin.Delete("/admission/documents/:id", handlers.DeleteAdmissionDocument)
	admin.Post("/admission/fees", handlers.CreateAdmissionFee)
	admin.Put("/admission/fees/:id", handlers.UpdateAdmissionFee)
	admin.Delete("/admission/fees/:id", handlers.DeleteAdmissionFee)

	// Alumni Management
	admin.Post("/alumni", handlers.CreateAlumni)
	admin.Put("/alumni/:id", handlers.UpdateAlumni)
	admin.Delete("/alumni/:id", handlers.DeleteAlumni)
	admin.Post("/alumni/stats", handlers.UpdateAlumniStats)

	// Placement Management
	admin.Post("/placements/stats", handlers.UpdatePlacementStats)
	admin.Post("/placements/recruiters", handlers.CreateRecruiter)
	admin.Put("/placements/recruiters/:id", handlers.UpdateRecruiter)
	admin.Delete("/placements/recruiters/:id", handlers.DeleteRecruiter)
	admin.Post("/placements/testimonials", handlers.CreatePlacementTestimonial)
	admin.Put("/placements/testimonials/:id", handlers.UpdatePlacementTestimonial)
	admin.Delete("/placements/testimonials/:id", handlers.DeletePlacementTestimonial)

	// Press Media Management
	admin.Post("/press-media", handlers.CreatePressMedia)
	admin.Put("/press-media/:id", handlers.UpdatePressMedia)
	admin.Delete("/press-media/:id", handlers.DeletePressMedia)

	admin.Post("/upload", handlers.UploadFile)

	// Static Files
	app.Static("/uploads", "./uploads")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
