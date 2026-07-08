package handlers

import (
	"fmt"
	"log"
	"net/http"

	"medsportation-be/internal/config"
	"medsportation-be/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/resend/resend-go/v2"
	"gorm.io/gorm"
)

type ConsultationRequest struct {
	Name              string `json:"name" binding:"required"`
	Email             string `json:"email" binding:"required,email"`
	Phone             string `json:"phone" binding:"required"`
	Organization      string `json:"organization"`
	InterestedService string `json:"interestedService"`
	Message           string `json:"message"`
}

func HandleConsultationRequest(db *gorm.DB, cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ConsultationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 1. Save to Database
		consultation := models.Consultation{
			Name:              req.Name,
			Email:             req.Email,
			Phone:             req.Phone,
			Organization:      req.Organization,
			InterestedService: req.InterestedService,
			Message:           req.Message,
		}

		if err := db.Create(&consultation).Error; err != nil {
			log.Printf("db.Create: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save consultation request"})
			return
		}

		// 2. Send Email via Resend
		if cfg.ResendApiKey != "" {
			client := resend.NewClient(cfg.ResendApiKey)
			subject := fmt.Sprintf("New Consultation Request: %s", req.Name)
			htmlContent := fmt.Sprintf(`
				<h2>New Consultation Request Received</h2>
				<p><strong>Name:</strong> %s</p>
				<p><strong>Email:</strong> %s</p>
				<p><strong>Phone:</strong> %s</p>
				<p><strong>Organization:</strong> %s</p>
				<p><strong>Interested Service:</strong> %s</p>
				<p><strong>Message:</strong> %s</p>
			`, req.Name, req.Email, req.Phone, req.Organization, req.InterestedService, req.Message)

			params := &resend.SendEmailRequest{
				From:    "Medsportation Notifications <quotes@medsportationlogistics.com>",
				To:      []string{cfg.NotificationEmail},
				Subject: subject,
				Html:    htmlContent,
			}

			sent, err := client.Emails.Send(params)
			if err != nil {
				log.Printf("Resend error (Consultation): %v", err)
			} else {
				log.Printf("Consultation email sent successfully: %s", sent.Id)
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "Consultation request submitted successfully", "id": consultation.ID})
	}
}

func GetAllConsultations(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var consultations []models.Consultation
		if err := db.Order("created_at desc").Find(&consultations).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch consultations"})
			return
		}
		c.JSON(http.StatusOK, consultations)
	}
}
