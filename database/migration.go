package database

import (
	"SiAkademik/models"
	"log"
)

func MigrateTables() {
	err := DB.AutoMigrate(
		&models.User{},             // Tabel users
		&models.Role{},             // Tabel roles
		&models.UserProfile{},      // Tabel user_profiles
		&models.Auditlog{},         // Tabel log_apis
		&models.Course{},           // Tabel courses
		&models.Semester{},         // Tabel semesters
		&models.CourseEnrollment{}, // Tabel course_enrollments
		&models.Grade{},            // Tabel grades
	)
	if err != nil {
		log.Fatalf("Error during migration: %v\n", err)
	}

	log.Println("Database tables migrated successfully!")
}
