package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"IronOps/internal/database"
	"IronOps/internal/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	// Initialize Database
	database.InitDB()
	db := database.DB

	fmt.Println("Starting database seeding...")

	// Seed Users
	seedUsers(db)

	// Seed Services and Instances
	seedServices(db)

	// Seed Alerts
	seedAlerts(db)

	// Seed Audit Logs
	seedAuditLogs(db)

	fmt.Println("Database seeding completed successfully!")
}

func seedUsers(db *gorm.DB) {
	users := []model.User{
		{Username: "admin", Role: model.RoleAdmin},
		{Username: "ops_lead", Role: model.RoleOps},
		{Username: "ops_junior", Role: model.RoleOps},
		{Username: "developer_a", Role: model.RoleViewer},
		{Username: "developer_b", Role: model.RoleViewer},
	}

	for _, u := range users {
		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Failed to hash password for %s: %v", u.Username, err)
			continue
		}
		u.Password = string(hashedPassword)

		var count int64
		db.Model(&model.User{}).Where("username = ?", u.Username).Count(&count)
		if count == 0 {
			if err := db.Create(&u).Error; err != nil {
				log.Printf("Failed to create user %s: %v", u.Username, err)
			} else {
				fmt.Printf("Created user: %s\n", u.Username)
			}
		}
	}
}

func seedServices(db *gorm.DB) {
	services := []model.Service{
		{Name: "OrderService", Owner: "TeamA", Env: "prod"},
		{Name: "PaymentService", Owner: "TeamA", Env: "prod"},
		{Name: "UserService", Owner: "TeamB", Env: "prod"},
		{Name: "CatalogService", Owner: "TeamB", Env: "dev"},
		{Name: "InventoryService", Owner: "TeamC", Env: "test"},
		{Name: "NotificationService", Owner: "TeamCommon", Env: "prod"},
	}

	for _, s := range services {
		var count int64
		db.Model(&model.Service{}).Where("name = ?", s.Name).Count(&count)
		if count == 0 {
			if err := db.Create(&s).Error; err != nil {
				log.Printf("Failed to create service %s: %v", s.Name, err)
				continue
			}
			fmt.Printf("Created service: %s\n", s.Name)

			// Add instances for this service
			numInstances := rand.Intn(3) + 1
			for i := 0; i < numInstances; i++ {
				instance := model.Instance{
					ServiceID:     s.ID,
					IP:            fmt.Sprintf("10.0.%d.%d", rand.Intn(255), rand.Intn(255)),
					Port:          8080 + rand.Intn(100),
					Status:        "running",
					LastHeartbeat: time.Now(),
				}
				if rand.Float32() < 0.2 {
					instance.Status = "stopped"
				}
				db.Create(&instance)
			}
		}
	}
}

func seedAlerts(db *gorm.DB) {
	// Get all instances
	var instances []model.Instance
	db.Find(&instances)
	if len(instances) == 0 {
		return
	}

	types := []string{"cpu_high", "memory_high", "disk_full", "offline"}
	statuses := []string{"firing", "resolved"}
	messages := []string{"CPU usage > 90%", "Memory usage > 85%", "Disk usage > 95%", "Instance not responding"}

	for i := 0; i < 15; i++ {
		idx := rand.Intn(len(instances))
		typeIdx := rand.Intn(len(types))

		alert := model.Alert{
			InstanceID: instances[idx].ID,
			Type:       types[typeIdx],
			Message:    messages[typeIdx],
			Status:     statuses[rand.Intn(len(statuses))],
		}
		// Set CreatedAt manually if needed, but GORM handles it.
		// To simulate past alerts, we can update it after creation or use specific struct if we want to override.
		// For now let's just create them.

		db.Create(&alert)

		// Update CreatedAt to spread them out
		db.Model(&alert).Update("created_at", time.Now().Add(-time.Duration(rand.Intn(72))*time.Hour))
	}
	fmt.Println("Created sample alerts")
}

func seedAuditLogs(db *gorm.DB) {
	actions := []string{"create_service", "stop_instance", "delete_user", "update_config", "deploy_version"}
	users := []string{"admin", "ops_lead", "developer_a"}
	results := []string{"success", "fail"}

	for i := 0; i < 20; i++ {
		log := model.AuditLog{
			User:   users[rand.Intn(len(users))],
			Action: actions[rand.Intn(len(actions))],
			Target: fmt.Sprintf("Resource-%d", rand.Intn(100)),
			Detail: fmt.Sprintf("Operation performed by %s", users[rand.Intn(len(users))]),
			Result: results[rand.Intn(len(results))],
		}
		db.Create(&log)

		// Update CreatedAt
		db.Model(&log).Update("created_at", time.Now().Add(-time.Duration(rand.Intn(96))*time.Hour))
	}
	fmt.Println("Created sample audit logs")
}
