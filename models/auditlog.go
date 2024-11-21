package models

import "time"

type Auditlog struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     *uint     `gorm:"index"` // Foreign key ke User.ID
	Method     string    `gorm:"size:10;not null"`
	Endpoint   string    `gorm:"size:255;not null"`
	Request    string    `gorm:"type:json"`
	Response   string    `gorm:"type:json"`
	StatusCode int       `gorm:"not null"`
	Duration   float64   `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
