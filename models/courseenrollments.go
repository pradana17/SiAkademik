package models

import "time"

type CourseEnrollment struct {
	ID         uint      `gorm:"primaryKey"`
	StudentID  uint      `gorm:"not null"`              // Foreign key ke tabel Users
	Student    User      `gorm:"foreignKey:StudentID"`  // Relasi many-to-one ke User
	CourseID   uint      `gorm:"not null"`              // Foreign key ke tabel Courses
	Course     Course    `gorm:"foreignKey:CourseID"`   // Relasi many-to-one ke Course
	SemesterID uint      `gorm:"not null"`              // Foreign key ke tabel Semesters
	Semester   Semester  `gorm:"foreignKey:SemesterID"` // Relasi many-to-one ke Semester
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
