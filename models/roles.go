package models

import "time"

type Role struct {
	ID         uint      `gorm:"primaryKey"`
	Name       string    `gorm:"unique;not null"` // Nama role (admin, lecturer, student)
	CreatedBy  string    `gorm:"size:50"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	ModifiedBy string    `gorm:"size:50"`
	ModifiedAt time.Time `gorm:"autoUpdateTime"`
	Users      []User    `gorm:"foreignKey:RoleID"` // Relasi one-to-many dengan User
}
