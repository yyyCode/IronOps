package database

import (
	"IronOps/internal/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) {
	var err error

	// Default DSN logic removed, using configuration provided DSN
	if dsn == "" {
		panic("database DSN is empty")
	}

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v. Please ensure MySQL is running and database 'ironops' exists, or set MYSQL_DSN env var.", err))
	}

	// Auto Migrate
	err = DB.AutoMigrate(
		&model.Service{},
		&model.Instance{},
		&model.User{},
		&model.AuditLog{},
		&model.Metric{},
		&model.Alert{},
		&model.AlertRule{},
		&model.AlertChannel{},
	)
	if err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}

	// Fix character set for Chinese support
	DB.Exec("ALTER TABLE alert_rules CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;")
	DB.Exec("ALTER TABLE alert_channels CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;")
}
