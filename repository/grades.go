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
