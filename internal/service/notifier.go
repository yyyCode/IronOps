package service

import (
	"IronOps/internal/database"
	"IronOps/internal/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type NotificationPayload struct {
	Title   string
	Message string
	Level   string // critical, warning, info
}

// SendNotification dispatches the alert to all enabled channels
func SendNotification(payload NotificationPayload) {
	var channels []model.AlertChannel
	if err := database.DB.Where("enabled = ?", true).Find(&channels).Error; err != nil {
		log.Printf("Failed to load alert channels: %v", err)
		return
	}

	log.Printf("Found %d enabled alert channels. Dispatching notification...", len(channels))

	for _, channel := range channels {
		go func(ch model.AlertChannel) {
			var config map[string]string
			if err := json.Unmarshal([]byte(ch.Config), &config); err != nil {
				log.Printf("Invalid config for channel %s: %v. Config: %s", ch.Name, err, ch.Config)
				return
			}

			url := config["url"]
			if url == "" {
				log.Printf("URL is empty for channel %s", ch.Name)
				return
			}

			log.Printf("Sending notification to channel %s (%s)", ch.Name, ch.Type)

			var err error
			switch ch.Type {
			case "feishu":
				err = sendToFeishu(url, payload)
			case "dingtalk":
				err = sendToDingTalk(url, payload)
			case "webhook":
				// Generic webhook
				err = postJSON(url, payload)
			default:
				log.Printf("Unknown channel type: %s", ch.Type)
			}

			if err != nil {
				log.Printf("Failed to send notification to %s: %v", ch.Name, err)
			}
		}(channel)
	}
}

// We need to implement specific senders
func sendToFeishu(url string, payload NotificationPayload) error {
	// Simple Feishu Webhook format
	// Docs: https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
	msg := map[string]interface{}{
		"msg_type": "text",
		"content": map[string]string{
			"text": fmt.Sprintf("[%s] IronOps Alert:\n%s\n%s", payload.Level, payload.Title, payload.Message),
		},
	}
	return postJSON(url, msg)
}

func sendToDingTalk(url string, payload NotificationPayload) error {
	// Simple DingTalk Webhook format
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": fmt.Sprintf("[%s] IronOps Alert:\n%s\n%s", payload.Level, payload.Title, payload.Message),
		},
	}
	return postJSON(url, msg)
}

func postJSON(url string, data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return fmt.Errorf("notification failed with status: %d, body: %s", resp.StatusCode, string(respBody))
	}

	// Log success for debugging (optional)
	// log.Printf("Notification sent to %s. Status: %d", url, resp.StatusCode)

	return nil
}
