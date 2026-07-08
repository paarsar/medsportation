package handlers

import (
	"log"
	"net/http"
	"time"

	"medsportation-be/internal/config"
	"medsportation-be/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(db *gorm.DB, cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("Login: Invalid request body: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		log.Printf("Login attempt for user: %s", req.Username)

		var user models.User
		if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
			log.Printf("Login: User not found: %s", req.Username)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Verify password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			log.Printf("Login: Password mismatch for user: %s", req.Username)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Generate JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString([]byte(cfg.JwtSecret))
		if err != nil {
			log.Printf("Login: Token signing error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		log.Printf("Login successful for user: %s", req.Username)
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}

func GetAllUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		user := models.User{
			Username: req.Username,
			Password: string(hashedPassword),
		}

		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		// Prevent self-deletion if we tracked current user ID, but for now simple delete
		if err := db.Delete(&models.User{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}

// CreateInitialAdmin creates a default admin if none exist
func CreateInitialAdmin(db *gorm.DB) error {
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		admin := models.User{
			Username: "admin",
			Password: string(hashedPassword),
		}
		return db.Create(&admin).Error
	}
	return nil
}
