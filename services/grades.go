package services

import (
	"SiAkademik/database"
	"SiAkademik/models"
	"SiAkademik/repository"
	"errors"
	"fmt"
	"log"
)

func CreateGrade(sem *models.Grade) error {
	// Cek apakah mahasiswa sudah terdaftar di kursus dan semester ini
	isEnrolled, err := repository.CheckStudentEnroll(sem.CourseID, sem.StudentID, sem.SemesterID)
	if err != nil {
		log.Printf("error check enrollment : %v", err)
		return errors.New("error check enrollment")
	}
	if !isEnrolled {
		log.Printf("student is not enrolled in the course for this semester")
		return errors.New("student is not enrolled in the course for this semester")
	}

	// Cek apakah grade sudah ada

	existing := repository.CheckExistingEnroll(sem.CourseID, sem.StudentID, sem.SemesterID)
	if existing == nil {
		log.Printf("grade already exists for the student in this course and semester")
		return errors.New("grade already exists for the student in this course and semester")
	}

	//cek apakah dosen mengajar matakuliah itu
	courses, err := repository.GetCourseById(sem.CourseID)
	if err != nil {
		return err
	}
	for _, checheckLec := range courses {
		if checheckLec.LecturerID != sem.GradedBy {
			return errors.New("you are not lecturer in this course")
		}
	}

	err = database.DB.Create(&sem).Error
	if err != nil {
		log.Print("failed to create grade: %w", err)
		return errors.New("failed to create grade")
	}
	return nil
}

func GetGrade(studentID uint) ([]models.Grade, error) {

	grade, err := repository.GetGrade(studentID)
	if err != nil {
		return nil, err
	}
	return grade, nil
}

func GetGPA(studentID, semesterID uint) (*models.GPAResponse, error) {

	semesters, err := repository.GetSemesterById(semesterID)
	if err != nil {
		log.Print("error", err)
		return nil, err
	}
	if len(semesters) == 0 {
		return nil, errors.New("semester id not found")
	}
	semesterName := semesters[0].Name

	grades, err := repository.GetGrade(studentID)
	if err != nil {
		log.Print("error", err)
		return nil, err
	}
	//filter grade by semester id
	var semesterGrades []models.Grade
	for _, grade := range grades {
		if grade.SemesterID == semesterID {
			semesterGrades = append(semesterGrades, grade)
		}
	}

	var totalCreditsSemester int
	var totalGradePoints float64
	gradeDetails := []models.GradeDetails{}

	for _, grade := range semesterGrades {
		courses, err := repository.GetCourseById(grade.CourseID)
		if err != nil {
			return nil, err
		}
		if len(courses) == 0 {
			return nil, fmt.Errorf("course not found for ID: %d", grade.CourseID)
		}
		course := courses[0]

		gradePoint := convertGradeToNumeric(grade.Grade)
		totalCreditsSemester += course.Credits
		totalGradePoints += float64(course.Credits) * gradePoint

		gradeDetails = append(gradeDetails, models.GradeDetails{
			CourseName: course.Name,
			Grade:      grade.Grade,
			Credits:    course.Credits,
		})
	}

	var semesterGPA float64
	if totalCreditsSemester > 0 {
		semesterGPA = totalGradePoints / float64(totalCreditsSemester)
	}

	// Step 4: Calculate cumulative GPA
	var totalCreditsCumulative int
	var totalGradePointsCumulative float64

	for _, grade := range grades {
		courses, err := repository.GetCourseById(grade.CourseID)
		if err != nil {
			return nil, err
		}
		if len(courses) == 0 {
			continue
		}
		course := courses[0]

		gradePoint := convertGradeToNumeric(grade.Grade)
		totalCreditsCumulative += course.Credits
		totalGradePointsCumulative += float64(course.Credits) * gradePoint
	}

	var cumulativeGPA float64
	if totalCreditsCumulative > 0 {
		cumulativeGPA = totalGradePointsCumulative / float64(totalCreditsCumulative)
	}

	// Step 5: Build response
	return &models.GPAResponse{
		StudentID:            studentID,
		SemesterName:         semesterName,
		TotalCreditsSemester: totalCreditsSemester,
		SemesterGPA:          semesterGPA,
		Grades:               gradeDetails,
		TotalCredits:         totalCreditsCumulative,
		GPACumulative:        cumulativeGPA,
	}, nil

}

func convertGradeToNumeric(grade string) float64 {
	switch grade {
	case "A":
		return 4.0
	case "A-":
		return 3.75
	case "B+":
		return 3.25
	case "B":
		return 3.0
	case "B-":
		return 2.75
	case "C+":
		return 2.25
	case "C":
		return 2.0
	case "C-":
		return 1.75
	case "D":
		return 1.0
	case "E":
		return 0.0
	default:
		return 0.0
	}
}
