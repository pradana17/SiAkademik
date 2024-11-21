package controllers

import (
	"SiAkademik/models"
	"SiAkademik/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSemester(c *gin.Context) {
	var sem models.Semester
	// Bind the request body to the model
	if err := c.ShouldBindJSON(&sem); err != nil {
		//log.Error("gagal parsing %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found"})
		return
	}
	sem.CreatedBy = username.(string)

	// Call the service to create
	if err := services.CreateSemester(&sem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Semester created successfully"})
}
