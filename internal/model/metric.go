package model

import "gorm.io/gorm"

type Metric struct {
	gorm.Model
	InstanceID uint    `json:"instance_id"`
	CPU        float64 `json:"cpu"` // Percentage
	Memory     float64 `json:"memory"` // Percentage
}
