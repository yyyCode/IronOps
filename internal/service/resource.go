package service

import (
	"IronOps/internal/database"
	"IronOps/internal/model"
)

func CreateService(service *model.Service) error {
	return database.DB.Create(service).Error
}

func ListServices() ([]model.Service, error) {
	var services []model.Service
	// Preload instances to show full tree
	err := database.DB.Preload("Instances").Find(&services).Error
	return services, err
}

func AddInstance(instance *model.Instance) error {
	return database.DB.Create(instance).Error
}
