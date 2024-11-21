package controllers

import (
	"SiAkademik/models"
	"SiAkademik/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
func CreateGrade(c *gin.Context) {
	var grade models.Grade
	// Bind the request body to the user model
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lecturerID, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Call the service to create grade
	if err := services.CreateGrade(lecturerID.(uint), &grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Grade created successfully"})
}
*/

func CreateGrade(c *gin.Context) {
	var input models.Grade

	// Bind data dari request body ke input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idLec, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ID lecturer not found"})
		return
	}
	input.GradedBy = idLec.(uint)

	// Panggil service untuk membuat grade
	err := services.CreateGrade(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Grade successfully created"})
}
