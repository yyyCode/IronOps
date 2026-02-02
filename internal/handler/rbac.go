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

func RegisterHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Warn("Register bind failed", zap.Error(err))
		response.ErrorWithStatus(c, http.StatusBadRequest, response.CodeParamError, err.Error())
		return
	}

	if err := service.CreateUser(&user); err != nil {
		logger.Error("Register create user failed", zap.Error(err))
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}

	logger.Info("User registered", zap.String("username", user.Username))
	response.Result(c, http.StatusCreated, response.CodeSuccess, "registered successfully", user)
}

func LoginHandler(c *gin.Context) {
	// Simple mock login
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&creds); err != nil {
		response.ErrorWithStatus(c, http.StatusBadRequest, response.CodeParamError, err.Error())
		return
	}

	user, err := service.GetUserByUsername(creds.Username)
	if err != nil {
		logger.Warn("Login failed - user not found", zap.String("username", creds.Username))
		response.ErrorWithStatus(c, http.StatusUnauthorized, response.CodeUnauthorized, "invalid credentials")
		return
	}

	// Verify password
	if !service.CheckPassword(user, creds.Password) {
		logger.Warn("Login failed - wrong password", zap.String("username", creds.Username))
		response.ErrorWithStatus(c, http.StatusUnauthorized, response.CodeUnauthorized, "invalid credentials")
		return
	}

	logger.Info("User logged in", zap.String("username", user.Username))
	response.Success(c, gin.H{
		"message": "login success",
		"user":    user,
	})
}

func ListUsersHandler(c *gin.Context) {
	users, err := service.ListUsers()
	if err != nil {
		logger.Error("List users failed", zap.Error(err))
		response.ErrorWithStatus(c, http.StatusInternalServerError, response.CodeServerBusy, err.Error())
		return
	}
	response.Success(c, users)
}

func ListRolesHandler(c *gin.Context) {
	roles := []model.RoleType{model.RoleAdmin, model.RoleOps, model.RoleViewer}
	response.Success(c, roles)
}
