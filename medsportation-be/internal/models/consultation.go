package models

import (
	"time"

	"gorm.io/gorm"
)

type Consultation struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	Name             string         `json:"name"`
	Email            string         `json:"email"`
	Phone            string         `json:"phone"`
	Organization     string         `json:"organization"`
	InterestedService string         `json:"interestedService"`
	Message          string         `json:"message"`
}
