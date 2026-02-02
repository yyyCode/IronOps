package handler

import (
	"IronOps/internal/pkg/logger"
	"IronOps/internal/pkg/response"
	"IronOps/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ControlInstanceHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorWithStatus(c, http.StatusBadRequest, response.CodeParamError, "invalid id")
		return
	}

	var req struct {
		Action string `json:"action"` // start, stop, restart
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorWithStatus(c, http.StatusBadRequest, response.CodeParamError, err.Error())
		return
	}

	if err := service.ControlInstance(uint(id), req.Action); err != nil {
		logger.Error("ControlInstance failed",
			zap.Uint("instance_id", uint(id)),
			zap.String("action", req.Action),
			zap.Error(err),
		)
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "action executed"})
}
