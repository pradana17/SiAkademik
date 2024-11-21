package repository

import (
	"SiAkademik/database"
	"SiAkademik/models"
)

func CreateGrade(grade *models.Grade) error {
	err := database.DB.Create(grade).Error
	if err != nil {
		return err
	}
	return nil
}

func GetGrade(studentID uint) ([]models.Grade, error) {
	var grade []models.Grade
	err := database.DB.Where("student_id = ?", studentID).Find(&grade).Error
	if err != nil {
		return nil, err
	}
	return grade, nil
}
