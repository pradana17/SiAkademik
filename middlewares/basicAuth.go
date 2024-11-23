package middlewares

import (
	"SiAkademik/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Cek apakah username dan password kosong
		if username == "" || password == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username and password are requirederrrr"})
			c.Abort()
			return
		}

		//AuthenticateUser untuk autentikasi
		user, err := services.AuthenticateUser(username, password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		//Get role by role id
		role, err := services.GetRoleByID(user.RoleID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Jika autentikasi berhasil, simpan user ke context dan lanjutkan ke handler berikutnya
		c.Set("userid", user.ID)
		c.Set("roleid", user.RoleID)
		c.Set("rolename", role.Name)
		c.Set("username", user.Username)

		c.Next()
	}
}
