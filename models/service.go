package models

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	Name                 string                   `json:"name" gorm:"size:100"`
	Id                   uuid.UUID                `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // Use UUID as ID
	SecAudit             SecAudit                 `json:"audits" gorm:"foreignKey:ServiceID"`
	OwnerID              uuid.UUID                `json:"teamid" gorm:"not null"` // Foreign Key for Team
	Metadata             JSONB                    `json:"metadata" gorm:"type:jsonb"`
	Release              string                   `json:"release" gorm:"size:100"`
	PublicFacing         string                   `json:"publicfacing" gorm:"size:100"`
	Team                 Team                     `gorm:"foreignKey:OwnerID"` // Establish the relationship
	SecerviceSecurityMap []ServiceSecurityControl `json:"securitycontrols" gorm:"foreignKey:ServiceID"`
	LastUpdated          time.Time                `json:"last_updated" gorm:"autoUpdateTime"` // This will automatically update when the row is modified
}
