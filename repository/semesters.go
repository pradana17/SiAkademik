package repository

import (
	"SiAkademik/database"
	"SiAkademik/models"
	"fmt"
	"log"
)

func CreateSemester(sem *models.Semester) error {
	err := database.DB.Create(sem).Error
	if err != nil {
		return err
	}
	return nil
}

func GetActiveSemester(sem *models.Semester) error {
	var i = true
	var count int64
	err := database.DB.Model(&models.Semester{}).Where("Is_Active = ?", i).Count(&count).Error
	if err != nil {
		return fmt.Errorf("failed to count active records: %w", err)
	}

	if count == 0 {
		return fmt.Errorf("no active records found")
	}
	if count > 1 {
		return fmt.Errorf("multiple active records found")
	}

	errs := database.DB.Where("Is_Active = ?", i).First(sem).Error
	if errs != nil {
		return errs
	}
	return nil
}

func GetSemesterById(semesterID uint) ([]models.Semester, error) {
	var semester []models.Semester
	err := database.DB.Where("id = ?", semesterID).Find(&semester).Error
	if err != nil {
		return nil, err
	}
	log.Print("data", semesterID)
	return semester, nil
}
