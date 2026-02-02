package database

import (
	"IronOps/internal/model"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	// Default DSN (Data Source Name)
	// format: user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:12345678@tcp(10.21.32.13:3306)/ironops?charset=utf8mb4&parseTime=True&loc=Local"
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
