package handler

import (
	"IronOps/internal/model"
	"IronOps/internal/monitor"
	"IronOps/internal/pkg/logger"
	"IronOps/internal/pkg/response"
	"IronOps/internal/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for dev
	},
}

func DashboardWSHandler(c *gin.Context) {
	// Authentication Check
	token := c.Query("token")
	if token == "" {
		// Try to get from header (some clients support it) or cookie if needed
		// But for now, enforce query param
		response.ErrorWithStatus(c, http.StatusUnauthorized, response.CodeUnauthorized, "missing token")
		return
	}

	// Simple token validation (in real app, verify JWT or session)
	// Here we check if the user exists based on username (token = username for this demo)
	// Or we can just check if it's not empty, assuming the frontend only sends it after login.
	// Let's do a basic check: valid user exists.
	// NOTE: In production, use proper JWT validation!
	user, err := service.GetUserByUsername(token)
	if err != nil || user == nil {
		response.ErrorWithStatus(c, http.StatusUnauthorized, response.CodeInvalidToken, "invalid token")
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("Failed to upgrade to websocket", zap.Error(err))
		return
	}
	defer conn.Close()

	// Send updates every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			stats, err := monitor.GetRealTimeStats()
			if err != nil {
				logger.Error("Failed to get system stats", zap.Error(err))
				continue
			}

			if err := conn.WriteJSON(stats); err != nil {
				logger.Info("Failed to write json to websocket (client disconnected?)", zap.Error(err))
				return // Client disconnected
			}

			// Integrate with Alert Engine: Evaluate rules against current stats
			// Using InstanceID 0 to represent the Monitoring Server itself
			go func(s *monitor.RealTimeStats) {
				metric := &model.Metric{
					InstanceID: 0,
					CPU:        s.CPUUsage,
					Memory:     s.MemoryUsage,
				}
				service.EvaluateRules(metric)
			}(stats)
		}
	}
}
