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

func CreateServiceHandler(c *gin.Context) {
	var svc model.Service
	if err := c.ShouldBindJSON(&svc); err != nil {
		response.ErrorWithStatus(c, http.StatusBadRequest, response.CodeParamError, err.Error())
		return
	}

	if err := service.CreateService(&svc); err != nil {
		logger.Error("CreateService failed", zap.Error(err))
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}

	response.Result(c, http.StatusCreated, response.CodeSuccess, "created", svc)
}

func ListServicesHandler(c *gin.Context) {
	services, err := service.ListServices()
	if err != nil {
		logger.Error("ListServices failed", zap.Error(err))
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}

	response.Success(c, services)
}

func AddInstanceHandler(c *gin.Context) {
	var instance model.Instance
	if err := c.ShouldBindJSON(&instance); err != nil {
		response.ErrorWithStatus(c, http.StatusBadRequest, response.CodeParamError, err.Error())
		return
	}

	if err := service.AddInstance(&instance); err != nil {
		logger.Error("AddInstance failed", zap.Error(err))
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}

	response.Result(c, http.StatusCreated, response.CodeSuccess, "created", instance)
}
