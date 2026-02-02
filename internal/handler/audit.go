package handler

import (
	"IronOps/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAuditLogsHandler(c *gin.Context) {
	logs, err := service.ListAuditLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}
