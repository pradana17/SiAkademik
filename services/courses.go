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

func CourseResponse(courses []models.Course) []models.CourseResponse {
	var responses []models.CourseResponse
	for _, course := range courses {
		responses = append(responses, models.CourseResponse{
			ID:   course.ID,
			Name: course.Name,
			Code: course.Code,
		})
	}
	return responses
}
