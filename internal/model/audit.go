package model

import (
	"gorm.io/gorm"
)

type AuditLog struct {
	gorm.Model
	User     string `json:"user"`
	Action   string `json:"action"`
	Target   string `json:"target"`
	Result   string `json:"result"` // success, fail
	Detail   string `json:"detail"`
}
