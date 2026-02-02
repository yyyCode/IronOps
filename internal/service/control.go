package service

import (
	"IronOps/internal/database"
	"IronOps/internal/model"
	"fmt"
)

func ControlInstance(instanceID uint, action string) error {
	var instance model.Instance
	if err := database.DB.First(&instance, instanceID).Error; err != nil {
		return err
	}

	// Mock execution
	fmt.Printf("Mock Executing action %s on instance %s (%s)\n", action, instance.IP, instance.Status)

	// Update Status based on action
	newStatus := instance.Status
	switch action {
	case "start":
		newStatus = "running"
	case "stop":
		newStatus = "stopped"
	case "restart":
		newStatus = "running"
	}

	return database.DB.Model(&instance).Update("status", newStatus).Error
}
