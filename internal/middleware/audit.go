package middleware

import (
	"IronOps/internal/database"
	"IronOps/internal/model"
	"github.com/gin-gonic/gin"
)

func AuditMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process request
		c.Next()

		// After request
		userVal, exists := c.Get("user")
		username := "anonymous"
		if exists {
			user := userVal.(*model.User)
			username = user.Username
		}

		// Only audit write operations (POST, PUT, DELETE) for MVP
		if c.Request.Method == "GET" {
			return
		}

		status := c.Writer.Status()
		result := "success"
		if status >= 400 {
			result = "fail"
		}

		// Create Audit Log
		log := model.AuditLog{
			User:   username,
			Action: c.Request.Method,
			Target: c.Request.URL.Path,
			Result: result,
			Detail: c.Errors.String(),
		}

		// In a real app, use a worker pool. Here direct DB insert is fine for MVP.
		database.DB.Create(&log)
	}
}
