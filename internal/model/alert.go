package model

import (
	"gorm.io/gorm"
)

// AlertRule defines the condition to trigger an alert
type AlertRule struct {
	gorm.Model
	Name        string  `json:"name"`
	MetricType  string  `json:"metric_type"` // cpu, memory, disk, response_time
	Condition   string  `json:"condition"`   // >, <, =
	Threshold   float64 `json:"threshold"`
	Severity    string  `json:"severity"`    // critical, warning, info
	Enabled     bool    `json:"enabled"`
	Description string  `json:"description"`
}

// AlertChannel defines where to send the notifications
type AlertChannel struct {
	gorm.Model
	Name    string `json:"name"`
	Type    string `json:"type"` // feishu, dingtalk, email, webhook
	Config  string `json:"config"` // JSON string storing url, token, etc.
	Enabled bool   `json:"enabled"`
}

// Update Alert struct to link to Rule
type Alert struct {
	gorm.Model
	InstanceID uint   `json:"instance_id"`
	RuleID     uint   `json:"rule_id"` // Optional: link to the rule that triggered it
	Type       string `json:"type"`    // cpu_high, etc. (Can be derived from Rule)
	Message    string `json:"message"`
	Status     string `json:"status"` // firing, resolved
}
