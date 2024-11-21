package services

import (
	"SiAkademik/models"
	"SiAkademik/repository"
	"errors"
	"log"
)

func CreateEnrollment(enroll *models.CourseEnrollment) error {

	semester, err := GetActiveSemester()
	if err != nil {
		return errors.New("error get active semester")
	}

	enroll.SemesterID = semester.ID

	existing, err := repository.CheckStudentEnroll(enroll.CourseID, enroll.StudentID, enroll.SemesterID)
	if err != nil {
		log.Print("error", err)
		return errors.New("cannot check student enroll")
	}

	if existing {
		log.Printf("already enrolled in the course for this semester")
		return errors.New("already enrolled in the course for this semester")
	}

	existCourse := repository.CheckCourse(enroll.CourseID)
	if existCourse != nil {
		log.Printf("course not exists")
		return errors.New("course not exists")
	}

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
