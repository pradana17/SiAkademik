package repository

import (
	"SiAkademik/database"
	"SiAkademik/models"
)

func CreateEnrollment(cour *models.CourseEnrollment) error {
	err := database.DB.Create(cour).Error
	if err != nil {
		return err
	}
	return nil
}

func GetEnrollment(courseID, studentID, semesterID uint) ([]models.CourseEnrollment, error) {
	var enroll []models.CourseEnrollment
	err := database.DB.Where("course_id = ? and student_id=? and semester_id=?", courseID, studentID, semesterID).First(enroll).Error
	if err != nil {
		return nil, err
	}
	return enroll, nil
}

func CheckStudentEnroll(courseID, studentID, semesterID uint) (bool, error) {
	var count int64
	err := database.DB.Model(&models.CourseEnrollment{}).
		Where("course_id = ? AND student_id = ? AND semester_id = ?", courseID, studentID, semesterID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CheckExistingEnroll(courseID, studentID, semesterID uint) error {
	var existingGrade models.Grade
	err := database.DB.Where("course_id = ? AND student_id = ? AND semester_id = ?", courseID, studentID, semesterID).
		First(&existingGrade).Error
	if err != nil {
		return err
	}
	return nil
}

func GetStudentEnrollment(studentID, semesterID uint) ([]models.CourseEnrollment, error) {
	var enroll []models.CourseEnrollment
	err := database.DB.Where("student_id = ? AND semester_id=?", studentID, semesterID).Find(&enroll).Error
	if err != nil {
		return nil, err
	}
	return enroll, nil
}
