package repository

import (
	"SiAkademik/database"
	"SiAkademik/models"
)

func CreateCourse(course *models.Course) error {
	err := database.DB.Create(course).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCourseByLectureId(LecturerID uint) ([]models.Course, error) {
	var courses []models.Course
	err := database.DB.Where("lecturer_id = ?", LecturerID).Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func GetCourseById(courseID uint) ([]models.Course, error) {
	var courses []models.Course
	err := database.DB.Where("id = ?", courseID).Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func CheckCourse(courseID uint) error {
	var exist models.Course
	err := database.DB.Where("id = ?", courseID).
		First(&exist).Error
	if err != nil {
		return err
	}
	return nil
}
