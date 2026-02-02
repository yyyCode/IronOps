package model

import (
	"time"

	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name      string     `json:"name"`
	Owner     string     `json:"owner"`
	Env       string     `json:"env"` // dev, test, prod
	Instances []Instance `json:"instances"`
}

type Instance struct {
	gorm.Model
	ServiceID     uint      `json:"service_id"`
	IP            string    `json:"ip"`
	Port          int       `json:"port"`
	Status        string    `json:"status"` // running, stopped, error
	LastHeartbeat time.Time `json:"last_heartbeat"`
}
