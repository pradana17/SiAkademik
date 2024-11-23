package middlewares

import (
	"SiAkademik/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is missing or invalid"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := services.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		user, err := services.GetUserByID(claims.UserID)
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

		// Set user ID ke context
		c.Set("userid", claims.UserID)
		c.Set("username", user.Username)
		c.Set("roleid", user.RoleID)
		c.Set("rolename", role.Name)

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
