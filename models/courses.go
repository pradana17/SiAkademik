package models

import "time"

type Course struct {
	ID          uint               `gorm:"primaryKey"`
	Name        string             `gorm:"not null"`               // Nama mata kuliah
	Code        string             `gorm:"size:5;unique;not null"` // Kode mata kuliah
	LecturerID  uint               `gorm:"not null"`               // Foreign key ke tabel Users (Lecturer)
	Lecturer    User               `gorm:"foreignKey:LecturerID"`  // Relasi many-to-one dengan User
	Credits     int                `gorm:"not null"`               // SKS
	Schedule    string             `gorm:"size:100"`
	CreatedBy   string             `gorm:"size:50"`
	CreatedAt   time.Time          `gorm:"autoCreateTime"`
	ModifiedBy  string             `gorm:"size:50"`
	ModifiedAt  time.Time          `gorm:"autoUpdateTime"`
	Enrollments []CourseEnrollment `gorm:"foreignKey:CourseID"` // Relasi one-to-many dengan CourseEnrollment
}

type CourseResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"` // Nama mata kuliah
	Code     string `json:"code"` // Kode mata kuliah
	Schedule string `json:"schedule"`
	Credits  int    `json:"credits"`
}
