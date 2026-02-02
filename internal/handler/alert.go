package handler

import (
	"IronOps/internal/database"
	"IronOps/internal/model"
	"IronOps/internal/pkg/logger"
	"IronOps/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// --- Alert Rules ---

func CreateAlertRuleHandler(c *gin.Context) {
	var rule model.AlertRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		response.ErrorWithStatus(c, http.StatusBadRequest, response.CodeParamError, err.Error())
		return
	}

	if err := database.DB.Create(&rule).Error; err != nil {
		logger.Error("CreateAlertRule failed", zap.Error(err))
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}

	response.Result(c, http.StatusCreated, response.CodeSuccess, "created", rule)
}

func ListAlertRulesHandler(c *gin.Context) {
	var rules []model.AlertRule
	if err := database.DB.Find(&rules).Error; err != nil {
		logger.Error("ListAlertRules failed", zap.Error(err))
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}
	response.Success(c, rules)
}

func DeleteAlertRuleHandler(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&model.AlertRule{}, id).Error; err != nil {
		logger.Error("DeleteAlertRule failed", zap.String("id", id), zap.Error(err))
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "deleted"})
}

func UpdateAlertRuleHandler(c *gin.Context) {
	id := c.Param("id")
	var rule model.AlertRule
	if err := database.DB.First(&rule, id).Error; err != nil {
		response.ErrorWithStatus(c, http.StatusNotFound, response.CodeParamError, "rule not found")
		return
	}

	var updateData model.AlertRule
	if err := c.ShouldBindJSON(&updateData); err != nil {
		response.ErrorWithStatus(c, http.StatusBadRequest, response.CodeParamError, err.Error())
		return
	}

	// Only allow updating specific fields
	rule.Enabled = updateData.Enabled
	// Add other fields if needed, but for now we focus on Enabled

	if err := database.DB.Save(&rule).Error; err != nil {
		logger.Error("UpdateAlertRule failed", zap.String("id", id), zap.Error(err))
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}
	response.Success(c, rule)
}

// --- Alert Channels ---

func CreateAlertChannelHandler(c *gin.Context) {
	var channel model.AlertChannel
	if err := c.ShouldBindJSON(&channel); err != nil {
		response.ErrorWithStatus(c, http.StatusBadRequest, response.CodeParamError, err.Error())
		return
	}

	if err := database.DB.Create(&channel).Error; err != nil {
		logger.Error("CreateAlertChannel failed", zap.Error(err))
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}

	response.Result(c, http.StatusCreated, response.CodeSuccess, "created", channel)
}

func ListAlertChannelsHandler(c *gin.Context) {
	var channels []model.AlertChannel
	if err := database.DB.Find(&channels).Error; err != nil {
		logger.Error("ListAlertChannels failed", zap.Error(err))
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}
	response.Success(c, channels)
}

func DeleteAlertChannelHandler(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&model.AlertChannel{}, id).Error; err != nil {
		logger.Error("DeleteAlertChannel failed", zap.String("id", id), zap.Error(err))
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "deleted"})
}

func UpdateAlertChannelHandler(c *gin.Context) {
	id := c.Param("id")
	var channel model.AlertChannel
	if err := database.DB.First(&channel, id).Error; err != nil {
		response.ErrorWithStatus(c, http.StatusNotFound, response.CodeParamError, "channel not found")
		return
	}

	var updateData model.AlertChannel
	if err := c.ShouldBindJSON(&updateData); err != nil {
		response.ErrorWithStatus(c, http.StatusBadRequest, response.CodeParamError, err.Error())
		return
	}

	// Only allow updating specific fields
	channel.Enabled = updateData.Enabled

	if err := database.DB.Save(&channel).Error; err != nil {
		logger.Error("UpdateAlertChannel failed", zap.String("id", id), zap.Error(err))
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}
	response.Success(c, channel)
}
