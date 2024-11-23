package controllers

import (
	"SiAkademik/models"
	"SiAkademik/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	// Ambil userID dari auth

	userid, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID Not Found"})
		return
	}

	// Panggil service untuk mengambil user berdasarkan ID
	up, err := services.GetUserProfile(userid.(uint))
	if err != nil {
		//c.JSON(http.StatusNotFound, gin.H{"error": "User Profile not found"})
		CreateUserProfile(c)
		return

	}

	// Mengembalikan data user
	c.JSON(http.StatusOK, gin.H{
		"User Profile": up,
	})
}

func UpdateUserProfile(c *gin.Context) {
	var userProfile models.UserProfile
	// Ambil data user yang akan diupdate dari body request
	if err := c.ShouldBindJSON(&userProfile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dataxxx"})
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
	userProfile.ModifiedBy = username.(string)

	if x := services.UpdateUserProfile(userid.(uint), userProfile); x != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": x.Error()})
		return
	}

	// Mengembalikan response sukses
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func CreateUserProfile(c *gin.Context) {
	var userProfile models.UserProfile
	userid, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Userid not found"})
		return
	}

	userProfile.UserID = userid.(uint)
	userProfile.CreatedBy = "system"
	// Call the service to create the user
	if err := services.CreateUserProfile(&userProfile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"user profile": &userProfile})
}
