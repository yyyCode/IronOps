package handler

import (
	"IronOps/internal/model"
	"IronOps/internal/monitor"
	"IronOps/internal/service"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for dev
	},
}

func DashboardWSHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade to websocket: %v", err)
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
				log.Printf("Failed to get system stats: %v", err)
				continue
			}

			if err := conn.WriteJSON(stats); err != nil {
				log.Printf("Failed to write json to websocket: %v", err)
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
