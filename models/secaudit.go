package models

import "github.com/google/uuid"

type SecAudit struct {
	Id                  uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // Use UUID as ID
	LastReviewed        string    `json:"last_reviewed" gorm:"size:100"`
	ReviewedBy          string    `json:"reviewed_by" gorm:"size:100"`
	ReviewSummary       string    `json:"reviewsummary" gorm:"size:100"`
	DataClassification  string    `json:"data_classification" gorm:"size:100"`
	DataRetentionPolicy string    `json:"data_retention_policy" gorm:"size:100"`
	Confidentiality     string    `json:"confidentiality" gorm:"size:100"`
	Avaibility          string    `json:"avaibility" gorm:"size:100"`
	Integrity           string    `json:"integrity" gorm:"size:100"`
	BusinessImpact      string    `json:"business_impact" gorm:"size:100"`
	Status              string    `gorm:"size:100"`
	ServiceID           uuid.UUID `gorm:"not null"` // Foreign Key for Service

}
