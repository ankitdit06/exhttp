package models

import "github.com/google/uuid"

type ServiceSecurityControl struct {
	ID                uuid.UUID       `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ServiceID         uuid.UUID       `gorm:"not null"`
	SecurityControlID uuid.UUID       `gorm:"not null"`
	Status            string          `gorm:"not null"`
	Service           Service         `json:"-" gorm:"foreignKey:ServiceID;constraint:OnDelete:CASCADE"`
	SecurityControl   SecurityControl `gorm:"foreignKey:SecurityControlID;constraint:OnDelete:CASCADE"`
}
