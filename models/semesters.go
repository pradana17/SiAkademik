package models

import "time"

type Semester struct {
	ID         uint      `gorm:"primaryKey"`
	Name       string    `gorm:"unique;not null"` // Nama semester (e.g., Semester 1, Semester 2)
	StartDate  time.Time `gorm:"not null"`
	EndDate    time.Time `gorm:"not null"`
	IsActive   bool      `gorm:"default:false"` // Status aktif semester
	CreatedBy  string    `gorm:"size:50"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	ModifiedBy string    `gorm:"size:50"`
	ModifiedAt time.Time `gorm:"autoUpdateTime"`
}
