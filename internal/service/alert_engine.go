package service

import (
	"IronOps/internal/database"
	"IronOps/internal/model"
	"fmt"
	"log"
)

// EvaluateRules checks the metric against all active rules
func EvaluateRules(metric *model.Metric) {
	var rules []model.AlertRule
	// Optimisation: Cache rules or only fetch rules for this metric type
	// For MVP: Fetch all enabled rules
	if err := database.DB.Where("enabled = ?", true).Find(&rules).Error; err != nil {
		log.Printf("Failed to fetch alert rules: %v", err)
		return
	}

	for _, rule := range rules {
		// 1. Check if rule matches metric type (e.g., cpu, memory)
		// Assuming metric struct has fields like CPU, Memory.
		// We need to map rule.MetricType to actual value in metric struct.
		var value float64
		switch rule.MetricType {
		case "cpu":
			value = metric.CPU
		case "memory":
			value = metric.Memory
		default:
			continue
		}

		// 2. Evaluate Condition
		triggered := false
		switch rule.Condition {
		case ">":
			if value > rule.Threshold {
				triggered = true
			}
		case "<":
			if value < rule.Threshold {
				triggered = true
			}
		case "=":
			if value == rule.Threshold {
				triggered = true
			}
		}

		if triggered {
			log.Printf("Rule %d (%s) triggered! Value: %.2f", rule.ID, rule.Name, value)
			fireAlert(rule, metric, value)
		} else {
			resolveAlert(rule, metric)
		}
	}
}

func fireAlert(rule model.AlertRule, metric *model.Metric, value float64) {
	// Check if already firing to avoid spam
	var existingAlert model.Alert
	result := database.DB.Where("rule_id = ? AND instance_id = ? AND status = ?", rule.ID, metric.InstanceID, "firing").First(&existingAlert)

	if result.RowsAffected > 0 {
		// Already firing, maybe update LastSeen or count
		return
	}

	log.Printf("Firing new alert for Rule %d on Instance %d", rule.ID, metric.InstanceID)

	// Create new alert
	msg := fmt.Sprintf("%s: Current value %.2f %s %.2f", rule.Description, value, rule.Condition, rule.Threshold)
	alert := model.Alert{
		InstanceID: metric.InstanceID,
		RuleID:     rule.ID,
		Type:       rule.MetricType + "_alert",
		Message:    msg,
		Status:     "firing",
	}
	database.DB.Create(&alert)

	// Send Notification
	SendNotification(NotificationPayload{
		Title:   fmt.Sprintf("Alert Triggered: %s", rule.Name),
		Message: fmt.Sprintf("Instance ID: %d\n%s", metric.InstanceID, msg),
		Level:   rule.Severity,
	})
}

func resolveAlert(rule model.AlertRule, metric *model.Metric) {
	// Check if there is an active alert to resolve
	var existingAlert model.Alert
	result := database.DB.Where("rule_id = ? AND instance_id = ? AND status = ?", rule.ID, metric.InstanceID, "firing").First(&existingAlert)

	if result.RowsAffected > 0 {
		// Resolve it
		database.DB.Model(&existingAlert).Update("status", "resolved")

		// Send Resolution Notification (Optional)
		SendNotification(NotificationPayload{
			Title:   fmt.Sprintf("Alert Resolved: %s", rule.Name),
			Message: fmt.Sprintf("Instance ID: %d is back to normal.", metric.InstanceID),
			Level:   "info",
		})
	}
}
