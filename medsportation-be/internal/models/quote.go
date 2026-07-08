package models

import (
	"time"

	"gorm.io/gorm"
)

type Quote struct {
	ID                  uint           `gorm:"primaryKey" json:"id"`
	CreatedAt           time.Time      `json:"createdAt"`
	UpdatedAt           time.Time      `json:"updatedAt"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"-"`
	OrganizationName    string         `json:"organizationName"`
	OrganizationType    string         `json:"organizationType"`
	ContactPerson       string         `json:"contactPerson"`
	Email               string         `json:"email"`
	Phone               string         `json:"phone"`
	ServiceType         string         `json:"serviceType"`
	PickupAddress       string         `json:"pickupAddress"`
	DeliveryAddress     string         `json:"deliveryAddress"`
	SpecialRequirements string         `json:"specialRequirements"` // Stored as comma-separated string
	AdditionalNotes     string         `json:"additionalNotes"`
}
