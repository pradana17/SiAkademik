package controllers

import (
	"SiAkademik/models"
	"SiAkademik/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func GetGPA(c *gin.Context) {
	var input models.Grade
	// Bind data dari request body ke input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studentID, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauhorized Student Id"})
		return
	}

	grade, err := services.GetGPA(studentID.(uint), uint(input.SemesterID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"data": grade})

}
