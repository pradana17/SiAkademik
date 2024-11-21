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

		// Gunakan fungsi AuthenticateUser untuk autentikasi
		user, err := services.AuthenticateUser(username, password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

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

func CheckRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil role dari context
		role, exists := c.Get("rolename")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role name not found"})
			c.Abort()
			return
		}

		// Cek apakah role yang didapat sesuai dengan requiredRole
		if role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden Access"})
			c.Abort()
			return
		}

		// Jika role sesuai, lanjutkan ke handler berikutnya
		c.Next()
	}
}
