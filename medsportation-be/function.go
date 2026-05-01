package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
)

// QuoteRequest represents the incoming JSON payload from the frontend
type QuoteRequest struct {
	OrganizationName    string   `json:"organizationName"`
	OrganizationType    string   `json:"organizationType"`
	ContactPerson       string   `json:"contactPerson"`
	Email               string   `json:"email"`
	Phone               string   `json:"phone"`
	ServiceType         string   `json:"serviceType"`
	PickupAddress       string   `json:"pickupAddress"`
	DeliveryAddress     string   `json:"deliveryAddress"`
	SpecialRequirements []string `json:"specialRequirements"`
	AdditionalNotes     string   `json:"additionalNotes"`
}

// RequestQuote is the HTTP Cloud Function entry point
func RequestQuote(w http.ResponseWriter, r *http.Request) {
	// 1. Handle CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Max-Age", "3600")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. Parse the request body
	var req QuoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("json.NewDecoder.Decode: %v", err)
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// 3. Send the email via SMTP
	if err := sendEmailSmtp(req); err != nil {
		log.Printf("sendEmailSmtp: %v", err)
		http.Error(w, "Error sending email notification", http.StatusInternalServerError)
		return
	}

	// 4. Respond to the frontend
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Quote request submitted successfully"})
}

func sendEmailSmtp(req QuoteRequest) error {
	// Configuration from environment variables
	from := os.Getenv("SMTP_EMAIL")    // e.g. info@yourdomain.com
	pass := os.Getenv("SMTP_PASSWORD") // Your GoDaddy/M365 password
	host := os.Getenv("SMTP_HOST")     // e.g. "smtp.office365.com" or "smtpout.secureserver.net"
	port := os.Getenv("SMTP_PORT")     // e.g. "587"
	to := os.Getenv("NOTIFICATION_RECIPIENT")

	if to == "" {
		to = from // Default to sending to self
	}

	// Construct the email message
	subject := fmt.Sprintf("Subject: New Quote Request: %s\r\n", req.OrganizationName)
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf(
		"New Quote Request Details:\n\n"+
			"Organization: %s (%s)\n"+
			"Contact: %s\n"+
			"Email: %s\n"+
			"Phone: %s\n"+
			"Service: %s\n"+
			"Pickup: %s\n"+
			"Delivery: %s\n"+
			"Requirements: %s\n"+
			"Notes: %s\n",
		req.OrganizationName, req.OrganizationType, req.ContactPerson, req.Email, req.Phone,
		req.ServiceType, req.PickupAddress, req.DeliveryAddress, 
		strings.Join(req.SpecialRequirements, ", "), req.AdditionalNotes,
	)

	msg := []byte(subject + mime + body)

	// Authentication
	auth := smtp.PlainAuth("", from, pass, host)

	// Send email
	addr := fmt.Sprintf("%s:%s", host, port)
	return smtp.SendMail(addr, auth, from, []string{to}, msg)
}
