package handler

import (
	"IronOps/internal/model"
	"IronOps/internal/pkg/logger"
	"IronOps/internal/pkg/response"
	"IronOps/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ReportMetricHandler(c *gin.Context) {
	var metric model.Metric
	if err := c.ShouldBindJSON(&metric); err != nil {
		response.ErrorWithStatus(c, http.StatusBadRequest, response.CodeParamError, err.Error())
		return
	}

	if err := service.ReportMetric(&metric); err != nil {
		logger.Error("ReportMetric failed", zap.Error(err))
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}

	response.Result(c, http.StatusCreated, response.CodeSuccess, "ok", nil)
}

func ListAlertsHandler(c *gin.Context) {
	alerts, err := service.ListAlerts()
	if err != nil {
		logger.Error("ListAlerts failed", zap.Error(err))
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}

	response.Success(c, alerts)
}

func GetDashboardStatsHandler(c *gin.Context) {
	stats, err := service.GetSystemStats()
	if err != nil {
		logger.Error("GetDashboardStats failed", zap.Error(err))
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}

	response.Success(c, stats)
}
