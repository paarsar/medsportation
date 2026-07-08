package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"medsportation-be/internal/config"
	"medsportation-be/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/resend/resend-go/v2"
	"gorm.io/gorm"
)

// QuoteRequest represents the incoming JSON payload from the frontend
type QuoteRequest struct {
	OrganizationName    string   `json:"organizationName" binding:"required"`
	OrganizationType    string   `json:"organizationType" binding:"required"`
	ContactPerson       string   `json:"contactPerson" binding:"required"`
	Email               string   `json:"email" binding:"required,email"`
	Phone               string   `json:"phone" binding:"required"`
	ServiceType         string   `json:"serviceType" binding:"required"`
	PickupAddress       string   `json:"pickupAddress" binding:"required"`
	DeliveryAddress     string   `json:"deliveryAddress" binding:"required"`
	SpecialRequirements []string `json:"specialRequirements"`
	AdditionalNotes     string   `json:"additionalNotes"`
}

func HandleQuoteRequest(db *gorm.DB, cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req QuoteRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 1. Save to Database
		quote := models.Quote{
			OrganizationName:    req.OrganizationName,
			OrganizationType:    req.OrganizationType,
			ContactPerson:       req.ContactPerson,
			Email:               req.Email,
			Phone:               req.Phone,
			ServiceType:         req.ServiceType,
			PickupAddress:       req.PickupAddress,
			DeliveryAddress:     req.DeliveryAddress,
			SpecialRequirements: strings.Join(req.SpecialRequirements, ", "),
			AdditionalNotes:     req.AdditionalNotes,
		}

		if err := db.Create(&quote).Error; err != nil {
			log.Printf("db.Create: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save quote request"})
			return
		}

		// 2. Send Email via Resend
		if cfg.ResendApiKey != "" {
			client := resend.NewClient(cfg.ResendApiKey)
			subject := fmt.Sprintf("New Quote Request: %s", req.OrganizationName)
			htmlContent := fmt.Sprintf(`
				<h2>New Quote Request Received</h2>
				<p><strong>Organization:</strong> %s (%s)</p>
				<p><strong>Contact Person:</strong> %s</p>
				<p><strong>Email:</strong> %s</p>
				<p><strong>Phone:</strong> %s</p>
				<p><strong>Service Type:</strong> %s</p>
				<p><strong>Pickup Address:</strong> %s</p>
				<p><strong>Delivery Address:</strong> %s</p>
				<p><strong>Special Requirements:</strong> %s</p>
				<p><strong>Additional Notes:</strong> %s</p>
			`, req.OrganizationName, req.OrganizationType, req.ContactPerson, req.Email, req.Phone,
				req.ServiceType, req.PickupAddress, req.DeliveryAddress,
				strings.Join(req.SpecialRequirements, ", "), req.AdditionalNotes)

			params := &resend.SendEmailRequest{
				From:    "Medsportation Quotes <quotes@medsportationlogistics.com>",
				To:      []string{cfg.NotificationEmail},
				Subject: subject,
				Html:    htmlContent,
			}

			sent, err := client.Emails.Send(params)
			if err != nil {
				log.Printf("Resend error (API Key starts with %s...): %v", cfg.ResendApiKey[:5], err)
				// We don't fail the request if email fails, but we log it
			} else {
				log.Printf("Email sent successfully to %s: %s", cfg.NotificationEmail, sent.Id)
			}
		} else {
			log.Println("Resend API key not configured (empty), skipping email notification")
		}

		c.JSON(http.StatusOK, gin.H{"message": "Quote request submitted successfully", "id": quote.ID})
	}
}

func GetAllQuotes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var quotes []models.Quote
		if err := db.Order("created_at desc").Find(&quotes).Error; err != nil {
			log.Printf("db.Find: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch quotes"})
			return
		}
		c.JSON(http.StatusOK, quotes)
	}
}

func DeleteQuote(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Quote{}, id).Error; err != nil {
			log.Printf("db.Delete: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete quote"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Quote deleted successfully"})
	}
}
