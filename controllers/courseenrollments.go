package controllers

import (
	"SiAkademik/models"
	"SiAkademik/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEnrollment(c *gin.Context) {
	var enroll models.CourseEnrollment
	// Bind the request body to the user model
	if err := c.ShouldBindJSON(&enroll); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	student, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": exists})
		return
	}

	enroll.StudentID = student.(uint)

	// Call the service to create the enrollment
	if err := services.CreateEnrollment(&enroll); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Enrollment created successfully"})
}
