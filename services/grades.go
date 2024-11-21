package services

import (
	"SiAkademik/database"
	"SiAkademik/models"
	"SiAkademik/repository"
	"errors"
	"log"
)

func CreateGrade(sem *models.Grade) error {
	// Cek apakah mahasiswa sudah terdaftar di kursus dan semester ini
	isEnrolled, err := repository.CheckStudentEnroll(sem.CourseID, sem.StudentID, sem.SemesterID)
	if err != nil {
		log.Printf("error check enrollment : %v", err)
		return errors.New("error check enrollment")
	}
	if !isEnrolled {
		log.Printf("student is not enrolled in the course for this semester")
		return errors.New("student is not enrolled in the course for this semester")
	}

	// Cek apakah grade sudah ada

	existing := repository.CheckExistingEnroll(sem.CourseID, sem.StudentID, sem.SemesterID)
	if existing == nil {
		log.Printf("grade already exists for the student in this course and semester")
		return errors.New("grade already exists for the student in this course and semester")
	}

	err = database.DB.Create(&sem).Error
	if err != nil {
		log.Print("failed to create grade: %w", err)
		return errors.New("failed to create grade")
	}
	return nil
}
