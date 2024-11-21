package models

import "time"

type Grade struct {
	ID         uint      `gorm:"primaryKey"`
	CourseID   uint      `gorm:"not null"`              // Foreign key ke tabel Courses
	Course     Course    `gorm:"foreignKey:CourseID"`   // Relasi many-to-one ke Course
	StudentID  uint      `gorm:"not null"`              // Foreign key ke tabel Users
	Student    User      `gorm:"foreignKey:StudentID"`  // Relasi many-to-one ke User
	SemesterID uint      `gorm:"not null"`              // Foreign key ke tabel Semesters
	Semester   Semester  `gorm:"foreignKey:SemesterID"` // Relasi many-to-one ke Semester
	Grade      string    `gorm:"size:2;not null"`       // Nilai huruf (A, B, C, dst.)
	GradedBy   uint      `gorm:"not null"`              // Foreign key ke tabel Users (Dosen)
	Grader     User      `gorm:"foreignKey:GradedBy"`   // Relasi many-to-one ke User
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

type GradeDetails struct {
	CourseName string `json:"course_name"`
	Grade      string `json:"grade"`
	Credits    int    `json:"credits"`
}

type GPAResponse struct {
	StudentID            uint           `json:"student_id"`
	SemesterName         string         `json:"semester_name"`
	TotalCreditsSemester int            `json:"total_credits_semester"`
	SemesterGPA          float64        `json:"semester_gpa"`
	Grades               []GradeDetails `json:"grades"`
	TotalCredits         int            `json:"total_credits"`
	GPACumulative        float64        `json:"gpa_cumulative"`
}
