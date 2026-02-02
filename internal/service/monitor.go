package service

import (
	"IronOps/internal/database"
	"IronOps/internal/model"
	"time"
)

func ReportMetric(metric *model.Metric) error {
	// 1. Save Metric
	if err := database.DB.Create(metric).Error; err != nil {
		return err
	}

	// 2. Update Instance Heartbeat
	database.DB.Model(&model.Instance{}).Where("id = ?", metric.InstanceID).
		Updates(map[string]interface{}{
			"last_heartbeat": time.Now(),
			"status":         "running",
		})

	// 3. Dynamic Rule Check
	EvaluateRules(metric)

	return nil
}

func ListAlerts() ([]model.Alert, error) {
	var alerts []model.Alert
	err := database.DB.Order("created_at desc").Find(&alerts).Error
	return alerts, err
}

type SystemStats struct {
	CPUUsage     float64 `json:"cpu_usage"`
	HealthScore  float64 `json:"health_score"`
	ResponseTime int     `json:"response_time"` // ms
	Stability    float64 `json:"stability"`
}

func GetSystemStats() (SystemStats, error) {
	stats := SystemStats{
		ResponseTime: 45,   // Mock default
		Stability:    92.5, // Mock default
	}

	// 1. Get latest average CPU from metrics (last 5 mins)
	var avgCPU float64
	// Just take the average of the last 10 metrics for simplicity
	result := database.DB.Model(&model.Metric{}).Select("avg(cpu)").Order("created_at desc").Limit(10).Scan(&avgCPU)
	if result.Error == nil {
		stats.CPUUsage = avgCPU
	}

	// 2. Calculate Health Score based on firing alerts
	var firingCount int64
	database.DB.Model(&model.Alert{}).Where("status = ?", "firing").Count(&firingCount)
	stats.HealthScore = 100.0 - float64(firingCount)*5.0
	if stats.HealthScore < 0 {
		stats.HealthScore = 0
	}

	return stats, nil
}
