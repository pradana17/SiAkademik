package services

import (
	"SiAkademik/models"
	"SiAkademik/repository"
	"errors"
)

func CreateCourse(course *models.Course) error {
	// validasi sebelum membuat course, pastikan id lecturer memiliki role dosen

	lecturer, err := GetUserByID(course.LecturerID)
	if err != nil {
		return errors.New("lecturer id not found")
	}

	role, err := GetRoleByID(lecturer.RoleID)
	if err != nil {
		return errors.New("role not found")
	}

	if role.Name != "dosen" {
		return errors.New("user not lecturer")
	}

	// Memanggil fungsi repository untuk menyimpan data user
	return repository.CreateCourse(course)
}

func GetCourseByLectureId(LecturerID uint) ([]models.Course, error) {

	courses, err := repository.GetCourseByLectureId(LecturerID)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func GetCourseById(ID uint) ([]models.Course, error) {

	courses, err := repository.GetCourseById(ID)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func CourseResponse(courses []models.Course) []models.CourseResponse {
	var responses []models.CourseResponse
	for _, course := range courses {
		responses = append(responses, models.CourseResponse{
			ID:       course.ID,
			Name:     course.Name,
			Code:     course.Code,
			Schedule: course.Schedule,
			Credits:  course.Credits,
		})
	}
	return responses
}

func GetStudentCourse(studentID, semesterID uint) ([]models.Course, error) {
	var courses []models.Course
	enroll, err := repository.GetStudentEnrollment(studentID, semesterID)
	if err != nil {
		return nil, err
	}
	//return enroll, nil
	for _, enrollment := range enroll {
		course, err := repository.GetCourseById(enrollment.CourseID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course...)
	}
	return courses, nil

}
