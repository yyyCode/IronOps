package model

import "gorm.io/gorm"

type RoleType string

const (
	RoleAdmin  RoleType = "admin"
	RoleOps    RoleType = "ops"
	RoleViewer RoleType = "viewer"
)

type User struct {
	gorm.Model
	Username string   `json:"username" gorm:"uniqueIndex;type:varchar(100)"`
	Password string   `json:"-"` // Store hash in real app
	Role     RoleType `json:"role"`
}
