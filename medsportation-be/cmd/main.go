package main

import (
	"log"
	"time"

	"medsportation-be/internal/config"
	"medsportation-be/internal/handlers"
	"medsportation-be/internal/middleware"
	"medsportation-be/internal/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	// 2. Initialize Database
	db, err := gorm.Open(sqlite.Open(cfg.DatabasePath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&models.Quote{}, &models.User{}, &models.Consultation{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Create default admin if none exists
	if err := handlers.CreateInitialAdmin(db); err != nil {
		log.Printf("Warning: could not create initial admin: %v", err)
	}

	// 3. Setup Gin Router
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust this for production
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Public Routes
	r.POST("/api/quote", handlers.HandleQuoteRequest(db, cfg))
	r.POST("/api/consultation", handlers.HandleConsultationRequest(db, cfg))
	r.POST("/api/login", handlers.Login(db, cfg))

	// Protected Admin Routes
	admin := r.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware(cfg.JwtSecret))
	{
		admin.GET("/quotes", handlers.GetAllQuotes(db))
		admin.DELETE("/quotes/:id", handlers.DeleteQuote(db))

		admin.GET("/consultations", handlers.GetAllConsultations(db))

		admin.GET("/users", handlers.GetAllUsers(db))
		admin.POST("/users", handlers.CreateUser(db))
		admin.DELETE("/users/:id", handlers.DeleteUser(db))
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 4. Start Server
	log.Printf("Server starting on port %s...", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
