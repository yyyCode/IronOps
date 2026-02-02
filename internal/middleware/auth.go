package middleware

import (
	"IronOps/internal/model"
	"IronOps/internal/pkg/response"
	"IronOps/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware simulates authentication by checking X-User header
// In real world, verify JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetHeader("X-User")
		if username == "" {
			// For MVP/Test, if no header, allow for now ONLY if testing?
			// No, strict mode is better for "Platform" feel.
			response.ErrorWithStatus(c, http.StatusUnauthorized, response.CodeUnauthorized, "unauthorized - missing X-User header")
			c.Abort()
			return
		}

		user, err := service.GetUserByUsername(username)
		if err != nil {
			response.ErrorWithStatus(c, http.StatusUnauthorized, response.CodeInvalidToken, "user not found")
			c.Abort()
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
			response.ErrorWithStatus(c, http.StatusUnauthorized, response.CodeUnauthorized, "unauthorized")
			c.Abort()
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
			response.ErrorWithStatus(c, http.StatusForbidden, response.CodeUnauthorized, "forbidden")
			c.Abort()
			return
		}
		c.Next()
	}
}
