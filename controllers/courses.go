package controllers

import (
	"SiAkademik/models"
	"SiAkademik/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCourse(c *gin.Context) {
	var course models.Course
	// Bind the request body to the user model
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found"})
		return
	}
	course.CreatedBy = username.(string)

	// Call the service to create the user
	if err := services.CreateCourse(&course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Course created successfully"})
}

func GetCourseByLectureId(c *gin.Context) {
	// Ambil userID dari auth

	lecturerID, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauhorized Lecturer Id"})
		return
	}

	// Panggil service untuk mengambil course berdasarkan LecID
	course, err := services.GetCourseByLectureId(lecturerID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	response := services.CourseResponse(course)
	// Mengembalikan data user
	c.JSON(http.StatusOK, gin.H{
		"course": response,
	})
}
