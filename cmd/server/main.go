package main

import (
	"IronOps/internal/database"
	"IronOps/internal/handler"
	"IronOps/internal/middleware"
	"IronOps/internal/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.InitDB()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api/v1")

	// Public Routes
	api.POST("/register", handler.RegisterHandler)
	api.POST("/login", handler.LoginHandler)

	// Protected Routes
	// Apply AuthMiddleware to all routes below
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	protected.Use(middleware.AuditMiddleware())
	{
		// Services
		// Create: Admin, Ops
		protected.POST("/services", middleware.RoleMiddleware(model.RoleOps), handler.CreateServiceHandler)
		// List: Viewer, Ops, Admin (Everyone authenticated)
		protected.GET("/services", middleware.RoleMiddleware(model.RoleViewer, model.RoleOps), handler.ListServicesHandler)

		// Instances
		// Add: Admin, Ops
		protected.POST("/instances", middleware.RoleMiddleware(model.RoleOps), handler.AddInstanceHandler)
		// Control: Admin, Ops
		protected.POST("/instances/:id/control", middleware.RoleMiddleware(model.RoleOps), handler.ControlInstanceHandler)

		// Monitoring
		// Report Metric: Admin, Ops (Agents should authenticate)
		protected.POST("/metrics", middleware.RoleMiddleware(model.RoleOps), handler.ReportMetricHandler)
		// List Alerts: Viewer, Ops, Admin
		protected.GET("/alerts", middleware.RoleMiddleware(model.RoleViewer, model.RoleOps), handler.ListAlertsHandler)

		// Alert Rules (Admin/Ops)
		protected.POST("/alert-rules", middleware.RoleMiddleware(model.RoleOps), handler.CreateAlertRuleHandler)
		protected.GET("/alert-rules", middleware.RoleMiddleware(model.RoleViewer, model.RoleOps), handler.ListAlertRulesHandler)
		protected.DELETE("/alert-rules/:id", middleware.RoleMiddleware(model.RoleAdmin), handler.DeleteAlertRuleHandler)

		// Alert Channels (Admin only)
		protected.POST("/alert-channels", middleware.RoleMiddleware(model.RoleAdmin), handler.CreateAlertChannelHandler)
		protected.GET("/alert-channels", middleware.RoleMiddleware(model.RoleAdmin), handler.ListAlertChannelsHandler)
		protected.DELETE("/alert-channels/:id", middleware.RoleMiddleware(model.RoleAdmin), handler.DeleteAlertChannelHandler)

		// Dashboard Stats
		protected.GET("/dashboard/stats", middleware.RoleMiddleware(model.RoleViewer, model.RoleOps), handler.GetDashboardStatsHandler)

		// Audit Logs
		// List: Admin only
		protected.GET("/audits", middleware.RoleMiddleware(model.RoleAdmin), handler.ListAuditLogsHandler)

		// User Management
		protected.GET("/users", middleware.RoleMiddleware(model.RoleAdmin), handler.ListUsersHandler)
		protected.GET("/roles", middleware.RoleMiddleware(model.RoleAdmin), handler.ListRolesHandler)
	}

	// WS Route - Public for now to avoid Auth header issues with standard WebSocket API
	api.GET("/ws/dashboard", handler.DashboardWSHandler)

	r.Run(":8080")
}
