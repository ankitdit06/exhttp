package models

import "github.com/google/uuid"

type SecurityControl struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`

	Name                   string                   `json:"name" gorm:"size:100"` // Example: "Encryption in Transit"
	Risk                   string                   `json:"risk" gorm:"size:100"`
	ServiceSecurityControl []ServiceSecurityControl `gorm:"foreignKey:SecurityControlID"`
}
