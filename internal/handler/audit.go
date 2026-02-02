package handler

import (
	"IronOps/internal/pkg/logger"
	"IronOps/internal/pkg/response"
	"IronOps/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ListAuditLogsHandler(c *gin.Context) {
	logs, err := service.ListAuditLogs()
	if err != nil {
		logger.Error("ListAuditLogs failed", zap.Error(err))
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}

	response.Success(c, logs)
}
