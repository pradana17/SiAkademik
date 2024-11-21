package controllers

import (
	"SiAkademik/models"
	"SiAkademik/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	// Bind the request body to the user model
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found"})
		return
	}

	user.CreatedBy = username.(string)

	// Call the service to create the user
	if err := services.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	// Ambil data user yang akan diupdate dari body request
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	userid, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID Not Found"})
		return
	}

	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found"})
		return
	}
	user.ModifiedBy = username.(string)

	if x := services.UpdateUser(userid.(uint), user); x != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": x.Error()})
		return
	}

	// Mengembalikan response sukses
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
