package middleware

import (
	"IronOps/internal/service"
	"net/http"
	"github.com/gin-gonic/gin"
	"IronOps/internal/model"
)

// AuthMiddleware simulates authentication by checking X-User header
// In real world, verify JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetHeader("X-User")
		if username == "" {
			// For MVP/Test, if no header, allow for now ONLY if testing? 
			// No, strict mode is better for "Platform" feel.
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized - missing X-User header"})
			return
		}

		user, err := service.GetUserByUsername(username)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

// RoleMiddleware checks if user has required role
func RoleMiddleware(allowedRoles ...model.RoleType) gin.HandlerFunc {
	return func(c *gin.Context) {
		userVal, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		user := userVal.(*model.User)

		// Admin has all permissions
		if user.Role == model.RoleAdmin {
			c.Next()
			return
		}

		allowed := false
		for _, role := range allowedRoles {
			if user.Role == role {
				allowed = true
				break
			}
		}

		if !allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}
		c.Next()
	}
}
