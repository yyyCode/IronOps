package handler

import (
	"IronOps/internal/database"
	"IronOps/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- Alert Rules ---

func CreateAlertRuleHandler(c *gin.Context) {
	var rule model.AlertRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&rule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, rule)
}

func ListAlertRulesHandler(c *gin.Context) {
	var rules []model.AlertRule
	if err := database.DB.Find(&rules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rules)
}

func DeleteAlertRuleHandler(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&model.AlertRule{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

// --- Alert Channels ---

func CreateAlertChannelHandler(c *gin.Context) {
	var channel model.AlertChannel
	if err := c.ShouldBindJSON(&channel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&channel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, channel)
}

func ListAlertChannelsHandler(c *gin.Context) {
	var channels []model.AlertChannel
	if err := database.DB.Find(&channels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, channels)
}

func DeleteAlertChannelHandler(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&model.AlertChannel{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
