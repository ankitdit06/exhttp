package models

import "github.com/google/uuid"

// Team represents a group of individuals managing services and components.
type Team struct {
	Id   uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name string    `json:"name" gorm:"size:100"`
}
