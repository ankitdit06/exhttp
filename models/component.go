package models

import (
	"time"

	"github.com/google/uuid"
)

type Component struct {
	Id            uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // Use UUID as ID
	Name          string    `json:"name" gorm:"size:100"`
	Version       string    `json:"version" gorm:"size:100"`
	URL           string    `json:"url" gorm:"size:100"`
	Platform      string    `json:"platform" gorm:"size:100"`
	DocumentRef   string    `json:"document_ref" gorm:"size:100"`
	APIRef        string    `json:"api_ref" gorm:"size:100"`
	ResponsibleID uuid.UUID `json:"team_id" gorm:"not null"`              // Foreign Key for Team
	Team          Team      `json:"team" gorm:"foreignKey:ResponsibleID"` // Relation to Team
	Metadata      JSONB     `json:"metadata" json:"metadata" gorm:"type:jsonb"`
	ServiceID     uuid.UUID `json:"service_id" gorm:"not null"`         // Foreign Key for Service
	Service       Service   `json:"-" gorm:"foreignKey:ServiceID"`      // Relation to Service
	LastUpdated   time.Time `json:"last_updated" gorm:"autoUpdateTime"` // This will automatically update when the row is modified

}
