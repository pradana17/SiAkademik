package services

import (
	"SiAkademik/models"
	"SiAkademik/repository"
	"errors"
	"log"
)

func CreateSemester(sem *models.Semester) error {
	// Memanggil fungsi repository untuk menyimpan data
	err := repository.CreateSemester(sem)
	if err != nil {
		log.Printf("cannot create semester : %v", err)
		return errors.New("error create semester")
	}
	return nil
}

func GetActiveSemester() (*models.Semester, error) {
	var sem models.Semester
	err := repository.GetActiveSemester(&sem)
	if err != nil {
		return nil, errors.New("error get active semester")
	}
	return &sem, nil
}
