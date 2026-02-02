package service

import (
	"IronOps/internal/database"
	"IronOps/internal/model"
)

func ListAuditLogs() ([]model.AuditLog, error) {
	var logs []model.AuditLog
	err := database.DB.Order("created_at desc").Find(&logs).Error
	return logs, err
}
