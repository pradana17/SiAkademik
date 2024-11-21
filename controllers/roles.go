package controllers

import (
	"SiAkademik/models"
	"SiAkademik/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	var role models.Role

	// Bind the request body to the user model
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Call the service to create the user
	if err := services.CreateRole(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Role created successfully"})
}

func DeleteRole(c *gin.Context) {

	roleIDParam := c.Param("id")

	// Konversi dari string ke uint
	roleID, err := strconv.ParseUint(roleIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}
	if err := services.DeleteRole(uint(roleID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}

// func GetRoleByID(c *gin.Context) {
// 	roleID, exists := c.Get("roleid")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Role ID not found"})
// 		return
// 	}

// 	// Pastikan roleID adalah tipe uint
// 	roleIDUint, ok := roleID.(uint)
// 	if !ok {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Role ID type"})
// 		return
// 	}

// 	// Panggil service untuk mengambil role berdasarkan ID
// 	role, err := services.GetRoleByID(roleIDUint)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
// 		return
// 	}

// 	// Mengembalikan data role
// 	c.JSON(http.StatusOK, gin.H{
// 		"id":   role.ID,
// 		"name": role.Name,
// 	})

// 	c.Set("rolename", role.Name)
// 	c.Next()
// }
