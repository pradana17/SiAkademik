package models

import "time"

type User struct {
	ID          uint        `gorm:"primaryKey"`
	Username    string      `gorm:"unique;not null"`   // Username unik
	Password    string      `gorm:"not null"`          // Password hash
	RoleID      uint        `gorm:"not null"`          // Foreign key ke Role
	Role        Role        `gorm:"foreignKey:RoleID"` // Relasi many-to-one dengan Role
	UserProfile UserProfile `gorm:"foreignKey:UserID"` // Relasi one-to-one dengan UserProfile
	CreatedBy   string      `gorm:"size:50"`
	CreatedAt   time.Time   `gorm:"autoCreateTime"`
	ModifiedBy  string      `gorm:"size:50"`
	ModifiedAt  time.Time   `gorm:"autoUpdateTime"`
}
