package services

import (
	"SiAkademik/models"
	"SiAkademik/repository"
	"errors"
)

func CreateEnrollment(enroll *models.CourseEnrollment) error {

	semester, err := GetActiveSemester()
	if err != nil {
		return errors.New("error get active semester")
	}

	enroll.SemesterID = semester.ID

	// Memanggil fungsi repository untuk menyimpan data user
	return repository.CreateEnrollment(enroll)

}

func GetEnrollment(courseID, studentID, semesterID uint) ([]models.CourseEnrollment, error) {
	enroll, err := repository.GetEnrollment(courseID, studentID, semesterID)
	if err != nil {
		return nil, err
	}
	return enroll, nil
}
