package models

import "time"

type UserProfile struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint   `gorm:"unique;not null"` // Foreign key ke tabel Users
	User       *User  `gorm:"constraint:OnDelete:CASCADE"`
	Name       string `gorm:"size:100"`
	DOB        *time.Time
	Address    string    `gorm:"type:text"`
	Gender     string    `gorm:"size:1"` // 'M' atau 'F'
	CreatedBy  string    `gorm:"size:50"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	ModifiedBy string    `gorm:"size:50"`
	ModifiedAt time.Time `gorm:"autoUpdateTime"`
}
